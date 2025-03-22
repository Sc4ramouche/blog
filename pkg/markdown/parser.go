package markdown

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node interface {
	Render() string
}

type Document struct {
	Children []Node
}

func (d *Document) Render() string {
	var result []string
	for _, node := range d.Children {
		result = append(result, node.Render())
	}
	return strings.Join(result, "\n")
}

type Heading struct {
	Level   int
	Content string
}

func (h *Heading) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<h%d>%s</h%d", h.Level, h.Content, h.Level)
	return result.String()
}

type Paragraph struct {
	Content string
}

func (p *Paragraph) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<p>%s</p>", p.Content)
	return result.String()
}

func ParseMarkdown(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Could not open file %s: %w", path, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	var document Document
	for scanner.Scan() {
		line := scanner.Text()
		count++

		if len(line) == 0 {
			continue
		}

		firstChar := line[0]
		switch firstChar {
		case '#':
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
			line := strings.TrimSpace(line[contentIndex:])
			document.Children = append(document.Children, &Heading{Level: level, Content: line})
		case '*':
		case '-':
			fmt.Println("LIST", count, line)
		default:
			document.Children = append(document.Children, &Paragraph{Content: line})
		}
	}

	fmt.Println(document.Render())
	return nil
}
