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
	fmt.Fprintf(&result, "<strong>")
	for _, node := range b.Children {
		result.WriteString(node.Render())
	}
	fmt.Fprintf(&result, "</strong>")
	return result.String()
}

func (i *Italic) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<em>")
	for _, node := range i.Children {
		result.WriteString(node.Render())
	}
	fmt.Fprintf(&result, "</em>")
	return result.String()
}

func (l *Link) Render() string {
	var results strings.Builder
	fmt.Fprintf(&results, "<a href=\"%s\">%s</a>", l.Url, l.Text)
	return results.String()
}

func (li *ListItem) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<li>")
	for _, node := range li.Children {
		fmt.Fprintf(&result, node.Render())
	}
	fmt.Fprintf(&result, "</li>")
	return result.String()
}

func (l *List) Render() string {
	var result strings.Builder
	fmt.Fprintf(&result, "<ul>")
	for _, listItem := range l.Children {
		fmt.Fprintf(&result, listItem.Render())
	}
	fmt.Fprintf(&result, "</ul>")
	return result.String()
}
