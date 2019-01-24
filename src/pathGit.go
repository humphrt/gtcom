package main

import (
	"os"
	"strings"
)

//GetPath descrase x time index the path
func GetPath(path string, index int) string {
	pathList := strings.Split(path, "/")
	pathList = pathList[:len(pathList)-index]
	return strings.Join(pathList, "/")
}

//GetRootPath return the path of the repository inital call with index = 0
//if the path is empty means no repository found
func GetRootPath(path string, index int) string {
	newpath := GetPath(path, index)
	if newpath == "" {
		return newpath
	}
	_, err := os.Stat(newpath + "/.git")
	if err != nil {
		return GetRootPath(path, index+1)
	}
	return newpath
}
