package message

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions - names begin with "Test"
// ----------------------------------------------------------------------------

func TestBuildMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "C", "D")
	fmt.Println(actual)
	assert.NotEmpty(test, actual)
}

func TestBuildMessageNoDetails(test *testing.T) {
	actual := BuildMessage("A", "B", "C")
	fmt.Println(actual)
}

func TestBuildMessageNoMessage(test *testing.T) {
	actual := BuildMessage("A", "B", "")
	fmt.Println(actual)
}

func TestBuildMessageNoLevel(test *testing.T) {
	actual := BuildMessage("A", "", "")
	fmt.Println(actual)
}

func TestBuildMessageUsingMap(test *testing.T) {
	detailsMap := map[string]string{
		"FirstVariable":  "First value",
		"SecondVariable": "Second value",
	}
	actual := BuildMessageUsingMap("A", "B", "C", detailsMap)
	fmt.Println(actual)
}
