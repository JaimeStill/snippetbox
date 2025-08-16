package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/JaimeStill/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/base.tmpl",
	// 	"./ui/html/partials/nav.tmpl",
	// 	"./ui/html/pages/home.tmpl",
	// }

	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }

	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// 	app.serverError(w, r, err)
	// }
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Snippet: snippet,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display a form for creating a new snippet")
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O Snail"
	content := "O Snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}
