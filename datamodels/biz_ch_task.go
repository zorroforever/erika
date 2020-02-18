package datamodels

import "time"

type BizChTask struct {
	UserId              int       `json:"userId" form:"userId" gorm:"Column:USER_ID"`
	ChId                int       `json:"chId" form:"chId" gorm:"Column:CH_ID"`
	TaskId              int       `json:"taskId" form:"taskId" gorm:"Column:TASK_ID"`
	Coin                float32   `json:"coin" form:"coin" gorm:"Column:COIN"`
	Experience          int       `json:"experience" form:"experience" gorm:"Column:EXPERIENCE"`
	Honor               int       `json:"honor" form:"honor" gorm:"Column:HONOR"`
	Log                 string    `json:"log" form:"log" gorm:"Column:LOG"`
	Remark              string    `json:"remark" form:"remark" gorm:"Column:REMARK"`
	JiraUrl             string    `json:"jiraUrl" form:"jiraUrl" gorm:"Column:JIRA_URL"`
	EstimateDevelopTime int       `json:"estimateDevelopTime" form:"estimateDevelopTime" gorm:"Column:ESTIMATE_DEVELOP_TIME"`
	StartTime           time.Time `json:"startTime" form:"startTime" gorm:"Column:START_TIME"`
	EndTime             time.Time `json:"endTime" form:"endTime" gorm:"Column:END_TIME"`
	CreateUser          string    `json:"createUser" form:"createUser" gorm:"Column:CREATE_USER"`
	CreateTime          string    `json:"createTime" form:"createTime" gorm:"Column:CREATE_TIME"`
	UpdateUser          string    `json:"updateUser" form:"updateUser" gorm:"Column:UPDATE_USER"`
	UpdateTime          string    `json:"updateTime" form:"updateTime" gorm:"Column:UPDATE_TIME"`
}
