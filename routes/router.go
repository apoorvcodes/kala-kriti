package routes


import ( 
	"github.com/gominima/minima"
)

func MainRouter() *minima.Router {
	return minima.NewRouter().Get("/graph/:format", GraphHandler()).Get("/list", ListHandler())
}