package repositories

import (
	"github.com/jinzhu/gorm"
	"iris/datamodels"
)

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type TaskRepository interface {
	GetAllTaskList() []datamodels.BizTask
}

func NewTaskDBRep(source *gorm.DB) TaskRepository {
	source = source.Table("BIZ_TASK")
	return &taskSQLRepository{source: source}
}

type taskSQLRepository struct {
	source *gorm.DB
}

func (r *taskSQLRepository) GetAllTaskList() (bizTask []datamodels.BizTask) {
	qc := r.source.Model(&datamodels.BizTask{})
	qc.Find(&bizTask)
	return bizTask
}
