package services

import (
	"context"
	"github.com/RobertMaulana/x-user-service/proto/comment"
	"github.com/RobertMaulana/x-user-service/proto/common"
	"github.com/RobertMaulana/x-comment-service/logger"
	"time"
)

var (
	CommentService commentServiceInterface = &commentService{}
)

type commentService struct {
}

type commentServiceInterface interface {
	GetOrganizationMembers(comment.CommentsClient, *comment.OrganizationNameRequest) (*common.Response, error)
}

func (s *commentService) GetOrganizationMembers(client comment.CommentsClient, data *comment.OrganizationNameRequest) (*common.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.GetOrganizationMembers(ctx, data)
	if err != nil {
		logger.Info("unable to get data from other services")
		return nil, err
	}
	return resp, nil
}
