package webtablezh

type E样式和格 struct {
	s单格样式 *S单格样式
	c表格单列 *C表格单列
}

func NewE样式和格(s单格样式 *S单格样式, c表格单列 *C表格单列) *E样式和格 {
	return &E样式和格{
		s单格样式: s单格样式,
		c表格单列: c表格单列,
	}
}

func (c *E样式和格) C居中对齐() *C表格单列 {
	c.s单格样式.a单格对齐 = align居中
	return c.c表格单列
}

func (c *E样式和格) L向左对齐() *C表格单列 {
	c.s单格样式.a单格对齐 = align左侧
	return c.c表格单列
}

func (c *E样式和格) R向右对齐() *C表格单列 {
	c.s单格样式.a单格对齐 = align右侧
	return c.c表格单列
}

func (c *E样式和格) C表格单列() *C表格单列 {
	return c.c表格单列
}
