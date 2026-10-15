package components

import (
	"cmp"
	"fmt"
	"slices"

	"workshop/internal/domain/sighting"
)

// YearGroup is one year on the timeline and the sightings in it, newest first.
type YearGroup struct {
	Year      int
	Sightings []sighting.Sighting
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
		groups[len(groups)-1].Sightings = append(groups[len(groups)-1].Sightings, s)
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

// morePhotos labels how many photos an entry has beyond the one shown.
func morePhotos(n int) string {
	if n == 1 {
		return "+1 photo"
	}
	return fmt.Sprintf("+%d photos", n)
}

// seenTimes labels how many sightings a year holds.
func seenTimes(n int) string {
	if n == 1 {
		return "1 sighting"
	}
	return fmt.Sprintf("%d sightings", n)
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
