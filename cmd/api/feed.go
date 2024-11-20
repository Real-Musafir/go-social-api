package main

import (
	"net/http"

	"github.com/Real-Musafir/social/internal/store"
)

func (app *application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {

	// pagination, filters,

	fq := store.PaginateFeedQuery {
		Limit: 20,
		Offset: 0,
		Sort: "desc",
	}

	fq, err := fq.Parse(r)
	if err!= nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()
	feed, err := app.store.Posts.GetUserFeed(ctx, int64(92), fq)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusOK, feed); err != nil {
		app.internalServerError(w, r, err)
	}
}