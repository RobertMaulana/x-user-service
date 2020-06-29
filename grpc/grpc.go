package grpc

import (
	"fmt"
	"github.com/RobertMaulana/x-user-service/proto/comment"
	"google.golang.org/grpc"
	"log"
	"os"
)

func CommentService() comment.CommentsClient {
	port := os.Getenv("COMMENT_SERVICE")
	fmt.Printf("port comment service %#v", port)
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect to", port, err)
	}
	fmt.Println(fmt.Sprintf("Grpc running at port %s...", port))
	return comment.NewCommentsClient(conn)
}
