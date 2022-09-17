package routes

import (
	"github.com/gominima/minima"
)

type KalaKritiData struct {
	Date     int  `bson:"date,omitempty"`
	Toxic    bool  `bson:"toxic,omitempty"`
	Readings string  `bson:"readings,omitempty"`
	Level    string  `bson:"level,omitempty"`
}
type DataArray struct {
	Data []*KalaKritiData
}

func ListHandler() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		res.OK().JSON(&DataArray{
			Data: []*KalaKritiData{
				{
					Date:     123,
					Toxic:    true,
					Readings: "abc",
					Level:    "Healthy",
				},
			},
		})
	}
}
