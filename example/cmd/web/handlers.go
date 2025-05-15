package main

import (
	"net/http"
	"fmt"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello design patterns")
}