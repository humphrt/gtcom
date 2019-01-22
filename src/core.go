package main

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func Core(data string, argv *ArgT) {
	currVersion := IncrementVersion(argv.Tag, data)
	newVersion := strconv.Itoa(currVersion.MajorX) + "." + strconv.Itoa(currVersion.MinorY) + "." + strconv.Itoa(currVersion.FixZ)
	commitMessage := "[v" + newVersion + "][" + argv.Tag + "] " + argv.Message
	createFileWithContent("VERSION", newVersion)
	exec.Command("sh", "-c", "git add VERSION").Output()
	out, err := exec.Command("sh", "-c", "git commit -m \""+commitMessage+"\"").Output()
	if err != nil {
		createFileWithContent("VERSION", data)
	}
	fmt.Printf("%s", out)
}

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
	default:
		CheckFatalError(errors.New(""), false, "ERROR: commit tag {", tag, "} doesn't exist")
	}
	return version
}

func GetVersion(data string) Version {
	var (
		vers Version
		err  error
	)

	s := strings.Split(data, ".")
	vers.MajorX, err = strconv.Atoi(s[0])
	CheckFatalError(err, false, "ERROR: Bad format in VERSION file")

	vers.MinorY, err = strconv.Atoi(s[1])
	CheckFatalError(err, false, "ERROR: Bad format in VERSION file")

	vers.FixZ, err = strconv.Atoi(s[2])
	CheckFatalError(err, false, "ERROR: Bad format in VERSION file")

	return vers
}
