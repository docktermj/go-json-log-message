// The message package formats messages into a JSON string.
package message

import ()

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Detail struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Message struct {
	Id      string      `json:"id"`
	Text    string      `json:"text"`
	Details interface{} `json:"details,omitempty"`
}
