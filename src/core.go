package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

//Core main loop of the program
func Core(data string, argv *ArgT, rootPath string) {
	currVersion := IncrementVersion(argv.Tag, data)
	newVersion := strconv.Itoa(currVersion.MajorX) + "." + strconv.Itoa(currVersion.MinorY) + "." + strconv.Itoa(currVersion.FixZ)
	commitMessage := "[v" + newVersion + "][" + argv.Tag + "] " + argv.Message

	createFileWithContent(rootPath+"/VERSION", newVersion)
	out, err := ExecutionGit(commitMessage, data, rootPath)
	if err != nil {
		createFileWithContent(rootPath+"/VERSION", data)
	}
	fmt.Printf(string(out))
}

//IncrementVersion increment the version in function of the tag
func IncrementVersion(tag string, data string) Version {
	version := GetVersion(data)
	switch {
	case isInArray(tag, tagsMajor):
		version.MajorX++
		version.MinorY = 0
		version.FixZ = 0
	case isInArray(tag, tagsMinor):
		version.MinorY++
		version.FixZ = 0
	case isInArray(tag, tagsFix):
		version.FixZ++
	case isInArray(tag, tagsNoinc):
	default:
		CheckFatalError(errors.New("commit tag { "+tag+" } doesn't exist !"), true)
	}
	return version
}

//GetVersion get the current version of the file version
func GetVersion(data string) Version {
	var (
		vers Version
		err  error
	)

	s := strings.Split(data, ".")
	vers.MajorX, err = strconv.Atoi(s[0])
	CheckFatalError(err, false, "Bad format in VERSION file")

	vers.MinorY, err = strconv.Atoi(s[1])
	CheckFatalError(err, false, "Bad format in VERSION file")

	vers.FixZ, err = strconv.Atoi(s[2])
	CheckFatalError(err, false, "Bad format in VERSION file")

	return vers
}

//ExecutionGit return the execution of git command
func ExecutionGit(commitMessage, data, rootPath string) ([]byte, error) {
	fmt.Println(rootPath)
	exec.Command("sh", "-c", "git add "+rootPath+"/VERSION").Output()
	return exec.Command("sh", "-c", "git commit -m \""+commitMessage+"\"").Output()
}
