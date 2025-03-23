package markdown

import (
	"bufio"
	"fmt"
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
	traversed := traverseContent(line) // TODO: better name please
	return &Paragraph{Children: traversed}
}

const (
	StateText = iota
	StateBold
	StateItalic
	StateLink
)

func traverseContent(line string) []Node {
	var nodes []Node
	state := StateText
	var buffer strings.Builder

	for i := 0; i < len(line); i++ {
		c := line[i]
		if c == '*' && state != StateBold {
			if i+1 < len(line) && line[i+1] == '*' {
				nodes = append(nodes, &Text{Content: buffer.String()})
				buffer.Reset()
				state = StateBold
				i = i + 1
				continue
			} else if state != StateItalic {
				state = StateItalic
			}
		}

		if c == '*' && state == StateBold {
			if i+1 < len(line) && line[i+1] == '*' {
				nodes = append(nodes, &Bold{Content: buffer.String()})
				buffer.Reset()
				state = StateText
				i++
				continue
			}
		}

		// if c == '*' && state == StateItalic {
		//
		// }

		buffer.WriteByte(c)
	}

	if buffer.Len() != 0 {
		switch state {
		case StateText:
			nodes = append(nodes, &Text{Content: buffer.String()})
		case StateBold:
			nodes = append(nodes, &Bold{Content: buffer.String()})
		}
	}

	return nodes
}
