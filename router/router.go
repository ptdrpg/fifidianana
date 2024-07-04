package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/efidy/controller"
)

type Router struct {
	R *gin.Engine
	C *controller.Controller
}

func NewRouter(r *gin.Engine, c *controller.Controller) *Router {
	return &Router{
		R: r,
		C: c,
	}
}

func (r *Router) RegisterRouter() {
	r.R.Static("/upload", "./image")

	apiR := r.R.Group("/api")
	v1 := apiR.Group("/v1")

	cr := v1.Group("/candidat")
	cr.GET("/", r.C.FindAllCandidat)
	cr.GET("/:id", r.C.FindCandidatById)
	cr.GET("/men", r.C.FindAllMen)
	cr.GET("/women", r.C.FindAllWoman)
	cr.POST("/", r.C.CreateCandidat)
	cr.POST("/avatar/:id", r.C.UploadCandidatAvatar)
	cr.PUT("/:id", r.C.UpdateCandidat)
	cr.DELETE("/:id", r.C.DeleteCandidat)

	br := v1.Group("/bulletin")
	br.GET("/", r.C.FindAllBulletin)
	br.GET("/", r.C.FindBulletinByNum)
	br.POST("/", r.C.SaveBulletin)

}
