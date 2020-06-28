package app

import (
	"github.com/RobertMaulana/x-user-service/grpc"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	grpc.CommentService()
	router.Run(":8081")
}