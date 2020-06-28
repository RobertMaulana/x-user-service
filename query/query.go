package query

const (
	GetAllOrganizationMembers = `
		SELECT login, avatar, followers, following
		FROM users
		WHERE organization_id = $1 ORDER BY followers DESC;
	`
)
