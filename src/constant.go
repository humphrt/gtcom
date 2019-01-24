package main

import "github.com/mkideal/cli"

//tagsMajor array of the major version
var tagsMajor = []string{
	"RELEASE",
}

//tagsMinor array of the minor version
var tagsMinor = []string{
	"FEATURE",
}

//tagsFix array of the fix version
var tagsFix = []string{
	"BUGFIX",
}

//tagsNoinc array of the no increment version
var tagsNoinc = []string{
	"DEV",
	"CONFIG",
	"UNITTEST",
	"CLEANUP",
	"INFRA",
	"DOC",
}

//ArgT argument passing in parameters
type ArgT struct {
	cli.Helper
	Tag     string `cli:"t,tag" usage:"commit tag:\t- RELEASE (inc. X.y.z)\n\t\t\t\t- FEATURE (inc. x.Y.z)\n\t\t\t\t- BUGFIX (inc. x.y.Z)\n\t\t\t\t- DEV, CONFIG, UNIT TEST, CLEANUP, INFRA, DOC (no inc.)\n\t\t (default: DEV)"`
	Message string `cli:"*m,message" usage:"commit message"`
}

//Version structure of the version
type Version struct {
	MajorX int
	MinorY int
	FixZ   int
}
