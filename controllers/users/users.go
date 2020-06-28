package users

import (
	"encoding/json"
	"github.com/RobertMaulana/x-user-service/domain/users"
	"github.com/RobertMaulana/x-user-service/grpc"
	"github.com/RobertMaulana/x-user-service/proto/comment"
	"github.com/RobertMaulana/x-user-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrganizationMembers(c *gin.Context) {
	organizationName := c.Param("organization_name")
	// Get Organization Id from comment-service
	resp, err := services.CommentService.GetOrganizationMembers(grpc.CommentService(), &comment.OrganizationNameRequest{Name: organizationName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var organization users.Organization
	errMarshal := json.Unmarshal(resp.Data, &organization)
	if errMarshal != nil {
		c.JSON(http.StatusInternalServerError, errMarshal)
		return
	}
	members, errGetMembers := services.UserService.GetMembers(organization.Id)
	if errGetMembers != nil {
		c.JSON(errGetMembers.Status, err)
		return
	}
	c.JSON(http.StatusOK, members)
}