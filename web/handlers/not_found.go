package handlers

import (
	"context"
	"log/slog"
	"net/http"
	"workshop/internal/domain/sighting"
	"workshop/web/views/pages"
)

func NotFound(logger *slog.Logger, sightings SightingsLister) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderNotFound(r.Context(), w, logger, sightings)
	}
}

// renderNotFound answers 404 with the lost-trail page and a few recent sightings to try instead.
func renderNotFound(ctx context.Context, w http.ResponseWriter, logger *slog.Logger, sightings SightingsLister) {
	recent, err := sightings.List(ctx, sighting.ListFilter{Limit: 24})
	if err != nil {
		logger.ErrorContext(ctx, "handlers: listing sightings for 404: "+err.Error())
		recent = nil
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	if err = pages.NotFound(recent).Render(ctx, w); err != nil {
		logger.ErrorContext(ctx, "handlers: rendering 404 page: "+err.Error())
	}
}
