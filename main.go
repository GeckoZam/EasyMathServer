package main

import (
	"github.com/GeckoZam/EasyMathServer/routes"
	"github.com/GeckoZam/EasyMathServer/storage"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	storage.InitializeDB()

	app := iris.Default()

	user := app.Party("/api/users")
	{
		user.Post("/register", routes.Register)
	}

	app.Listen(":4000")
}
