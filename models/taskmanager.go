package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

// 任务管理类，用于执行数据CRUD
type TaskManager struct {}

// 全局唯一的TaskManager对象
var DefaultTaskManager TaskManager

// ListTasks 获取所有任务
// o.QueryTable("tableName").Filter().All()类似Django的objects
func (m *TaskManager) ListTasks() []*Task {
	o := orm.NewOrm()
	var tasks []*Task
	_, err := o.QueryTable("task").Filter("done", false).All(&tasks)
	if err != nil {
		fmt.Println(err)
	}
	return tasks
}

// CreateTask 创建一个任务
func (m *TaskManager) CreateTask(title string) *Task {
	task := NewTask(title)
	o := orm.NewOrm()
	_, err := o.Insert(task)
	if err != nil {
		fmt.Println(err)
	}
	return task
}

// UpdateTask 更新一个任务
func (m *TaskManager) UpdateTask(id int, done bool) *Task {
	o := orm.NewOrm()
	task := &Task{Id: id}
	err := o.Read(task)
	if err != nil {
		fmt.Println(err)
	}
	task.Done = done
	_, err = o.Update(task, "done")
	if err != nil {
		fmt.Println(err)
	}
	return task
}
