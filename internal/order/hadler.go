package order

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	orderIDURL     = "/order/"
	orderSearchURL = "/order"
)

type (
	Handler struct {
		order Service
	}
)

func NewHandler(service *Service) *Handler {
	return &Handler{*service}
}

func (h *Handler) Register(r chi.Router) {
	router := chi.NewRouter()
	router.Get(orderIDURL, func(w http.ResponseWriter, r *http.Request) {
		h.getOrderByID(w, r)
	})

	router.Get(orderSearchURL, func(w http.ResponseWriter, r *http.Request) {
		h.searchOrder(w)
	})

	r.Mount("/", router)
}

func (h *Handler) getOrderByID(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	oid := r.Form.Get("ouid")
	orderByUID, err := h.order.Get(oid)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		// w.Write([]byte(err.Error()))
		BadSearch(w, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	OrderResponseTable(w, orderByUID)
}

func (h *Handler) searchOrder(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	Search(w)
}
