package hello

import (
    "github.com/sirupsen/logrus"
)

func Greet() string {
    logrus.Infoln("Greeting from hello")
    return "hello"
}
