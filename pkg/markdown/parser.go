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
	traversed, err := parseInlineContent(line)
	if err != nil {
		fmt.Println("Inline parse error:", err)
	}
	return &Paragraph{Children: traversed}
}

// TODO: Debug view for AST could be handy
func parseInlineContent(line string) ([]InlineNode, error) {
	var nodes []InlineNode
	nodesStack := []InlineNode{}
	var buffer strings.Builder

	for i := 0; i < len(line); i++ {
		c := line[i]
		var currentNode InlineNode
		if len(nodesStack) != 0 {
			currentNode = nodesStack[len(nodesStack)-1]
		} else {
			currentNode = nil
		}

		if c == '*' {
			if i+1 < len(line) && line[i+1] == '*' {
				switch currentNode.(type) {
				case nil:
					nodes = append(nodes, &Text{buffer.String()})
					nodesStack = append(nodesStack, &Bold{Children: []Node{}})
				case *Bold:
					if len(nodesStack) == 1 {
						appendContent(currentNode, buffer.String())
						nodes = append(nodes, currentNode)
						nodesStack = nodesStack[:len(nodesStack)-1]
					} else {
						appendContent(currentNode, buffer.String())
						boldNode := nodesStack[len(nodesStack)-1]
						nodesStack = nodesStack[:len(nodesStack)-1]
						parentNode := nodesStack[len(nodesStack)-1]
						appendChildNode(parentNode, boldNode)
					}
				default:
					appendContent(currentNode, buffer.String())
					nodesStack = append(nodesStack, &Bold{Children: []Node{}})
				}
				i++
			} else {
				switch currentNode.(type) {
				case nil:
					nodes = append(nodes, &Text{buffer.String()})
					nodesStack = append(nodesStack, &Italic{Children: []Node{}})
				case *Italic:
					if len(nodesStack) == 1 {
						appendContent(currentNode, buffer.String())
						nodes = append(nodes, currentNode)
						nodesStack = nodesStack[:len(nodesStack)-1]
					} else {
						appendContent(currentNode, buffer.String())
						italicNode := nodesStack[len(nodesStack)-1]
						nodesStack = nodesStack[:len(nodesStack)-1]
						parentNode := nodesStack[len(nodesStack)-1]
						appendChildNode(parentNode, italicNode)
					}
				default:
					appendContent(currentNode, buffer.String())
					nodesStack = append(nodesStack, &Italic{Children: []Node{}})
				}
			}
			buffer.Reset()
			continue
		}

		buffer.WriteByte(c)
	}

	if buffer.Len() != 0 {
		nodes = append(nodes, &Text{Content: buffer.String()})
	}

	if len(nodesStack) != 0 {
		nodeType := nodesStack[len(nodesStack)-1].Type()
		var tagName string

		switch nodeType {
		case BoldNode:
			tagName = "bold (**)"
		case ItalicNode:
			tagName = "italic (*)"
		default:
			tagName = "unknown"
		}

		return nodes, fmt.Errorf("unclosed %s tag", tagName)
	}

	return nodes, nil
}
