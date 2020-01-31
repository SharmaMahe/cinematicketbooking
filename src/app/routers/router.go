package routers

import (
	"app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // Indicate MainController.Post method to handle POST requests.
	beego.Router("/booking", &controllers.MainController{}, "post:BookSeat")

	// Indicate MainController.Post method to handle POST requests.
	beego.Router("/cancel-booking", &controllers.MainController{}, "post:CancelSeat")
}
