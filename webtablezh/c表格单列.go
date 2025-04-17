package webtablezh

import (
	"fmt"

	"github.com/yyle88/tern"
)

type C表格单列 struct {
	cn单列列名 string   //表格列的名称
	vs单列内容 []string //表格列的内容-有多行，因此这里就是表示该列在多行的内容
	th表头样式 *S单格样式   //标题列的样式
	td单元格式 *S单格样式   //内容列的样式
}

func NewC表格单列(name单列列名 string) *C表格单列 {
	return &C表格单列{
		cn单列列名: name单列列名,
		vs单列内容: []string{},
		th表头样式: NewS单格样式(), //这个需要用本体的指针
		td单元格式: NewS单格样式(), //这个需要用本体的指针
	}
}

func (c *C表格单列) V值(v string) *C表格单列 {
	c.vs单列内容 = append(c.vs单列内容, v)
	return c
}

func (c *C表格单列) H样() *E样式和格 {
	return NewE样式和格(c.th表头样式, c)
}

func (c *C表格单列) D样() *E样式和格 {
	return NewE样式和格(c.td单元格式, c)
}

func (c *C表格单列) V整数(v int) *C表格单列 {
	return c.V值(fmt.Sprintf("%d", v))
}

func (c *C表格单列) V两位小数(v float64) *C表格单列 {
	return c.V值(fmt.Sprintf("%.2f", v))
}

func (c *C表格单列) V三位小数(v float64) *C表格单列 {
	return c.V值(fmt.Sprintf("%.3f", v))
}

func (c *C表格单列) V错误(e error) *C表格单列 {
	return c.V值(tern.BF(e != nil, func() string { return e.Error() }))
}
