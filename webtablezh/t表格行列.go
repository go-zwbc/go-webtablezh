package webtablezh

import (
	"strings"

	"github.com/yyle88/tern"
)

type T表格行列 []*C表格单列

func (cs T表格行列) Gen网页表格() string {
	var th表头列表 []string
	for _, c单列 := range cs {
		th表头列表 = append(th表头列表, c单列.new表头字符串TH())
	}
	var th表头内容 = "<tr>" + strings.Join(th表头列表, "") + "</tr>"

	var sz首列高度 = tern.BF(len(cs) > 0, func() int { return len(cs[0].vs单列内容) })
	for _, c单列 := range cs {
		for idx := len(c单列.vs单列内容); idx < sz首列高度; idx++ {
			c单列.V值("") //补个空白就行
		}
	}

	var td各行内容 = make([]string, 0, sz首列高度)
	for idx := 0; idx < sz首列高度; idx++ {
		td单行的值 := make([]string, 0, len(cs))
		for _, c单格 := range cs {
			td单行的值 = append(td单行的值, c单格.new表格字符串TD(idx))
		}
		td单行内容 := "<tr>" + strings.Join(td单行的值, "") + "</tr>"

		td各行内容 = append(td各行内容, td单行内容)
	}

	return `<table border="1">` + th表头内容 + strings.Join(td各行内容, "") + `</table>`
}

func (c *C表格单列) new表头字符串TH() string {
	var align = c.th表头样式.a单格对齐
	var value = c.cn单列列名
	return `<th align="` + align + `">` + value + "</th>"
}

func (c *C表格单列) new表格字符串TD(rowNum int) string {
	var align = c.td单元格式.a单格对齐
	var value = tern.BF(len(c.vs单列内容) > 0, func() string { return c.vs单列内容[rowNum] })
	return `<td align="` + align + `">` + value + "</td>"
}
