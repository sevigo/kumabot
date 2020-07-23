package version

import (
	"net/http"

	"github.com/sevigo/kumabot/pkg/core"
	"github.com/sevigo/kumabot/pkg/handler/json"
	"github.com/sevigo/kumabot/pkg/version"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// New returns a new health check router
func New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Handle("/", Handler())
	return r
}

// Handler creates an http.HandlerFunc that performs system healthchecks
func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v := core.Version{
			Revision: version.Revision,
			Version:  version.Version,
		}
		json.JSON(w, &v, 200)
	}
}
