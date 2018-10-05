package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/michaljemala/playground-go/chirouter/pkg/advertiser"
	"github.com/michaljemala/playground-go/chirouter/pkg/campaign"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(10 * time.Second))

	r.Mount("/advertiser", advertiser.NewRestHandler())
	r.Mount("/campaign", campaign.NewRestHandler())
	r.NotFound(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "NO! NO! NO! a.k.a. Triple UPPERCASED NO:-)")
	})
	http.ListenAndServe(":8888", r)
}
