package repositories

import (
	"errors"
	"github.com/jinzhu/gorm"
	"iris/commons"
	"iris/datamodels"
	"iris/datasource"
	"strconv"
	"time"
)

// UserRepository handles the basic operations of a user entity/model.
// It's an interface in order to be testable, i.e a memory user repository or
// a connected to an sql database.
type TaskRepository interface {
	GetAllTaskList() commons.Page
	GetTaskById(taskId int) datamodels.BizTask
	ScrambleTask(userId int64, taskId int) (bool, error)
	GetTaskListByMapId(mapId int) (count int, list []datamodels.BizTask)
	DoPersonGetTask(userId int, chId int, taskId int) (bool, error)
}

type taskSQLRepository struct {
	source *gorm.DB
}

func NewTaskDBRep() TaskRepository {
	return &taskSQLRepository{source: datasource.DB}
}

func (r *taskSQLRepository) GetAllTaskList() commons.Page {
	qc := r.source.Table("BIZ_TASK").Model(&datamodels.BizTask{})
	var bizTask []datamodels.BizTask
	qc.Find(&bizTask)
	page := commons.Page{
		Data:       bizTask,
		PageNo:     1,
		PageSize:   commons.PageSize,
		TotalCount: len(bizTask),
		TotalPage:  (len(bizTask) / commons.PageSize) + 1,
	}
	return page
}

func (r *taskSQLRepository) GetTaskById(taskId int) (task datamodels.BizTask) {
	qc := r.source.Table("BIZ_TASK").Model(&datamodels.BizTask{})
	qc.Where("ID = ?", taskId).Find(&task)
	return task
}

func (r *taskSQLRepository) ScrambleTask(userId int64, taskId int) (bool, error) {
	task := r.GetTaskById(taskId)
	bizUserTask := datamodels.BizUserTask{
		UserId:     int(userId),
		TaskId:     taskId,
		Coin:       task.Coin,
		Experience: task.Experience,
		Honor:      task.Honor,
		CreateUser: strconv.FormatInt(userId, 10),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	r.source.Table("BIZ_USER_TASK").Create(&bizUserTask)
	bizTask := datamodels.BizTask{}
	if r.source.Table("BIZ_TASK").Where("ID = ? ", taskId).First(&bizTask).RecordNotFound() {
		return false, errors.New("任务不存在")
	}
	r.source.Table("BIZ_TASK").Model(&datamodels.BizTask{}).Where("ID = ?", taskId).Update(datamodels.BizTask{Status: 1, UpdateUser: strconv.FormatInt(userId, 10),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05")})
	return true, nil
}
func (r *taskSQLRepository) DoPersonGetTask(userId int, chId int, taskId int) (bool, error) {
	task := r.GetTaskById(taskId)
	bizChTask := datamodels.BizChTask{
		UserId:     userId,
		ChId:       chId,
		TaskId:     taskId,
		Coin:       task.Coin,
		Experience: task.Experience,
		Honor:      task.Honor,
		CreateUser: strconv.Itoa(chId),
		CreateTime: commons.GetNowStr(),
	}
	r.source.Table("BIZ_CH_TASK").Create(&bizChTask)
	bizTask := datamodels.BizTask{}
	if r.source.Table("BIZ_TASK").Where("ID = ? ", taskId).First(&bizTask).RecordNotFound() {
		return false, errors.New("任务不存在")
	}
	r.source.Table("BIZ_TASK").Model(&datamodels.BizTask{}).Where("ID = ?", taskId).Update(datamodels.BizTask{Status: 1, UpdateUser: strconv.Itoa(chId),
		UpdateTime: commons.GetNowStr()})
	return true, nil

}
func (r *taskSQLRepository) GetTaskListByMapId(mapId int) (count int, list []datamodels.BizTask) {
	qc := r.source.Table("BIZ_TASK").Model(&datamodels.BizTask{})
	qc.Where("MAP_ID = ?", mapId).Count(&count)
	if count > 0 {
		qc.Where("MAP_ID = ?", mapId).Find(&list)
	}
	return count, list
}
