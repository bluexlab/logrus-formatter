package formatter

import (
	"os"

	"github.com/sirupsen/logrus"
)

// InitLogger setup formatter for the default logger
func InitLogger() {
	formatter := Formatter{}
	formatter.HideKeys = true
	formatter.ShowFullLevel = true
	formatter.NoColors = true
	formatter.NoFieldsColors = true

	logrus.SetFormatter(&formatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.TraceLevel)
}
