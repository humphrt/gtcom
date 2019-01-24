package main

import (
	"fmt"
	"os"
)

//CheckError check if the err error is nil,
//show the error if the boolean ShowError is true. If arguments(args) are nil the error (err) is print automaticly
//show arguments (args) if you need to print some messages
func CheckError(err error, showError bool, args ...interface{}) {
	if err != nil {
		if args == nil {
			fmt.Println("Error : ", err)
		} else {
			fmt.Println(args...)
			if showError {
				fmt.Println("Error: ", err)
			}
		}
	}
}

//CheckFatalError check if the err error is nil overwrise exit 1,
//show the error if the boolean ShowError is true. If arguments(args) are nil the error (err) is print automaticly
//show arguments (args) if you need to print some messages
func CheckFatalError(err error, showError bool, args ...interface{}) {
	if err != nil {
		if args == nil {
			fmt.Println("Error : ", err)
		} else {
			fmt.Println(args...)
			if showError {
				fmt.Println("Error: ", err)
			}
		}
		os.Exit(1)
	}
}
