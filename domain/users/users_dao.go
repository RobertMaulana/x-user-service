package users

import (
	"database/sql"
	"fmt"
	"github.com/RobertMaulana/x-comment-service/datasource/postgre/comments_db"
	"github.com/RobertMaulana/x-comment-service/logger"
	"github.com/RobertMaulana/x-user-service/query"
	"github.com/RobertMaulana/x-comment-service/utils/errors"
	"net/http"
)

func (organization *Organization) GetOrganizationMembers() (*ApiGeneralResponse, *errors.RestErr) {
	stmt, err := comments_db.Client.Prepare(query.GetAllOrganizationMembers)
	if err != nil {
		logger.Error("error when trying to prepare get members statement", err)
		return nil, errors.InternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(organization.Id)
	if err != nil {
		logger.Error("error when trying to get members statement", err)
		return nil, errors.InternalServerError("database error")
	}
	var results []OrganizationMembers
	for rows.Next() {
		var res OrganizationMembers
		if err := rows.Scan(&res.Login, &res.Avatar, &res.Followers, &res.Following); err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.NotFound("data not found")
			}
			logger.Error("error when trying to execute members", err)
			return nil, errors.InternalServerError("database error")
		}
		results = append(results, res)
	}
	if results == nil {
		results = []OrganizationMembers{}
	}
	res := &ApiGeneralResponse{
		Status: http.StatusOK,
		Data: results,
		Message: "success",
	}
	return res, nil
}