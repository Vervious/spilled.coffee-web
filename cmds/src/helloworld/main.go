package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"bob.expert/lib/designer"
)

func main() {
	var buf bytes.Buffer

	b := &buf

	// key: variable name. Value: default value.
	requestedGetVars := map[string]string{
		"exampleget": "",
	}
	requestedPostVars := map[string]string{
		"examplepost": "",
	}

	err := designer.Get(requestedGetVars, requestedPostVars)
	if err != nil {
		fmt.Fprintln(b, err)
	}

	fmt.Fprintf(b, "Parsed get vars: %v<br/>\n\n", requestedGetVars)
	fmt.Fprintf(b, "Parsed post vars: %v<br/>\n\n", requestedPostVars)

	fmt.Fprintf(b, "Server time at %s is %s\n",
		os.Getenv("SERVER_NAME"), time.Now().Format(time.RFC1123))
	b = designer.SimpleView("helloworld", requestedGetVars, requestedPostVars, b)
	// fmt.Println("Content-type: text/html")
	// fmt.Printf("Content-Length: %d\n\n", b.Len())
	b.WriteTo(os.Stdout)
}
