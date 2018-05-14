package main

import (
	"bufio"
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

	requestMethod := os.Getenv("REQUEST_METHOD")
	switch requestMethod {
	case "POST":
		scanner := bufio.NewScanner(os.Stdin)
		postData := ""
		for scanner.Scan() {
			postData += scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			postData += fmt.Sprintln(err)
		}
		fmt.Fprintf(b, "post data: %v<br/>\n\n", postData)
	case "GET":
		queryString := os.Getenv("QUERY_STRING")
		fmt.Fprintf(b, "query string: %v<br/>\n\n", queryString)
	default:
		fmt.Fprintf(b, "No input received.\n")
	}

	fmt.Fprintf(b, "Server time at %s is %s\n",
		os.Getenv("SERVER_NAME"), time.Now().Format(time.RFC1123))
	b = designer.SimpleView("helloworld", requestedGetVars, requestedPostVars, b)
	// fmt.Println("Content-type: text/html")
	// fmt.Printf("Content-Length: %d\n\n", b.Len())
	b.WriteTo(os.Stdout)
}
