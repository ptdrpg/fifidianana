package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/efidy/app"
	"github.com/ptdrpg/efidy/controller"
	"github.com/ptdrpg/efidy/repository"
	"github.com/ptdrpg/efidy/router"
)

func main () {
	mainR:= gin.Default()
	app.Connexion()
	db := app.DB
	repo := repository.NewRepository(db)
	controller := controller.NewController(db, repo)
	r := router.NewRouter(mainR, controller)
	r.RegisterRouter()

	r.R.Run(":4400")
}