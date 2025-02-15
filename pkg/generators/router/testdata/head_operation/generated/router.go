package openapi

// This file is auto-generated, don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi"
)

// HeadHandlerHandler handles the operations of the 'HeadHandler' handler group.
type HeadHandlerHandler interface {
	// HeadOp
	HeadOp(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new router for the spec and the given handlers.
//
//
//
//
//
//
func NewRouter(
	headHandlerHandler HeadHandlerHandler,
) http.Handler {

	r := chi.NewRouter()

	// 'HeadHandler' group

	// '/a/path'
	r.Options("/a/path", optionsHandlerFunc(
		http.MethodHead,
	))
	r.Head("/a/path", headHandlerHandler.HeadOp)

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
