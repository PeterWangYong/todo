package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/PeterWangYong/todo/models"
	"github.com/astaxie/beego"
	"strconv"
)

type TaskController struct {
	beego.Controller
}

// List 获取所有已完成任务
// c.ServeJSON会将c.Data["json"]中保存的内容转换为JSON字符串并写入ResponseWriter
func (c *TaskController) List() {
	tasks := models.DefaultTaskManager.ListTasks()
	c.Data["json"] = tasks
	c.ServeJSON()
}

// Create 创建一个新任务
// 客户端提交 {title: "title"} JSON格式
func (c *TaskController) Create() {
	temp := struct {
		Title string `json:"title"`
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &temp); err != nil {
		fmt.Println("empty task: ", err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	title := temp.Title
	task := models.DefaultTaskManager.CreateTask(title)
	c.Data["json"] = task
	c.ServeJSON()
}

// Update 更新一个任务的完成状态
// 客户端提交 {done: true} JSON格式
// 通过Input.Param(":id")获取动态URL参数
// 使用 strconv.ParseInt将字符串转换为int64
func (c *TaskController) Update() {
	temp := struct {
		Done bool `json:"done"`
	}{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &temp); err != nil {
		fmt.Println("empty task: ", err)
		c.Ctx.Output.SetStatus(400)
		return
	}
	id := c.Ctx.Input.Param(":id")
	intid, _ := strconv.ParseInt(id, 10, 64)
	task := models.DefaultTaskManager.UpdateTask(int(intid), temp.Done)
	c.Data["json"] = task
	c.ServeJSON()
}