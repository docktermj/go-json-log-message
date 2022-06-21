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

	// Create a simple message.

	aMessage := message.BuildMessage("A", "B", "C", "D", programName, buildVersion, buildIteration)
	fmt.Println(aMessage)

	// Create a message with a dictionary of details.

	detailsMap := map[string]string{
		"programName":    programName,
		"buildVersion":   buildVersion,
		"buildIteration": buildIteration,
	}
	anotherMessage := message.BuildMessageUsingMap("Unique-id", "Error level", "Error message", detailsMap)
	fmt.Println(anotherMessage)

	// Parse a message.

	parsedMessage := message.ParseMessage(anotherMessage)
	details, ok := parsedMessage.Details.(map[string]interface{})
	if !ok {
		fmt.Printf("Error: %t", ok)
	}

	fmt.Println(parsedMessage.Id)
	fmt.Println(parsedMessage.Level)
	fmt.Println(parsedMessage.Text)
	fmt.Println(details["buildVersion"])

}
