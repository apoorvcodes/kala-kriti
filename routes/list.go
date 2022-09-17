package routes

import (
	"github.com/gominima/minima"

)
type KalaKritiData struct {
   Date int
   Toxic bool
   Readings string
   Level string
}
type DataArray struct {
	Data []*KalaKritiData
}


func ListHandler() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
      res.OK().JSON(&DataArray{
		Data: []*KalaKritiData{
			{
				Date: 123,
				Toxic: true,
				Readings: "abc",
				Level: "Healthy",
			},
		},
	  })
	}
}