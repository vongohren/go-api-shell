package main

import (
  "net/http"
  "github.com/Snorlock/go-api-shell/handlers"
  "github.com/Snorlock/go-api-shell/middleware"
  "github.com/Snorlock/go-api-shell/db"
  "github.com/Snorlock/mux"
)

func NewRouter(env *db.Env) *http.ServeMux {
  // implement your router variables here, example uses mux and only need a http router capable with go http
  var authorize = true
  router := mux.NewRouter();
  router.Handle("/auth/{provider}", middleware.AuthHandler{middleware.Handler{env, !authorize}, handlers.BeginAuthHandler}).Methods("GET")
  router.Handle("/auth/{provider}/callback", middleware.AuthHandler{middleware.Handler{env, !authorize}, handlers.CallBack}).MakePrivate()
  router.Handle("/add", middleware.TokenHandler{middleware.Handler{env, authorize}, handlers.AddHandler}).Methods("POST")
  router.Handle("/list", middleware.TokenHandler{middleware.Handler{env, authorize}, handlers.GetListHandler}).Methods("GET")

  mx := http.NewServeMux()
  mx.Handle("/", router)
  return mx
}
