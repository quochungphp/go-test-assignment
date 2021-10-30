package user

type User struct {
	ID       string `sql:"primary_key;auto_increment" json:"id"`
	Password string `sql:"omitempty" json:"password"`
	MaxTodo  string `sql:"size:100;not null;" json:"maxTodo"`
}
