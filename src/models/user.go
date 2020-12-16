package models

import (
	"gensjaak/go-for-quizz/src/types"
)

// CreateUser <- Inserts an user
func CreateUser(data types.UserAccount) (bool, error) {
	sql := `
		INSERT INTO
			user_accounts (
				first_name,
				last_name,
				user_name,
				password,
				user_role_id
			)
		VALUES
			($1, $2, $3, $4, $5)
	`

	_, err := Database.Exec(sql, data.FirstName, data.LastName, data.UserName, data.Password, data.UserRoleID)

	if err != nil {
		return false, err
	}

	return true, nil
}
