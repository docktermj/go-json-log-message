package main

import (
	"fmt"

	"github.com/docktermj/go-json-log-message/message"
)

// Values updated via "go install -ldflags" parameters.

var programName string = "unknown"
var buildVersion string = "0.0.0"
var buildIteration string = "0"

// ----------------------------------------------------------------------------
// Main
// ----------------------------------------------------------------------------

func main() {
	aMessage := message.BuildMessage("A", "B", "C", "D", programName, buildVersion, buildIteration)
	fmt.Println(aMessage)

	detailsMap := map[string]string{
		"programName":    programName,
		"buildVersion":   buildVersion,
		"buildIteration": buildIteration,
	}
	anotherMessage := message.BuildMessageUsingMap("E", "F", detailsMap)
	fmt.Println(anotherMessage)

}
