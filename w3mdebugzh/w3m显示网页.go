package w3mdebugzh

import (
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/osexec"
)

func Open打开链接(link链接 string) {
	w3mopenpage.Open(osexec.NewOsCommand().WithDebug(), link链接)
}

func Show渲染网页(page网页 string) {
	w3mopenpage.Show(osexec.NewOsCommand().WithDebug(), page网页)
}
