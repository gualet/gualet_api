package router

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    name        string
    method      string
    pattern     string
    handlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    addUserRoutes(router)

    return router
}
