package routes

import (
	"github.com/gominima/minima"
)

func GraphHandler() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
      res.Send("Hello")
	}
}