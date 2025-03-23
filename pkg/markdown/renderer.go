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
	fmt.Fprintf(&result, "<p>")
	for _, n := range p.Children {
		result.Write([]byte(n.Render()))
	}
	fmt.Fprintf(&result, "</p>")
	return result.String()
}

func (t *Text) Render() string {
	return t.Content
}

func (b *Bold) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<strong>%s</strong>", b.Content)
	return result.String()
}
