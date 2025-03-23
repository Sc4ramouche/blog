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
	Content string
    Children []Node
}

// Inline Nodes

type Text struct {
    Content string
}

type Bold struct {
    Content string
}
