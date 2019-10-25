package routers

import (
	"github.com/PeterWangYong/todo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/task", &controllers.TaskController{}, "get:List;post:Create")
    beego.Router("/task/:id", &controllers.TaskController{}, "put:Update")
}
