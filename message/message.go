/*
Package message ...
*/
package message

import (
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

// Build a message given details as strings.
func BuildMessage(id string, text string, details ...string) string {

	var resultStruct Message

	// Construct details.

	detailMap := make(map[string]interface{})
	for index, value := range details {
		detailMap[strconv.Itoa(index+1)] = value
	}

	// Construct entire message.

	if len(details) > 0 {
		resultStruct = Message{
			Id:      id,
			Text:    text,
			Details: detailMap,
		}
	} else {
		resultStruct = Message{
			Id:   id,
			Text: text,
		}
	}

	// Convert to JSON.

	result, _ := json.Marshal(resultStruct)
	return string(result)
}

// Build a message given details as a map of strings.
func BuildMessageUsingMap(id string, text string, details map[string]string) string {

	// Construct entire message.

	resultStruct := Message{
		Id:      id,
		Text:    text,
		Details: details,
	}

	result, _ := json.Marshal(resultStruct)
	return string(result)
}
