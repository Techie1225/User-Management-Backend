package repository

import (
	"fmt"
	"user-management-backend/internal/database"
	"user-management-backend/internal/models"

	"github.com/Masterminds/squirrel"
)

var sqlBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question) // Use `?` for MySQL

func GetUsers() ([]models.User, error) {
	if database.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	query, args, err := sqlBuilder.Select("id", "fname", "lname", "email", "age").From("users").ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.FName, &user.LName, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUserByID(id int) (models.User, error) {
	if database.DB == nil {
		return models.User{}, fmt.Errorf("database connection is not initialized")
	}

	query, args, err := sqlBuilder.Select("id", "fname", "lname", "email", "age").From("users").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return models.User{}, err
	}

	row := database.DB.QueryRow(query, args...)
	var user models.User
	err = row.Scan(&user.ID, &user.FName, &user.LName, &user.Email, &user.Age)
	return user, err
}

func AddUser(user models.User) error {
	if database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query, args, err := sqlBuilder.Insert("users").Columns("fname", "lname", "email", "age").Values(user.FName, user.LName, user.Email, user.Age).ToSql()
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(query, args...)
	return err
}

func UpdateUser(user models.User) error {
	if database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query, args, err := sqlBuilder.Update("users").
		Set("fname", user.FName).
		Set("lname", user.LName).
		Set("email", user.Email).
		Set("age", user.Age).
		Where(squirrel.Eq{"id": user.ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(query, args...)
	return err
}

func DeleteUser(id int) error {
	if database.DB == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	query, args, err := sqlBuilder.Delete("users").Where(squirrel.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}

	_, err = database.DB.Exec(query, args...)
	return err
}
