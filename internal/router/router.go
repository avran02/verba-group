package router

import (
	"net/http"

	"github.com/avran02/verba-group/internal/controller"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	swagger "github.com/swaggo/http-swagger"
)

type Router struct {
	*chi.Mux
	c controller.Controller
}

func (r *Router) getSwaggerRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/docs/openapi.yml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./docs/openapi.yml")
	})
	router.Get("/swagger/*", swagger.Handler(
		swagger.URL("/docs/openapi.yml"),
	))
	router.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
	})
	return router
}

func (r *Router) getTodoListRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/", r.c.CreateTaskHandler)
	router.Get("/", r.c.ListTasksHandler)
	router.Get("/{id}", r.c.GetTaskHandler)
	router.Put("/{id}", r.c.UpdateTaskHandler)
	router.Delete("/{id}", r.c.DeleteTaskHandler)
	return router
}

func New(controller controller.Controller) *Router {
	router := Router{c: controller}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Mount("/", router.getSwaggerRoutes())
	r.Mount("/tasks", router.getTodoListRoutes())

	router.Mux = r
	return &router
}
