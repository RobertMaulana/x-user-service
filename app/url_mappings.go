package app

import (
	"github.com/RobertMaulana/x-user-service/controllers/ping"
	"github.com/RobertMaulana/x-user-service/controllers/users"
)

func mapUrls() {
	// Kubernetes need to check health pod using this open public endpoint
	router.GET("/ping", ping.Ping)

	// Rest route
	xendit := router.Group("/orgs")
	{
		xendit.GET("/:organization_name/members", users.GetOrganizationMembers)
	}
}