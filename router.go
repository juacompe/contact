package contact

import "github.com/ant0ine/go-json-rest/rest"

type Router struct {
	All, Get, Delete, Update, Add *rest.Route
}

func NewRouter() *Router {
	handler := NewHandler()

	All := &rest.Route{"GET", "/contact", handler.All}
	Get := &rest.Route{"GET", "/contact/:id", handler.Get}
	Delete := &rest.Route{"DELETE", "/contact/:id", handler.Delete}
	Update := &rest.Route{"PUT", "/contact/:id", handler.Get}
	Add := &rest.Route{"POST", "/contact", handler.Add}

	return &Router{All: All, Get: Get, Delete: Delete, Update: Update, Add: Add}
}

func MakeRestRouter() (rest.App, error) {
	contact := NewRouter()

	return rest.MakeRouter(
		contact.All,
		contact.Get,
		contact.Delete,
		contact.Update,
		contact.Add,
	)
}