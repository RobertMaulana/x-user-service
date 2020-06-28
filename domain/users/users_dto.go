package users

type ApiGeneralResponse struct {
	Status int `json:"status"`
	Data []OrganizationMembers `json:"data"`
	Message string `json:"message"`
}

type OrganizationMembers struct {
	Login string `json:"login"`
	Avatar string `json:"avatar"`
	Followers int64 `json:"followers"`
	Following int64 `json:"following"`
}

type Organization struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}
