package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func createFileWithContent(filename string, content string) int {
	f, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return 1
	}
	f.WriteString(content)

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}

func getVersionFileContent() string {
	data, err := ioutil.ReadFile("VERSION")
	CheckFatalError(err, true, "File reading error:")
	return string(data)
}

func ParseFile(data string) error {
	r, _ := regexp.Compile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
	valid := r.MatchString(data)
	if !valid {
		return errors.New("error")
	}
	return nil
}
