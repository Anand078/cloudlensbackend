package logging

import (
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func Init() {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{})
}

func RedError(message string) {
	red := color.New(color.FgRed).SprintFunc()
	Logger.Errorf(red(message))
}

func LogError(message string) {
	Logger.Println("ERROR:", message)
}
