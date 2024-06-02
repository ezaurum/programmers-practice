package main

import (
	"fmt"
	"strings"
)

type Tag struct {
	Content  string
	RowSpan  int
	ColSpan  int
	Children []*Tag
	Index    int
	Parent   *Tag
	Name     string
}

func (t *Tag) AddChild(child *Tag) {
	child.Index = len(t.Children)
	child.Parent = t
	t.Children = append(t.Children, child)
}

func (t *Tag) AddContent(content string) {
	t.Content = content
}

func (t *Tag) String() string {
	var builder strings.Builder
	if t.Name == "table" {
		builder.WriteString("\n")
	}
	builder.WriteString("<")
	builder.WriteString(t.Name)
	if t.RowSpan > 1 {
		builder.WriteString(" rowspan=\"")
		builder.WriteString(fmt.Sprint(t.RowSpan))
		builder.WriteString("\"")
	}
	if t.ColSpan > 1 {
		builder.WriteString(" colspan=\"")
		builder.WriteString(fmt.Sprint(t.ColSpan))
		builder.WriteString("\"")
	}
	builder.WriteString(">")
	if t.Name == "td" {
		builder.WriteString(t.Content)
	}
	for _, child := range t.Children {
		builder.WriteString(child.String())
	}
	if t.Name == "body" {
		builder.WriteString("\n")
	}
	builder.WriteString("</")
	builder.WriteString(t.Name)
	builder.WriteString(">")
	return builder.String()
}

func Parse(html string) *Tag {
	var root *Tag
	var currentTag *Tag
	tokens := strings.FieldsFunc(html, func(r rune) bool {
		return r == '<' || r == '>'
	})

	for _, token := range tokens {
		switch {

		case strings.HasPrefix(token, "td"):
			arr := strings.Split(token, " ")
			t := &Tag{Name: arr[0],
				RowSpan: 1,
				ColSpan: 1}
			if nil != currentTag {
				currentTag.AddChild(t)
			}
			currentTag = t
			for _, s := range arr[1:] {
				if strings.HasPrefix(s, "rowspan") {
					_, _ = fmt.Sscanf(s, "rowspan=\"%d\"", &t.RowSpan)
				} else if strings.HasPrefix(s, "colspan") {
					_, _ = fmt.Sscanf(s, "colspan=\"%d\"", &t.ColSpan)
				}
			}
		case token == "body":
			fallthrough
		case token == "tr":
			fallthrough
		case token == "table":
			t := &Tag{Name: token}
			if nil != currentTag {
				currentTag.AddChild(t)
			}
			currentTag = t
			if nil == root {
				root = t
			}
		case token == "/body":
			return currentTag
		case token == "/td":
			fallthrough
		case token == "/tr":
			fallthrough
		case token == "/table":
			currentTag = currentTag.Parent
		default:
			if currentTag != nil {
				currentTag.AddContent(token)
			}
		}
	}

	return root
}

func (t *Tag) UpdateRowCol(row, col int) {
	t.RowSpan += row
	t.ColSpan += col
}

func (t *Tag) Flatten() {
	if len(t.Children) == 0 {
		return
	}

	if t.Name == "td" {
		if len(t.Children) == 1 && t.Children[0].Name == "table" {
			targetTable := t.Children[0]
			targetTable.Flatten()
			trs := targetTable.Children
			if len(trs) == 0 {
				t.RemoveChild(targetTable)
				return
			}
			// 일단 자기 자신은 없애고 tr올린다
			parentTr := t.Parent
			parentTr.RemoveChild(t)

			addingRows := trs[1:]

			// 같은 가로축 rowspan을 추가한다.
			rowSpan := len(trs)
			for _, child := range parentTr.Children {
				if child.RowSpan < rowSpan {
					child.RowSpan = rowSpan
				}
			}

			parentTable := parentTr.Parent
			// tr을 parentTable에 올린다
			firstRow := trs[0]
			countTD := len(firstRow.Children)

			currentTDIndex := t.Index
			parentTr.AddChildrenAfter(firstRow.Children, currentTDIndex-1)
			parentTable.AddChildrenAfter(addingRows, parentTr.Index)

			// 같은 세로축 colspan을 추가한다.
			for _, parentTableTr := range parentTable.Children {
				if parentTableTr != parentTr {
					for _, child := range parentTableTr.Children {
						if child.Index == currentTDIndex && child.ColSpan < countTD {
							child.ColSpan = countTD
						}
					}
				}
			}
			parentTable.Flatten()
			return
		}
	}

	for _, child := range t.Children {
		child.Flatten()
	}
}

func (t *Tag) RemoveChild(t2 *Tag) {
	tt := -1
	for i, child := range t.Children {
		if child == t2 {
			tt = i
			break
		}
	}
	if tt > -1 {
		t.Children = append(t.Children[:tt], t.Children[tt+1:]...)
	}
	return
}

func (t *Tag) AddChildrenAfter(rows []*Tag, index int) {
	if index >= len(t.Children) {
		index = len(t.Children) - 1
	}
	tags := t.Children[:index+1]
	i := append(rows, t.Children[index+1:]...)
	t.Children = append(tags, i...)
	for i, child := range t.Children {
		child.Index = i
	}
}

func Solution(input string) string {
	inputTag := Parse(input)
	inputTag.Flatten()
	return inputTag.String()
}
