package testdata

import (
	"fmt"
	"testing"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error: %d - %s", e.Code, e.Message)
}

type SpecialError struct {
	Code    int
	Message string
}

func (e *SpecialError) Error() string {
	return fmt.Sprintf("Error: %d - %s", e.Code, e.Message)
}

func doSomething1() error {
	// Simulating an error
	return &SpecialError{
		Code:    500,
		Message: "Something went wrong",
	}
}
func TestERrr(t *testing.T) {
	err := doSomething1()
	if err != nil {
		switch errType := err.(type) {
		case *CustomError:
			fmt.Println("Custom Error:", errType.Code, "-", errType.Message)
		case *SpecialError:
			fmt.Println("special")

		default:
			fmt.Println("Unknown Error:", err)
		}
	}
}
