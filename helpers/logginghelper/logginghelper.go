package logginghelper

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

/*LoggingHelper : Author - Mahesh Chinvar
Purpose : Logging
*/

var Logger *logrus.Logger

// Logger initilisation
func init() {
	Logger = logrus.New()
	// open a file
	f, err := os.OpenFile("DezervGoLangTask.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}

	// Log as JSON instead of the default ASCII formatter.
	Logger.Formatter = &log.JSONFormatter{}

	// Output to stderr instead of stdout, could also be a file.
	Logger.Out = f

	// Only log the warning severity or above.
	Logger.Level = log.DebugLevel

}

//Info to print info
func Info(args ...interface{}) {
	Logger.Info("", args)
}

//Error to print Error
func Error(args ...interface{}) {
	Logger.Error("", args)
}

//Warn to print Warn
func Warn(args ...interface{}) {
	Logger.Warn("", args)
}

//Debug to print Debug
func Debug(args ...interface{}) {
	Logger.Debug("", args)
}

//Panic to print Debug
func Panic(args ...interface{}) {
	Logger.Panic("", args)
}
