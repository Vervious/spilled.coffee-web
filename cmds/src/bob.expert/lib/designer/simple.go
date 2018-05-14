package designer

import "bytes"

func SimpleView(title string, b *bytes.Buffer) *bytes.Buffer {
	s := b.String()

	// do some inefficient concatenation, for now
	s = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>` + title + `</title>
</head>
<body>
<p style="font-size: 1.5em; padding-left: 1.5em; border-left: 2px solid gray;">
` + s + `
</p>
</body>
</html>
`
	s = `
Content-Type: text/html
Content-Length: ` + len(s) + "\n\n"

	// mutate b (for now), may want to consider immutables
	b.Reset()
	b.WriteString(s)
	return b
}
