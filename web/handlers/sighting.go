package handlers

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"uuid"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

type SightingFinder interface {
	Find(ctx context.Context, id uuid.UUID) (sighting.Sighting, error)
	List(ctx context.Context, filter sighting.ListFilter) ([]sighting.Sighting, error)
}

func Sighting(logger *slog.Logger, sightings SightingFinder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			renderNotFound(ctx, w, logger, sightings)
			return
		}

		s, err := sightings.Find(ctx, id)
		if errors.Is(err, sighting.ErrNotFound) {
			renderNotFound(ctx, w, logger, sightings)
			return
		}
		if err != nil {
			logger.ErrorContext(ctx, "handlers: finding sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to load this sighting.")
			return
		}

		nearby, err := sightings.List(ctx, sighting.ListFilter{Country: s.Location.CountryCode, Limit: 5})
		if err != nil {
			logger.ErrorContext(ctx, "handlers: listing nearby sightings: "+err.Error())
			nearby = nil
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err = pages.Sighting(s, others(nearby, s.ID, 4)).Render(ctx, w); err != nil {
			logger.ErrorContext(ctx, "handlers: rendering sighting: "+err.Error())
			renderError(ctx, w, logger, http.StatusInternalServerError, "Unable to render page.")
		}
	}
}

// others drops the sighting with id from list and keeps at most n.
func others(list []sighting.Sighting, id uuid.UUID, n int) []sighting.Sighting {
	out := make([]sighting.Sighting, 0, n)
	for _, s := range list {
		if s.ID != id && len(out) < n {
			out = append(out, s)
		}
	}
	return out
}
