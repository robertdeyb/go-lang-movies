package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

//helper function for http router
func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		//adding the context back to the new http router
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

//receiver of app pointer to the application
func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)
	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovie)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getAllGenres)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:genre_id", app.getAllMoviesByGenre)
	router.POST("/v1/admin/editmovie", app.wrap(secure.ThenFunc(app.editMovie)))
	// router.HandlerFunc(http.MethodPost, "/v1/admin/editmovie", app.editMovie)
	router.GET("/v1/admin/deletemovie/:id", app.wrap(secure.ThenFunc(app.deleteMovie)))
	// router.HandlerFunc(http.MethodGet, "/v1/admin/deletemovie/:id", app.deleteMovie)
	return app.enableCORS(router)
}
