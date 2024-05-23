package main

import (
	"brujulavirtual-auth/src/register"
	"net/http"
)

func App(mux *http.ServeMux) {
	register.Module(mux)
}
