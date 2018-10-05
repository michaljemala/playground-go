package advertiser

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type restHandler struct {
	r chi.Router
}

func NewRestHandler() http.Handler {
	h := &restHandler{r: chi.NewRouter()}
	h.r.Get("/", h.List)
	h.r.Get("/{id}", h.Get)
	h.r.Post("/", h.Create)
	h.r.Put("/{id}", h.Update)
	h.r.Delete("/{id}", h.Delete)
	return h
}

func (h *restHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(rw, r)
}

func (h *restHandler) List(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s -> advertiser.restHandler.List", r.Method, r.RequestURI)
}

func (h *restHandler) Get(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s -> advertiser.restHandler.Get(%s)", r.Method, r.RequestURI, chi.URLParam(r, "id"))
}

func (h *restHandler) Create(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s -> advertiser.restHandler.Create", r.Method, r.RequestURI)
}

func (h *restHandler) Update(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s -> advertiser.restHandler.Update(%s)", r.Method, r.RequestURI, chi.URLParam(r, "id"))
}

func (h *restHandler) Delete(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s %s -> advertiser.restHandler.Delete(%)", r.Method, r.RequestURI, chi.URLParam(r, "id"))
}
