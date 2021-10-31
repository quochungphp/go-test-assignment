package user

import "time"

type Users struct {
	ID          string    `sql:"primary_key;auto_increment" json:"id"`
	Username    string    `sql:"size:50;not null;" json:"username"`
	Password    string    `sql:"omitempty" json:"password"`
	MaxTodo     string    `sql:"int;not null;" json:"maxTodo"`
	CreatedDate time.Time `sql:"default:CURRENT_TIMESTAMP" json:"created_date"`
}
