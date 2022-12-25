package logic

import (
	"backend/database"
)

type Logic struct {
	//DB *sql.DB
	Database *database.Database
}

func (l *Logic) ReturnLogic() ([]database.User, error) {
	users, err := l.Database.GetUser()

	if err != nil {
		//bad
	}

	return users, nil
}
