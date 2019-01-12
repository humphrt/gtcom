package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"io/ioutil"
	"strconv"
	"strings"
	"github.com/mkideal/cli"
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
    if err != nil {
        fmt.Println("File reading error", err)
        return ""
    }
		return string(data)
}

type argT struct {
	cli.Helper
	Tag 				string `cli:"t,tag" usage:"commit tag:\t- RELEASE (inc. X.y.z)\n\t\t\t\t- FEATURE (inc. x.Y.z)\n\t\t\t\t- BUGFIX (inc. x.y.Z)\n\t\t\t\t- DEV, CONFIG, UNIT TEST, CLEANUP, INFRA, DOC (no inc.)\n\t\t (default: DEV)"`
	Message    	string `cli:"*m,message" usage:"commit message"`
}

var tags_major = []string {
	"RELEASE",
}

var tags_minor = []string {
	"FEATURE",
}

var tags_fix = []string {
	"BUGFIX",
}

var tags_noinc = []string {
	"DEV",
	"CONFIG",
	"UNIT TEST",
	"CLEANUP",
	"INFRA",
	"DOC",
}

func isInArray(str string, arr []string) bool {
	for _, element := range arr {
		if str == element {
			return true
		}
	}
	return false
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)

		if argv.Tag == "" {
			argv.Tag = "DEV"
		}

		if _, err := os.Stat("./VERSION"); os.IsNotExist(err) {
			createFileWithContent("VERSION", "0.0.0")
		}
	
		data := getVersionFileContent()	
		r,  _ := regexp.Compile("^[0-9]+\\.[0-9]+\\.[0-9]+$")
	
		if ! r.MatchString(data) {
			fmt.Println("ERROR: Bad format in VERSION file\n")
			os.Exit(1)
		}
	
		s := strings.Split(data, ".")
		major, err := strconv.Atoi(s[0])
		if err != nil {
			fmt.Println("ERROR: Bad format in VERSION file\n")
			os.Exit(1)
		}
		minor, err := strconv.Atoi(s[1])
		if err != nil {
			fmt.Println("ERROR: Bad format in VERSION file\n")
			os.Exit(1)
		}
	
		fix, err := strconv.Atoi(s[2])
		if err != nil {
			fmt.Println("ERROR: Bad format in VERSION file\n")
			os.Exit(1)
		}
	
		if isInArray(argv.Tag, tags_major) {
			major += 1
		} else if isInArray(argv.Tag, tags_minor) {
			minor += 1
		} else if isInArray(argv.Tag, tags_fix) {
			fix += 1
		} else if ! isInArray(argv.Tag, tags_noinc) {
			fmt.Println("ERROR: commit tag {", argv.Tag, "} doesn't exist\n")
			os.Exit(1)
		}
	
		new_version := strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(fix)
		commit_message := "[v" + new_version + "][" + argv.Tag + "] " + argv.Message
		createFileWithContent("VERSION", new_version)
		exec.Command("sh","-c", "git add VERSION").Output()
		out, err := exec.Command("sh","-c", "git commit -m \"" + commit_message + "\"").Output()
		if err != nil {
			createFileWithContent("VERSION", data)
		}
		fmt.Printf("%s", out)
		
		return nil
	}))
}
