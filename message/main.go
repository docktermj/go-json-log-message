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
	Level   string      `json:"level,omitempty"`
	Text    string      `json:"text,omitempty"`
	Details interface{} `json:"details,omitempty"`
}
