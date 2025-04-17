package webtablezh_test

import (
	"strings"
	"testing"

	"github.com/go-zwbc/go-webtablezh/webtablezh"
)

func TestT表格行列_Gen网页表格(t *testing.T) {
	tab := webtablezh.T表格行列{
		webtablezh.NewC表格单列("A").V值("1").V值("abc").H样().L向左对齐().D样().R向右对齐(),
		webtablezh.NewC表格单列("B").V值("2").V值("xyz").H样().C居中对齐().D样().C居中对齐(),
		webtablezh.NewC表格单列("C").V值("3").V值("uvw").H样().R向右对齐().D样().L向左对齐(),
	}
	//FIREFOX.ShowPage("表格", tab.GenTable())
	t.Log(convertContent(tab.Gen网页表格()))
}

func TestT表格行列_Gen网页表格1(t *testing.T) {
	tab := webtablezh.T表格行列{
		webtablezh.NewC表格单列("A").V值("1").V值("abc").H样().L向左对齐().D样().R向右对齐(),
		webtablezh.NewC表格单列("B").V值("2").V值("xyz").H样().C居中对齐().D样().C居中对齐(),
		webtablezh.NewC表格单列("C").V值("3").V值("uvw").H样().R向右对齐().D样().L向左对齐(),
	}
	//FIREFOX.ShowPage("表格", tab.GenTable())
	t.Log(convertContent(tab.Gen网页表格()))
}

func TestT表格行列_Gen网页表格2(t *testing.T) {
	//假设你的类型是这样的
	type rowType struct {
		A string
		B string
		C string
	}
	//假设你的数据是这样的
	var rows = []*rowType{
		&rowType{
			A: "a",
			B: "b",
			C: "c",
		},
		&rowType{
			A: "abc",
			B: "xyz",
			C: "uvw",
		},
		&rowType{
			A: "1",
			B: "2",
			C: "3",
		},
		&rowType{
			A: "aaaaaaa",
			B: "bbbbbbb",
			C: "ccccccc",
		},
		&rowType{
			A: "u",
			B: "v",
			C: "w",
		},
	}

	//接下来定义各个列的名称
	var a = webtablezh.NewC表格单列("A").H样().L向左对齐().D样().R向右对齐()
	var b = webtablezh.NewC表格单列("B").H样().C居中对齐().D样().C居中对齐()
	var c = webtablezh.NewC表格单列("C").H样().R向右对齐().D样().L向左对齐()
	for _, row := range rows {
		a.V值(row.A)
		b.V值(row.B)
		c.V值(row.C)
	}
	tab := webtablezh.T表格行列{a, b, c}
	//FIREFOX.ShowPage("表格", tab.GenTable())
	t.Log(convertContent(tab.Gen网页表格()))
}

func convertContent(content string) string {
	content = strings.ReplaceAll(content, "<tr>", "\n\t<tr>\n\t\t")
	content = strings.ReplaceAll(content, "</tr>", "\n\t</tr>\n")
	content = strings.ReplaceAll(content, "\n\n", "\n")
	return "表格的内容:\n" + content
}
