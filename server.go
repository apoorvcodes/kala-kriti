package main

import (
	"fmt"
	"kala/database"
	"kala/routes"
	"os"
    "github.com/joho/godotenv"
	"github.com/gominima/cors"
	"github.com/gominima/middlewares"
	"github.com/gominima/minima"
)

func main() {
	m := minima.New()
     err := godotenv.Load()
	 if err != nil {
		panic(err)
	 }

	db := Db.Connect(os.Getenv("DATABASE_URL"))
	fmt.Print(db)
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

