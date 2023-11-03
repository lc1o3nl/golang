package main

import (
	"fmt"
)

type Logger struct {
	instance *Logger
}

func GetInstance() *Logger {
	if Logger.instance == nil {
		Logger.instance = new(Logger)
	}
	return Logger.instance
}

func (l *Logger) Print(message string) {
	fmt.Println(message)
}

func main() {
	logger := GetInstance()

	logger.Print("Hello, world!")

	logger2 := GetInstance()

	if logger == logger2 {
		fmt.Println("Это тот же экземпляр")
	} else {
		fmt.Println("Это другой экземпляр")
	}
}
