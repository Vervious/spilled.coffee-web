package designer

import (
	"bytes"
	"strconv"
)

func SimpleView(title string, requestedGetVars map[string]string, requestedPostVars map[string]string, b *bytes.Buffer) *bytes.Buffer {
	s := b.String()

	// create text boxes
	textBoxes := ""
	if len(requestedGetVars) > 0 {
		textBoxes += "<form method=\"get\">\n"
		for k, v := range requestedGetVars {
			textBoxes += k + ": <input type='text' name='" + k + "' value='" + v + "'><br>\n"
		}
		textBoxes += "<input type='submit' value='Submit'>\n</form>"
	}
	if len(requestedPostVars) > 0 {
		textBoxes += "<form method=\"post\">\n"
		for k, v := range requestedPostVars {
			textBoxes += k + ": <input type='text' name='" + k + "' value='" + v + "'><br>\n"
		}
		textBoxes += "<input type='submit' value='Submit'>\n</form>"
	}

	// do some inefficient concatenation, for now
	s = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>` + title + `</title>
</head>
<body>
<div style="font-size: 1.5em; padding-left: 1.5em;">
` + textBoxes + `
</div>
<p style="font-size: 1.5em; padding-left: 1.5em; border-left: 2px solid gray;">
` + s + `
</p>
</body>
</html>
`
	s = `Content-Type: text/html
Content-Length: ` + strconv.Itoa(len(s)) + "\n\n" + s

	// mutate b (for now), may want to consider immutables
	b.Reset()
	b.WriteString(s)
	return b
}
