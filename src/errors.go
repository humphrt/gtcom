package main

import (
	"fmt"
	"os"
)

func CheckError(err error, showError bool, args ...interface{}) {
	if err != nil {
		fmt.Println(args...)
		if showError {
			fmt.Println("Error: ", err)
		}
	}
}

func CheckFatalError(err error, showError bool, args ...interface{}) {
	if err != nil {
		fmt.Println(args...)
		if showError {
			fmt.Println("Error: ", err)
		}
	}
	os.Exit(1)
}
