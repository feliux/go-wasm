//go:build wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"
)

// var c chan bool

// func init() {
// 	c = make(chan bool)
// }

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("formatJSON", jsonWrapper())
	// Ensure that the Go program is running when JavaScript calls it
	<-make(chan struct{})
	// <-c
}

func prettyJson(input string) (string, error) {
	var raw any
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	// c <- true
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	// Func is a wrapped Go function that can be called by JavaScript. The FuncOf function can be used to create a Func type.
	// The function which is passed to FuncOf will be called synchronously from Javascript. The first parameter of this function
	// is Javascript’s 'this' keyword. 'this' refers to JavaScript’s global object. The second parameter is a slice of []js.Value which
	// represents the arguments that will be passed to the Javascript function call
	jsonfunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result := map[string]any{
				"error": "Invalid no of arguments passed",
			}
			return result
		}
		// Get the document property of JavaScript from the global scope
		jsDoc := js.Global().Get("document")
		// The Truthy function in line no. 7 is JavaScript’s way of testing for nil. If truthy returns false, it means the property doesn’t exist.
		// Hence the appropriate error string is returned to JavaScript. We do not explicitly return a Go error type
		if !jsDoc.Truthy() {
			result := map[string]any{
				"error": "Unable to get document object",
			}
			return result
		}
		jsonOuputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOuputTextArea.Truthy() {
			result := map[string]any{
				"error": "Unable to get output text area",
			}
			return result
		}
		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			errStr := fmt.Sprintf("unable to parse JSON. Error %s occurred\n", err)
			result := map[string]any{
				"error": errStr,
			}
			return result
		}
		// Use the Set method to set the value property of the jsonoutput text area to the formatted JSON string
		jsonOuputTextArea.Set("value", pretty)
		return nil
	})
	return jsonfunc
}
