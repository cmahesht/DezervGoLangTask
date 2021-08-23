package api

import (
	"DezervGoLangTask/api/modules/login"

	"github.com/labstack/echo"
)

/*API : Author - Mahesh Chinvar
Purpose : Init api group
*/

//Init api binding
func Init(e *echo.Echo) {

	// group for open rest api
	o := e.Group("/o")

	// group for resticted JWT Auth rest api
	r := e.Group("/r")

	// meeting api registration
	login.Init(o, r)
}
