package handlers

import (
	"github.com/go-chi/chi/v5"

	"wb_L0/internal/order"
	"wb_L0/internal/services"
)

type (
	Handlers struct {
		Order  *order.Handler
		router chi.Router
	}
)

func NewHandlers(services *services.Services, r chi.Router) *Handlers {
	return &Handlers{
		Order:  order.NewHandler(&services.Order),
		router: r,
	}
}

func (h *Handlers) Register() {
	r := chi.NewRouter()
	r.Route("/api", func(route chi.Router) {
		route.Route("/v1", func(route chi.Router) {
			h.Order.Register(route)
		})
	})

	h.router.Mount("/", r)
}
