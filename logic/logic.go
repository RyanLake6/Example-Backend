package logic

import (
	"backend/database"
	"errors"
)

type Logic struct {
	//DB *sql.DB
	Database *database.Database
}

func (l *Logic) ReturnLogic() ([]database.User, error) {
	// return nil, errors.New("this is a test error")
	users, err := l.Database.GetUser()

	if err != nil {
		return nil, errors.New("Error has occured")
	}

	return users, nil
}
