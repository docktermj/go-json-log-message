# go-json-log-message

## Synopsis

A simple Go package that formats (error) messages as JSON.
When pretty printed, the messages have the following format:

```json
{
    "id": "a-unique-message-id",
    "text": "Explanation of the error",
    "details": {
        "FirstVariable": "First value",
        "SecondVariable": "Second value"
    }
}
```

To import the go package:

```go
import (
    :
    "github.com/docktermj/go-json-log-message/message"
)
```

For examples of use, see:

1. [main.go](main.go)
1. [message_test.go](message/message_test.go)
