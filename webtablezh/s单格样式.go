package webtablezh

const align居中 = "CENTER"
const align左侧 = "LEFT"
const align右侧 = "RIGHT"

type S单格样式 struct {
	a单格对齐 string
}

func NewS单格样式() *S单格样式 {
	return &S单格样式{
		a单格对齐: align居中, //默认居中显示
	}
}
