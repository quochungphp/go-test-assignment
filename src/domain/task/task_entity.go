package task

type Task struct {
	ID        string `gorm:"primary_key" json:"id"`
	UserId    string `gorm:"size:100" json:"user_id"`
	Content   string `gorm:"size:1000" json:"content"`
	CreatedAt string `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
