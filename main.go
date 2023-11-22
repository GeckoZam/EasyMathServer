package main

import (
	"github.com/GeckoZam/EasyMathServer/routes"
	"github.com/GeckoZam/EasyMathServer/storage"
	"github.com/go-playground/validator/v10"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()
	storage.InitializeDB()

	app := iris.Default()
	app.Validator = validator.New()

	user := app.Party("/api/users")
	{
		user.Post("/register", routes.Register)
		user.Post("/login", routes.Login)
	}

	app.Listen(":4000")
}
