package main

import (
	"fmt"
	"net/http"

	"github.com/Real-Musafir/social/internal/store"
)

type CreatePostPayload struct {
	Title 		string		`json:"title"`
	Content 	string		`json:"content"`
	Tags		[]string	`json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request)  {

	var payload CreatePostPayload
	fmt.Printf("Check before payload", payload)
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Printf("Check after payload", payload)

	

	post := &store.Post {
		Title: payload.Title,
		Content: payload.Content,
		Tags: payload.Tags,
		UserId: 1,
	}
	
	ctx := r.Context()

	if err:= app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}