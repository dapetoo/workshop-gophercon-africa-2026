package pages

import "workshop/internal/domain/sighting"

// topCountries ranks the countries with the most sightings.
func topCountries(sightings []sighting.Sighting, n int) []sighting.Tally {
	return sighting.Top(sighting.CountBy(sightings, func(s sighting.Sighting) string {
		return s.Location.Country
	}), n)
}
