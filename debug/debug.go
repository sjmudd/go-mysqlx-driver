package debug

import (
	"fmt"
	"os"
)

/*

Use various environment variables to determine verbose logging
- DEBUG
- DEBUG_MSG protobuf messages only
*/

// Msg writes a debug message to stdout iif DEBUG environment variable is set
func Msg(format string, a ...interface{}) (n int, err error) {
	if os.Getenv("DEBUG") == "" {
		return 0, nil
	}
	return fmt.Printf("DEBUG: "+format+"\n", a...)
}

// MsgProtobuf writes a debug message iif environment variable DEBUG_MSG is set
func MsgProtobuf(format string, a ...interface{}) (n int, err error) {
	if os.Getenv("DEBUG_MSG") == "" {
		return 0, nil
	}
	return fmt.Printf("MSG: "+format+"\n", a...)
}
