package handlers

import(
  "net/http"
  "fmt"
  "github.com/Snorlock/go-api-shell/db"
  "github.com/Snorlock/mux"
)

type Apis struct {
  Paths []Route
}

type Route struct {
  Path  string
  Methods []string
}

func IndexHandler(env *db.Env, routes []*mux.Route, w http.ResponseWriter, r *http.Request) error {
  fmt.Fprint(w, "Hello world")
  return nil
}
