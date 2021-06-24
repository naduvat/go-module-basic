package world

import (
    "github.com/sirupsen/logrus"
)

func Greet() string {
    logrus.Infoln("Greeting from world")
    return "world"
}
