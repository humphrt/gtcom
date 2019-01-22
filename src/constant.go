package main

import "github.com/mkideal/cli"

var tagsMajor = []string{
	"RELEASE",
}

var tagsMinor = []string{
	"FEATURE",
}

var tagsFix = []string{
	"BUG FIX",
}

var tagsNoinc = []string{
	"DEV",
	"CONFIG",
	"UNIT TEST",
	"CLEANUP",
	"INFRA",
	"DOC",
}

type ArgT struct {
	cli.Helper
	Tag     string `cli:"t,tag" usage:"commit tag:\t- RELEASE (inc. X.y.z)\n\t\t\t\t- FEATURE (inc. x.Y.z)\n\t\t\t\t- BUGFIX (inc. x.y.Z)\n\t\t\t\t- DEV, CONFIG, UNIT TEST, CLEANUP, INFRA, DOC (no inc.)\n\t\t (default: DEV)"`
	Message string `cli:"*m,message" usage:"commit message"`
}

type Version struct {
	MajorX int
	MinorY int
	FixZ   int
}
