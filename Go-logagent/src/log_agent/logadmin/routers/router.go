package routers

import (
	"github.com/astaxie/beego"
	"logadmin/controllers"
	"logadmin/controllers/AppController"
	"logadmin/controllers/LogController"
)

func init() {
    beego.Router("/ljw", &controllers.MainController{})
	beego.Router("/index", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/list", &AppController.AppController{}, "*:AppList")
	beego.Router("/app/apply", &AppController.AppController{}, "*:AppApply")

	beego.Router("/app/create", &AppController.AppController{}, "*:AppCreate")

	beego.Router("/log/apply", &LogController.LogController{}, "*:LogApply")
	beego.Router("/log/list", &LogController.LogController{}, "*:LogList")
	beego.Router("/log/create", &LogController.LogController{}, "*:LogCreate")

}
