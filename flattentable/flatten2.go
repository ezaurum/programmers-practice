package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Cell struct {
	row, col   int
	rowspan    int
	colspan    int
	content    string
	isOccupied bool
}

type Table struct {
	rows [][]*Cell
}

func parseTable(html string) *Table {
	table := &Table{}
	stack := []*Table{}
	var row []*Cell
	var cell *Cell
	inCell := false
	content := ""

	tokens := strings.Split(html, "<")
	for _, token := range tokens {
		token = strings.TrimSpace(token)
		if token == "" {
			continue
		}
		if strings.HasPrefix(token, "table") {
			if table != nil {
				stack = append(stack, table)
			}
			table = &Table{}
		} else if strings.HasPrefix(token, "tr") {
			row = []*Cell{}
		} else if strings.HasPrefix(token, "/tr") {
			table.rows = append(table.rows, row)
		} else if strings.HasPrefix(token, "td") {
			cell = &Cell{rowspan: 1, colspan: 1}
			if strings.Contains(token, `rowspan="`) {
				fmt.Sscanf(token, `td rowspan="%d"`, &cell.rowspan)
			}
			if strings.Contains(token, `colspan="`) {
				fmt.Sscanf(token, `td colspan="%d"`, &cell.colspan)
			}
			inCell = true
			content = ""
		} else if strings.HasPrefix(token, "/td") {
			cell.content = content
			row = append(row, cell)
			inCell = false
		} else if strings.HasPrefix(token, "/table") {
			if len(stack) > 0 {
				parent := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				for _, r := range table.rows {
					for _, c := range r {
						if c != nil {
							parent.rows = append(parent.rows, []*Cell{c})
						}
					}
				}
				table = parent
			}
		} else {
			if inCell {
				content += "<" + token
			}
		}
	}

	return table
}

func flattenTable(table *Table) *Table {
	maxRows := 100
	maxCols := 100
	flatTable := &Table{rows: make([][]*Cell, maxRows)}

	for i := range flatTable.rows {
		flatTable.rows[i] = make([]*Cell, maxCols)
		for j := range flatTable.rows[i] {
			flatTable.rows[i][j] = &Cell{}
		}
	}

	rowOffset := 0
	colOffset := 0

	for _, row := range table.rows {
		for _, cell := range row {
			for flatTable.rows[rowOffset][colOffset].isOccupied {
				colOffset++
			}
			for i := 0; i < cell.rowspan; i++ {
				for j := 0; j < cell.colspan; j++ {
					flatTable.rows[rowOffset+i][colOffset+j] = &Cell{
						row:        rowOffset,
						col:        colOffset,
						rowspan:    cell.rowspan,
						colspan:    cell.colspan,
						content:    cell.content,
						isOccupied: true,
					}
				}
			}
			colOffset++
		}
		rowOffset++
		colOffset = 0
	}

	return flatTable
}

func renderTable(table *Table) string {
	var buf bytes.Buffer

	buf.WriteString("<table>")

	for _, row := range table.rows {
		if row[0] == nil || row[0].content == "" {
			continue
		}
		buf.WriteString("<tr>")
		for _, cell := range row {
			if cell == nil || cell.content == "" {
				continue
			}
			if cell.isOccupied {
				buf.WriteString("<td")
				if cell.rowspan > 1 {
					buf.WriteString(fmt.Sprintf(` rowspan="%d"`, cell.rowspan))
				}
				if cell.colspan > 1 {
					buf.WriteString(fmt.Sprintf(` colspan="%d"`, cell.colspan))
				}
				buf.WriteString(">")
				buf.WriteString(cell.content)
				buf.WriteString("</td>")
			}
		}
		buf.WriteString("</tr>")
	}
	buf.WriteString("</table>")

	return buf.String()
}

func SolutionByChatGPT(input string) string {

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	var result strings.Builder
	result.WriteString("<body>")
	for _, line := range lines {
		if line == "<body>" || line == "</body>" {
			continue
		}
		table := parseTable(line)
		flatTable := flattenTable(table)
		result.WriteString(renderTable(flatTable))
	}
	result.WriteString("</body>")
	return result.String()
}
