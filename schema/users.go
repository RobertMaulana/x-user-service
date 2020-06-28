package schema

type Users struct {
	Base
	Login	string
	Avatar	string
	Followers	int
	Following	int
	OrganizationId int
}

func (Users) TableName() string {
	return "users"
}

func (Users) Pk() string {
	return "id"
}

func (t Users) Ref() string {
	return t.TableName() + "(" + t.Pk() + ")"
}

func (t Users) AddForeignKeys() {
}

func (t Users) InsertDefaults() {
	Database.Exec(`INSERT INTO users (login, avatar, followers, following, organization_id)
	SELECT * FROM (SELECT 'robert', 'https://google.com', 1500, 500, 1) AS tmp
		WHERE NOT EXISTS (
    		SELECT login FROM users WHERE login = 'robert'
	) LIMIT 1;`)
	Database.Exec(`INSERT INTO users (login, avatar, followers, following, organization_id)
	SELECT * FROM (SELECT 'maulana', 'https://google.com1', 324, 44, 1) AS tmp
		WHERE NOT EXISTS (
    		SELECT login FROM users WHERE login = 'maulana'
	) LIMIT 1;`)
	Database.Exec(`INSERT INTO users (login, avatar, followers, following, organization_id)
	SELECT * FROM (SELECT 'andi', 'https://google.com1', 324, 44, 2) AS tmp
		WHERE NOT EXISTS (
    		SELECT login FROM users WHERE login = 'andi'
	) LIMIT 1;`)
	Database.Exec(`INSERT INTO users (login, avatar, followers, following, organization_id)
	SELECT * FROM (SELECT 'rudi', 'https://google.com3', 2423, 344, 2) AS tmp
		WHERE NOT EXISTS (
    		SELECT login FROM users WHERE login = 'rudi'
	) LIMIT 1;`)
}
