package markdown

import (
	"bufio"
	"os"
	"strings"
)

func parseFile(file *os.File) (*Document, error) {
	scanner := bufio.NewScanner(file)

	var document Document
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			continue
		}

		var node Node
		firstChar := line[0]
		switch firstChar {
		case '#':
			node = parseHeading(line)
		default:
			// Note: markdown specification treats consequitive non-empty lines
			// as a single paragraph. In my implementation single paragraph
			// is represented by single line. Maybe I'll reconsider later.
			node = parseParagraph(line)
		}
		document.Children = append(document.Children, node)
	}

	return &document, nil
}

func parseHeading(line string) Node {
	level := 0
	contentIndex := 0
	for i, char := range line {
		if char == '#' {
			level++
		} else {
			contentIndex = i
			break
		}
	}
	content := strings.TrimSpace(line[contentIndex:])
	return &Heading{Level: level, Content: content}
}

func parseParagraph(line string) Node {
	return &Paragraph{Content: line}
}
