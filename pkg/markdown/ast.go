package markdown

import "time"

type Node interface {
	Render() string
}
type Document struct {
	Children []Node
	Title    string
	Date     time.Time
	Draft    string
	Path     string
}
type Heading struct {
	Level   int
	Content string
}
type Paragraph struct {
	Children []InlineNode
}
type ListItem struct {
	Children   []InlineNode
	NestedList *List
	Level      int
}

func newListItem(content []InlineNode) *ListItem {
	return &ListItem{Children: content}
}

type List struct {
	Children []ListItem
	Level    int
}

func (l *List) addNestedList(list *List) {
	if len(l.Children) > 0 {
		l.Children[len(l.Children)-1].NestedList = list
	}
}

func newList(level int) *List {
	return &List{Children: []ListItem{}, Level: level}
}

type NodeType int

const (
	TextNode NodeType = iota
	ItalicNode
	BoldNode
	InlineCodeNode
	LinkNode
)

type InlineNode interface {
	Node
	Type() NodeType
}

type Text struct {
	Content  string
	nodeType NodeType
}

func newTextNode(content string) *Text {
	return &Text{Content: content, nodeType: TextNode}
}
func (t *Text) Type() NodeType {
	return t.nodeType
}

type Bold struct {
	Children []Node
	nodeType NodeType
}

func newBoldNode() *Bold {
	return &Bold{Children: []Node{}, nodeType: BoldNode}
}

func (b *Bold) Type() NodeType {
	return b.nodeType
}

type Italic struct {
	Children []Node
	nodeType NodeType
}

func newItalicNode() *Italic {
	return &Italic{Children: []Node{}, nodeType: ItalicNode}
}

func (i *Italic) Type() NodeType {
	return i.nodeType
}

type Link struct {
	Url      string
	Text     string
	nodeType NodeType
}

func newLinkNode() *Link {
	return &Link{Url: "", Text: "", nodeType: LinkNode}
}

func (l *Link) Type() NodeType {
	return l.nodeType
}

type InlineCode struct {
	Children []Node
	nodeType NodeType
}

func newInlineCodeNode() *InlineCode {
	return &InlineCode{Children: []Node{}, nodeType: InlineCodeNode}
}

func (c *InlineCode) Type() NodeType {
	return c.nodeType
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
	case *InlineCode:
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
