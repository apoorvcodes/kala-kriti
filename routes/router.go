package routes

import (
	"github.com/gominima/minima"
)

func MainRouter() *minima.Router {
	return minima.NewRouter().Get("/graph/:format", GraphHandler()).Get("/list", ListHandler()).Post(("/save/:toxic/:read/:level"), SaveHandler()).Get("/save/:toxic/:read/:level", SaveHandler())
}
