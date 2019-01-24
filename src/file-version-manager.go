package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

//createFileWithContent Create the file, with content
func createFileWithContent(pathfile string, content string) {
	fmt.Println(pathfile)
	f, err := os.Create(pathfile)
	CheckFatalError(err, true)

	f.WriteString(content)

	err = f.Close()
	CheckFatalError(err, true)
}

//GetVersionFileContent get data of file version
func GetVersionFileContent(pathfile string) (string, error) {
	data, err := ioutil.ReadFile(pathfile)
	return string(data), err
}

//ParseFile parse the file version
func ParseFile(data string) error {
	r, _ := regexp.Compile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
	valid := r.MatchString(data)
	if !valid {
		return errors.New("error")
	}
	return nil
}
