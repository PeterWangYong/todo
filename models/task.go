package models

// 导入相应数据库驱动，sql.Open时会寻找相关信息
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	)

// Task 任务数据模型
// 默认int类型的Id被自动声明为自增长类型
type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Done bool `json:"done"`
}

func NewTask(title string) *Task {
	task := Task{
		Title: title,
		Done:  false,
	}
	return &task
}

func init() {
	// 注册驱动，sqlite驱动默认已经注册
	//orm.RegisterDriver()
	// 注册default数据库，名称必须为default便于orm.NewOrm()时调用
	orm.RegisterDataBase("default", "sqlite3", "todo.db")
	// 注册数据模型
	orm.RegisterModel(new(Task))
	// 根据数据模型生成数据表
	orm.RunSyncdb("default", false, true)
}