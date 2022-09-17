package main

import (
	"fmt"
	"github.com/gominima/cors"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
	"github.com/joho/godotenv"
	"kala/routes"
	"os"
)

func main() {
	m := minima.New()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	m.UseRaw(middleware.Logger)
	m.UseRaw(middleware.RouteHeaders().Handler)
	m.UseRaw(middleware.Heartbeat("/graph"))
	m.UseRaw(middleware.Recoverer)
	m.Use(cors.New().AllowAll())
	m.Get("/", func(res *minima.Response, req *minima.Request) {
		res.JSON("Hello")
	})

	m.UseRouter(routes.MainRouter())
	fmt.Print(os.Getenv("PORT"))
	m.Listen(":" + os.Getenv("PORT"))
}
