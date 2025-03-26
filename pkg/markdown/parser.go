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
	listStack := []*List{}
	currentIndentationLevel := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			if len(listStack) > 0 {
				headList := listStack[0]
				document.Children = append(document.Children, headList)
				listStack = listStack[:0]
				currentIndentationLevel = 0
			}
			continue
		}

		if isListMarker(line) {
			newLineIndentationLevel := determineIndentation(line)
			if len(listStack) == 0 {
				listStack = append(listStack, newList(0))
			} else if newLineIndentationLevel > currentIndentationLevel {
				topmostList := listStack[len(listStack)-1]
				newListNode := newList(newLineIndentationLevel)
				topmostList.addNestedList(newListNode)
				listStack = append(listStack, newListNode)
				currentIndentationLevel = newLineIndentationLevel
			} else {
				for len(listStack) > 1 && listStack[len(listStack)-1].Level > newLineIndentationLevel {
					listStack = listStack[:len(listStack)-1]
				}
				currentIndentationLevel = newLineIndentationLevel
			}

			currentList := listStack[len(listStack)-1]
			parseListItem(currentList, line)
		} else {
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
	}

	if len(listStack) > 0 {
		document.Children = append(document.Children, listStack[0])
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
	inlineNodes, err := parseInlineContent(line)
	if err != nil {
		fmt.Println("Inline parse error:", err)
	}
	return &Paragraph{Children: inlineNodes}
}

func parseListItem(list *List, line string) {
	contentIndex := strings.Index(line, "- ")
	content := line[contentIndex+2:]
	inlineNodes, err := parseInlineContent(content)
	if err != nil {
		fmt.Println("Inline parse error:", err)
	}
	listItem := newListItem(inlineNodes)
	list.Children = append(list.Children, *listItem)
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
					nodes = append(nodes, newTextNode(buffer.String()))
					nodesStack = append(nodesStack, newBoldNode())
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
					nodesStack = append(nodesStack, newBoldNode())
				}
				i++
			} else {
				switch currentNode.(type) {
				case nil:
					nodes = append(nodes, newTextNode(buffer.String()))
					nodesStack = append(nodesStack, newItalicNode())
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
					nodesStack = append(nodesStack, newItalicNode())
				}
			}
			buffer.Reset()
			continue
		}

		if c == '[' {
			switch currentNode.(type) {
			case nil:
				nodes = append(nodes, newTextNode(buffer.String()))
				nodesStack = append(nodesStack, newLinkNode())
			default:
				appendContent(currentNode, buffer.String())
				nodesStack = append(nodesStack, newLinkNode())
			}
			buffer.Reset()
			continue
		}

		if c == ']' {
			switch currentNode.(type) {
			case *Link:
				linkNode := currentNode.(*Link)
				linkNode.Text = buffer.String()
			}
			buffer.Reset()
			continue
		}

		if c == '(' && currentNode.Type() == LinkNode {
			buffer.Reset()
			continue
		}
		if c == ')' && currentNode.Type() == LinkNode {
			linkNode := currentNode.(*Link)
			linkNode.Url = buffer.String()
			buffer.Reset()

			if len(nodesStack) == 1 {
				nodes = append(nodes, linkNode)
				nodesStack = nodesStack[:len(nodesStack)-1]
			} else {
				nodesStack = nodesStack[:len(nodesStack)-1]
				parentNode := nodesStack[len(nodesStack)-1]
				appendChildNode(parentNode, linkNode)
			}
			continue
		}

		buffer.WriteByte(c)
	}

	if buffer.Len() != 0 {
		nodes = append(nodes, newTextNode(buffer.String()))
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

func isListMarker(line string) bool {
	trimmed := strings.TrimSpace(line)
	return strings.HasPrefix(trimmed, "- ")
}

func determineIndentation(line string) int {
	index := strings.Index(line, "- ")
	return index / 2
}
