package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	Host = "localhost:8081"
)


func TestGetOrganizationMembers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, err := http.NewRequest("GET", Host + "/orgs/xendit/members", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp := httptest.NewRecorder()
	assert.Equal(t, resp.Code, 200)
}

