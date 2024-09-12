package models

type User struct {
	ID    int    `json:"id"`
	FName string `json:"fname"`
	LName string `json:"lname"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
