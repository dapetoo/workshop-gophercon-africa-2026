package components

import (
	"cmp"
	"fmt"
	"net/url"
	"slices"
	"strings"
	"unicode"

	"workshop/internal/domain/sighting"
)

// perYear is how many sightings a year shows before the rest fold away.
const perYear = 9

// YearGroup is one year on the timeline: the first few sightings, newest first, and the rest.
type YearGroup struct {
	Year  int
	Shown []sighting.Sighting
	More  []sighting.Sighting
}

// ByYear sorts sightings newest first and groups them by the year they happened.
func ByYear(sightings []sighting.Sighting) []YearGroup {
	sorted := slices.SortedFunc(slices.Values(sightings), func(a, b sighting.Sighting) int {
		return b.HappenedAt.Compare(a.HappenedAt)
	})
	var groups []YearGroup
	for _, s := range sorted {
		y := s.HappenedAt.Year()
		if len(groups) == 0 || groups[len(groups)-1].Year != y {
			groups = append(groups, YearGroup{Year: y})
		}
		g := &groups[len(groups)-1]
		if len(g.Shown) < perYear {
			g.Shown = append(g.Shown, s)
		} else {
			g.More = append(g.More, s)
		}
	}
	return groups
}

// photo picks the first still image with a URL.
func photo(media []sighting.Media) (sighting.Media, bool) {
	for _, m := range media {
		if m.Type == "StillImage" && m.URL != "" {
			return m, true
		}
	}
	return sighting.Media{}, false
}

// sized asks iNaturalist for a smaller copy of a photo ("medium" or "large"); other hosts are left alone.
func sized(raw, size string) string {
	u, err := url.Parse(raw)
	if err != nil || !strings.HasSuffix(u.Host, "inaturalist-open-data.s3.amazonaws.com") && u.Host != "static.inaturalist.org" {
		return raw
	}
	dir, file, ok := strings.Cut(u.Path[1:], "/original.")
	if !ok || strings.Contains(file, "/") {
		return raw
	}
	u.Path = "/" + dir + "/" + size + "." + file
	return u.String()
}

// Hero picks the sighting for the big banner: a big cat or elephant if one has a photo, else the latest.
func Hero(sightings []sighting.Sighting) (sighting.Sighting, bool) {
	for _, want := range []string{"leopard", "lion", "cheetah", "african elephant", "giraffe"} {
		for _, s := range sightings {
			if _, ok := photo(s.Media); ok && strings.EqualFold(displayName(s.Species), want) {
				return s, true
			}
		}
	}
	return Latest(sightings)
}

// Latest is the most recent sighting that has a photo.
func Latest(sightings []sighting.Sighting) (sighting.Sighting, bool) {
	var best sighting.Sighting
	found := false
	for _, s := range sightings {
		if _, ok := photo(s.Media); ok && (!found || s.HappenedAt.After(best.HappenedAt)) {
			best, found = s, true
		}
	}
	return best, found
}

// Regular is one of the most-seen animals, with a sighting to show it by.
type Regular struct {
	Name     string
	Count    int
	Sighting sighting.Sighting
}

// Regulars lists the n most-seen animals that have a photo, each with its latest photographed sighting.
func Regulars(sightings []sighting.Sighting, n int) []Regular {
	var out []Regular
	for _, t := range TopAnimals(sightings, len(sightings)) {
		var pick sighting.Sighting
		found := false
		for _, s := range sightings {
			if _, ok := photo(s.Media); ok && displayName(s.Species) == t.Name && (!found || s.HappenedAt.After(pick.HappenedAt)) {
				pick, found = s, true
			}
		}
		if found {
			out = append(out, Regular{Name: t.Name, Count: t.Count, Sighting: pick})
		}
		if len(out) == n {
			break
		}
	}
	return out
}

// title capitalises each word, since GBIF's common names arrive in any case.
func title(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		r := []rune(w)
		r[0] = unicode.ToUpper(r[0])
		words[i] = string(r)
	}
	return strings.Join(words, " ")
}

// plural labels a count with a noun, adding an s past one.
func plural(n int, noun string) string {
	if n == 1 {
		return "1 " + noun
	}
	return fmt.Sprintf("%d %ss", n, noun)
}

// cmpOr returns s, or fallback when s is empty.
func cmpOr(s, fallback string) string {
	return cmp.Or(s, fallback)
}

// TopAnimals ranks the animals seen most often.
func TopAnimals(sightings []sighting.Sighting, n int) []sighting.Tally {
	return sighting.Top(sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return displayName(s.Species)
	}), n)
}
