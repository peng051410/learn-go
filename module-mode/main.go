package main

import (
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
)

func main() {
	logrus.Println("hello, go module mode")
	logrus.Println(uuid.NewString())
}