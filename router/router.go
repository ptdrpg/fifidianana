package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ptdrpg/efidy/controller"
	docs "github.com/ptdrpg/efidy/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//  "net/http"
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS request")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (r *Router) RegisterRouter() {
	r.R.Static("/upload", "./image")
	r.R.Use(CORSMiddleware())
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.R.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiR := r.R.Group("/api")
	v1 := apiR.Group("/v1")

	cr := v1.Group("/candidat")
	cr.GET("/", r.C.FindAllCandidat)
	cr.GET("/:id", r.C.FindCandidatByNum)
	cr.GET("/men", r.C.FindAllMen)
	cr.GET("/women", r.C.FindAllWoman)
	cr.POST("/", r.C.CreateCandidat)
	cr.POST("/avatar/:id", r.C.UploadCandidatAvatar)
	cr.PUT("/:id", r.C.UpdateCandidat)
	cr.DELETE("/:id", r.C.DeleteCandidat)

	br := v1.Group("/bulletin")
	br.GET("/", r.C.FindAllBulletin)
	br.GET("/:id", r.C.FindBulletinByNum)
	br.POST("/", r.C.SaveBulletin)

}
