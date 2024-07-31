package server

import (
	"github.com/Turgho/go-opportunities.git/router"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	router.InitializeRoutes(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
