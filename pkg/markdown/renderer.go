package markdown

import (
	"fmt"
	"strings"
)

func (d *Document) Render() string {
	var result []string
	for _, node := range d.Children {
		result = append(result, node.Render())
	}
	return strings.Join(result, "\n")
}

func (h *Heading) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<h%d>%s</h%d>", h.Level, h.Content, h.Level)
	return result.String()
}

func (p *Paragraph) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<p>%s</p>", p.Content)
	return result.String()
}
