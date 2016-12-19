package routers

import (
	"hstats/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/jt", &controllers.MainController{}, "get:Dispath")
}
