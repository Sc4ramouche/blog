package markdown

type Node interface {
	Render() string
}
type Document struct {
	Children []Node
}
type Heading struct {
	Level   int
	Content string
}
type Paragraph struct {
	Children []Node
}

type InlineNode interface {
	Node
}
type Text struct {
	Content string
}
type Bold struct {
	Content string
}
