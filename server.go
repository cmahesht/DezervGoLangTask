package main

import (
	"DezervGoLangTask/api"
	"DezervGoLangTask/helpers/confighelper"
	"DezervGoLangTask/helpers/databasehelper"
	"DezervGoLangTask/helpers/logginghelper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

/*main : Author - GaneshL
Purpose : main program execution started
*/
func main() {
	startServer()
}

/*startServer : Author - GaneshL
Purpose : Start go lang echo server
*/
func startServer() {

	e := echo.New()

	// Initilisation of config helper
	confighelper.InitViper()

	// Middlewares Added
	e.Use(middleware.Logger())

	e.HideBanner = true

	// Api route initilisation
	api.Init(e)

	// mongo driver initilisation
	sessionCreationErr := databasehelper.InitMongo()
	if sessionCreationErr != nil {
		logginghelper.Error("ERROR : Error occurred in InitMongo from mongodb : mgo.Dial ")
	}

	// Get Server Port From Config
	serverPort := confighelper.GetConfig("serverPort")

	// Started echo server
	err := e.Start(":" + serverPort)
	if err != nil {
		logginghelper.Error("Unable to start server on :" + serverPort)
	}

}
