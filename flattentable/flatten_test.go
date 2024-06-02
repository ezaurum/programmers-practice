package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	const input = `<body>
<table><tr><td>A</td><td><table><tr><td>B</td><td>C</td></tr></table></td></tr></table>
<table><tr><td>The</td><td>rain</td><td><table><tr><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr></table></td></tr><tr><td><table><tr><td>in</td></tr><tr><td>in</td></tr></table></td><td>Spain</td><td>stays</td></tr></table>
</body>`
	const output = `<body>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
<table><tr><td rowspan="2">The</td><td rowspan="2">rain</td><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr><tr><td>in</td><td rowspan="2">Spain</td><td rowspan="2" colspan="2">stays</td></tr><tr><td>in</td></tr></table>
</body>`
	r := Solution(input)
	if r != output {
		t.Error("Expected ", output, ", got ", r)
	}
}

func TestSolution2(t *testing.T) {
	const input = `<body>
<table><tr><td>A</td><td><table><tr><td>B</td><td><table><tr><td>C</td></tr></table></td></tr></table></td></tr></table>
</body>`
	const output = `<body>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
</body>`
	r := Solution(input)
	if r != output {
		t.Error("Expected ", output, ", got ", r)
	}
}

func TestSolution3(t *testing.T) {
	const input = `<body>
<table><tr><td>A</td><td><table><tr><td>B</td><td>C</td></tr></table></td></tr></table>
<table><tr><td>The</td><td>rain</td><td><table><tr><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr></table></td></tr><tr><td><table><tr><td>in</td></tr><tr><td>in</td></tr></table></td><td>Spain</td><td>stays</td></tr></table>
<table><tr><td>A</td><td><table><tr><td>B</td><td>C</td></tr></table></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td><table></table></td></tr></table>
<table><tr><td><table><tr><td><table><tr><td><table></table></td></tr></table></td></tr></table></td></tr></table>
<table><tr><td><table><tr><td><table><tr><td><table><tr><td>deep</td></tr></table></td></tr></table></td></tr></table></td></tr></table>
<table><tr><td><table><tr><td>foo</td><td>bar</td></tr></table></td></tr><tr><td>spanned</td></tr></table>
<table><tr><td>A</td><td><table><tr><td>B</td></tr><tr><td>C</td></tr></table></td></tr></table>
<table><tr><td>The</td><td>rain</td><td><table><tr><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr></table></td></tr><tr><td><table><tr><td>in</td></tr><tr><td>in</td></tr></table></td><td>Spain</td><td>stays</td></tr></table>
<table><tr><td><table><tr><td>2</td><td><table><tr><td>3</td></tr><tr><td><table><tr><td>5</td><td>4</td></tr></table></td></tr></table></td></tr></table></td></tr><tr><td>1</td></tr></table>
<table><tr><td><table><tr><td>1</td><td><table><tr><td><table><tr><td>2</td><td><table><tr><td><table><tr><td>3</td><td><table><tr><td><table><tr><td>4</td><td><table><tr><td><table><tr><td>5</td><td>GOAL</td></tr></table></td></tr><tr><td>a</td></tr></table></td></tr></table></td></tr><tr><td>b</td></tr></table></td></tr></table></td></tr><tr><td>c</td></tr></table></td></tr></table></td></tr><tr><td>d</td></tr></table></td></tr></table></td></tr><tr><td>e</td></tr></table>
<table><tr><td><table><tr><td>AA</td><td>BB</td><td>CC</td></tr><tr><td>DD</td><td>EE</td><td>FF</td></tr><tr><td>GG</td><td>HH</td><td>II</td></tr></table></td><td>b</td><td>c</td></tr><tr><td>d</td><td><table><tr><td>A1</td><td>B1</td><td>C1</td></tr><tr><td><table><tr><td>A2</td><td>B2</td><td>C3</td></tr><tr><td>D1</td><td>E1</td><td>F2</td></tr><tr><td>G1</td><td><table><tr><td>A3</td><td>B3</td><td>C5</td></tr><tr><td>D2</td><td>E2</td><td><table><tr><td><table><tr><td>A4</td><td>B4</td><td>C7</td></tr><tr><td>D3</td><td><table><tr><td>A5</td><td>B5</td><td>C9</td></tr><tr><td></td><td>E3ababababababababababababababababababababababababababababababababababababababababababababababababab</td><td>F4</td></tr><tr><td>G2</td><td>H1</td><td>I1</td></tr></table></td><td>F6</td></tr><tr><td>G3</td><td>H2</td><td>I2</td></tr></table></td><td>B6</td><td>C11</td></tr><tr><td>D5</td><td>E4</td><td>F8</td></tr><tr><td>G4</td><td>H3</td><td>I3</td></tr></table></td></tr><tr><td>G5</td><td>H4</td><td>I4</td></tr></table></td><td>I5</td></tr></table></td><td>E5</td><td>F10</td></tr><tr><td>G6</td><td>H5</td><td>I6</td></tr></table></td><td>f</td></tr><tr><td>g</td><td>h</td><td>i</td></tr></table>
</body>
`
	const output = `<body>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
<table><tr><td rowspan="2">The</td><td rowspan="2">rain</td><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr><tr><td>in</td><td rowspan="2">Spain</td><td rowspan="2" colspan="2">stays</td></tr><tr><td>in</td></tr></table>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td>deep</td></tr></table>
<table><tr><td>foo</td><td>bar</td></tr><tr><td colspan="2">spanned</td></tr></table>
<table><tr><td rowspan="2">A</td><td>B</td></tr><tr><td>C</td></tr></table>
<table><tr><td rowspan="2">The</td><td rowspan="2">rain</td><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr><tr><td>in</td><td rowspan="2">Spain</td><td rowspan="2" colspan="2">stays</td></tr><tr><td>in</td></tr></table>
<table><tr><td rowspan="2">2</td><td colspan="2">3</td></tr><tr><td>5</td><td>4</td></tr><tr><td colspan="3">1</td></tr></table>
<table><tr><td rowspan="5">1</td><td rowspan="4">2</td><td rowspan="3">3</td><td rowspan="2">4</td><td>5</td><td>GOAL</td></tr><tr><td colspan="2">a</td></tr><tr><td colspan="3">b</td></tr><tr><td colspan="4">c</td></tr><tr><td colspan="5">d</td></tr><tr><td colspan="6">e</td></tr></table>
<table><tr><td>AA</td><td>BB</td><td>CC</td><td rowspan="3" colspan="13">b</td><td rowspan="3">c</td></tr><tr><td>DD</td><td>EE</td><td>FF</td></tr><tr><td>GG</td><td>HH</td><td>II</td></tr><tr><td rowspan="13" colspan="3">d</td><td colspan="11">A1</td><td>B1</td><td>C1</td><td rowspan="13">f</td></tr><tr><td>A2</td><td colspan="9">B2</td><td>C3</td><td rowspan="11">E5</td><td rowspan="11">F10</td></tr><tr><td>D1</td><td colspan="9">E1</td><td>F2</td></tr><tr><td rowspan="9">G1</td><td>A3</td><td>B3</td><td colspan="7">C5</td><td rowspan="9">I5</td></tr><tr><td rowspan="7">D2</td><td rowspan="7">E2</td><td>A4</td><td colspan="3">B4</td><td>C7</td><td rowspan="5">B6</td><td rowspan="5">C11</td></tr><tr><td rowspan="3">D3</td><td>A5</td><td>B5</td><td>C9</td><td rowspan="3">F6</td></tr><tr><td></td><td>E3ababababababababababababababababababababababababababababababababababababababababababababababababab</td><td>F4</td></tr><tr><td>G2</td><td>H1</td><td>I1</td></tr><tr><td>G3</td><td colspan="3">H2</td><td>I2</td></tr><tr><td colspan="5">D5</td><td>E4</td><td>F8</td></tr><tr><td colspan="5">G4</td><td>H3</td><td>I3</td></tr><tr><td>G5</td><td>H4</td><td colspan="7">I4</td></tr><tr><td colspan="11">G6</td><td>H5</td><td>I6</td></tr><tr><td colspan="3">g</td><td colspan="13">h</td><td>i</td></tr></table>
</body>`
	r := Solution(input)
	if r != output {
		t.Error("Expected ", output, ", got ", r)
	}
}

func TestSolutionGPT(t *testing.T) {
	const input = `<body>
<table><tr><td>A</td><td><table><tr><td>B</td><td><table><tr><td>C1</td></tr><tr><td>C2</td></tr></table></td></tr><tr><td>A2</td><td>B2</td><td>C2</td></tr></table></td></tr></table>
</body>`
	const output = `<body>
<table><tr><td rowspan="2">A</td><td rowspan="2">B</td><td>C1</td></tr><tr><td>C2</td></tr><tr><td>A2</td><td>B2</td><td>C2</td></tr>
</table>
</body>`
	r := Solution(input)
	if r != output {
		t.Error("Expected ", output, ", got ", r)
	}
}
func TestSolution4(t *testing.T) {
	const input = `<body>
<table><tr><td>A</td><td><table><tr><td>B</td><td>C</td></tr></table></td></tr></table>
<table><tr><td>The</td><td>rain</td><td><table><tr><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr></table></td></tr><tr><td><table><tr><td>in</td></tr><tr><td>in</td></tr></table></td><td>Spain</td><td>stays</td></tr></table>
<table><tr><td>A</td><td><table><tr><td>B</td><td>C</td></tr></table></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td><table></table></td></tr></table>
<table><tr><td><table><tr><td><table><tr><td><table></table></td></tr></table></td></tr></table></td></tr></table>
<table><tr><td><table><tr><td><table><tr><td><table><tr><td>deep</td></tr></table></td></tr></table></td></tr></table></td></tr></table>
<table><tr><td><table><tr><td>foo</td><td>bar</td></tr></table></td></tr><tr><td>spanned</td></tr></table>
<table><tr><td>A</td><td><table><tr><td>B</td></tr><tr><td>C</td></tr></table></td></tr></table>
<table><tr><td>The</td><td>rain</td><td><table><tr><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr></table></td></tr><tr><td><table><tr><td>in</td></tr><tr><td>in</td></tr></table></td><td>Spain</td><td>stays</td></tr></table>
<table><tr><td><table><tr><td>2</td><td><table><tr><td>3</td></tr><tr><td><table><tr><td>5</td><td>4</td></tr></table></td></tr></table></td></tr></table></td></tr><tr><td>1</td></tr></table>
<table><tr><td><table><tr><td>1</td><td><table><tr><td><table><tr><td>2</td><td><table><tr><td><table><tr><td>3</td><td><table><tr><td><table><tr><td>4</td><td><table><tr><td><table><tr><td>5</td><td>GOAL</td></tr></table></td></tr><tr><td>a</td></tr></table></td></tr></table></td></tr><tr><td>b</td></tr></table></td></tr></table></td></tr><tr><td>c</td></tr></table></td></tr></table></td></tr><tr><td>d</td></tr></table></td></tr></table></td></tr><tr><td>e</td></tr></table>
<table><tr><td><table><tr><td>AA</td><td>BB</td><td>CC</td></tr><tr><td>DD</td><td>EE</td><td>FF</td></tr><tr><td>GG</td><td>HH</td><td>II</td></tr></table></td><td>b</td><td>c</td></tr><tr><td>d</td><td><table><tr><td>A1</td><td>B1</td><td>C1</td></tr><tr><td><table><tr><td>A2</td><td>B2</td><td>C3</td></tr><tr><td>D1</td><td>E1</td><td>F2</td></tr><tr><td>G1</td><td><table><tr><td>A3</td><td>B3</td><td>C5</td></tr><tr><td>D2</td><td>E2</td><td><table><tr><td><table><tr><td>A4</td><td>B4</td><td>C7</td></tr><tr><td>D3</td><td><table><tr><td>A5</td><td>B5</td><td>C9</td></tr><tr><td></td><td>E3ababababababababababababababababababababababababababababababababababababababababababababababababab</td><td>F4</td></tr><tr><td>G2</td><td>H1</td><td>I1</td></tr></table></td><td>F6</td></tr><tr><td>G3</td><td>H2</td><td>I2</td></tr></table></td><td>B6</td><td>C11</td></tr><tr><td>D5</td><td>E4</td><td>F8</td></tr><tr><td>G4</td><td>H3</td><td>I3</td></tr></table></td></tr><tr><td>G5</td><td>H4</td><td>I4</td></tr></table></td><td>I5</td></tr></table></td><td>E5</td><td>F10</td></tr><tr><td>G6</td><td>H5</td><td>I6</td></tr></table></td><td>f</td></tr><tr><td>g</td><td>h</td><td>i</td></tr></table>
</body>
`
	const output = `<body>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
<table><tr><td rowspan="2">The</td><td rowspan="2">rain</td><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr><tr><td>in</td><td rowspan="2">Spain</td><td rowspan="2" colspan="2">stays</td></tr><tr><td>in</td></tr></table>
<table><tr><td>A</td><td>B</td><td>C</td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td></td></tr></table>
<table><tr><td>deep</td></tr></table>
<table><tr><td>foo</td><td>bar</td></tr><tr><td colspan="2">spanned</td></tr></table>
<table><tr><td rowspan="2">A</td><td>B</td></tr><tr><td>C</td></tr></table>
<table><tr><td rowspan="2">The</td><td rowspan="2">rain</td><td>1</td><td>2</td></tr><tr><td>3</td><td>4</td></tr><tr><td>in</td><td rowspan="2">Spain</td><td rowspan="2" colspan="2">stays</td></tr><tr><td>in</td></tr></table>
<table><tr><td rowspan="2">2</td><td colspan="2">3</td></tr><tr><td>5</td><td>4</td></tr><tr><td colspan="3">1</td></tr></table>
<table><tr><td rowspan="5">1</td><td rowspan="4">2</td><td rowspan="3">3</td><td rowspan="2">4</td><td>5</td><td>GOAL</td></tr><tr><td colspan="2">a</td></tr><tr><td colspan="3">b</td></tr><tr><td colspan="4">c</td></tr><tr><td colspan="5">d</td></tr><tr><td colspan="6">e</td></tr></table>
<table><tr><td>AA</td><td>BB</td><td>CC</td><td rowspan="3" colspan="13">b</td><td rowspan="3">c</td></tr><tr><td>DD</td><td>EE</td><td>FF</td></tr><tr><td>GG</td><td>HH</td><td>II</td></tr><tr><td rowspan="13" colspan="3">d</td><td colspan="11">A1</td><td>B1</td><td>C1</td><td rowspan="13">f</td></tr><tr><td>A2</td><td colspan="9">B2</td><td>C3</td><td rowspan="11">E5</td><td rowspan="11">F10</td></tr><tr><td>D1</td><td colspan="9">E1</td><td>F2</td></tr><tr><td rowspan="9">G1</td><td>A3</td><td>B3</td><td colspan="7">C5</td><td rowspan="9">I5</td></tr><tr><td rowspan="7">D2</td><td rowspan="7">E2</td><td>A4</td><td colspan="3">B4</td><td>C7</td><td rowspan="5">B6</td><td rowspan="5">C11</td></tr><tr><td rowspan="3">D3</td><td>A5</td><td>B5</td><td>C9</td><td rowspan="3">F6</td></tr><tr><td></td><td>E3ababababababababababababababababababababababababababababababababababababababababababababababababab</td><td>F4</td></tr><tr><td>G2</td><td>H1</td><td>I1</td></tr><tr><td>G3</td><td colspan="3">H2</td><td>I2</td></tr><tr><td colspan="5">D5</td><td>E4</td><td>F8</td></tr><tr><td colspan="5">G4</td><td>H3</td><td>I3</td></tr><tr><td>G5</td><td>H4</td><td colspan="7">I4</td></tr><tr><td colspan="11">G6</td><td>H5</td><td>I6</td></tr><tr><td colspan="3">g</td><td colspan="13">h</td><td>i</td></tr></table>
</body>`
	r := Solution(input)
	if r != output {
		//t.Error("Expected ", output, ", got ", r)
		split0 := strings.Split(input, "\n")
		split := strings.Split(output, "\n")
		split1 := strings.Split(r, "\n")
		for i, expected := range split {
			result := split1[i]
			if result != expected {
				inputString := split0[i]
				fmt.Println(i, inputString)
				fmt.Println(i, expected)
				fmt.Println(i, result)
			}
		}
	}
}
