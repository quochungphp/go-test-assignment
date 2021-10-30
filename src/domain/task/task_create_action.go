package task

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/pkg/errors"
	"github.com/twinj/uuid"
)

type TaskCreateAction struct {
	Db *pg.DB
}

func (T TaskCreateAction) Execute(Content string, UserID string) (taskDetail Task, err error) {
	t := time.Now()
	currentDate := t.Format("2006-01-02")
	taskDetail = Task{ID: uuid.NewV4().String(), Content: Content, UserId: UserID, CreatedDate: currentDate}

	count, err := T.Db.Model(new(Task)).Where("user_id = ?", UserID).Where("created_date = ?", currentDate).Count()
	if err != nil {
		return Task{}, errors.Wrapf(err, "Create a task error")
	}

	if count >= 5 {
		return Task{}, errors.Errorf("Task over 5 tasks, can not create")
	}

	_, err = T.Db.Model(&taskDetail).Insert()
	if err != nil {
		return Task{}, errors.Wrapf(err, "Create a task error")
	}
	return taskDetail, nil
}
