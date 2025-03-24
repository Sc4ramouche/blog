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
	Children []InlineNode
}

type NodeType int

const (
	TextNode NodeType = iota
	ItalicNode
	BoldNode
)

type InlineNode interface {
	Node
	Type() NodeType
}

type Text struct {
	Content string
    nodeType NodeType
}

func (t *Text) Type() NodeType {
	return t.nodeType
}

type Bold struct {
	Children []Node
}

func (b *Bold) Type() NodeType {
	return BoldNode
}

type Italic struct {
	Children []Node
}

func (i *Italic) Type() NodeType {
	return ItalicNode
}

func appendContent(node InlineNode, content string) {
	if content == "" {
		return
	}
	
	switch n := node.(type) {
	case *Text:
		n.Content += content
	case *Bold:
		n.Children = append(n.Children, newTextNode(content))
	case *Italic:
		n.Children = append(n.Children, newTextNode(content))
	}
}

func appendChildNode(parent InlineNode, child InlineNode) {
	switch p := parent.(type) {
	case *Bold:
		p.Children = append(p.Children, child)
	case *Italic:
		p.Children = append(p.Children, child)
	}
}

func newTextNode(content string) *Text {
	return &Text{Content: content, nodeType: TextNode}
}
