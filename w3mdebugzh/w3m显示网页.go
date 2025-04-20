package w3mdebugzh

import (
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/osexec"
	"github.com/yyle88/zaplog"
)

func Open打开链接(link链接 string) {
	w3mopenpage.Open(osexec.NewOsCommand().WithDebug(), link链接)
}

func Show渲染网页(page网页 string) {
	w3mopenpage.Show(osexec.NewOsCommand().WithDebug(), page网页)
}

func Show标题内容(s标题 string, s内容 string) {
	zaplog.SUG.Debug("----\n", eroticgo.PINK.Sprint(s标题), "\n----")
	page网页 := `<!DOCTYPE html><html lang="zh"><head><meta charset="UTF-8"><title>` + s标题 + `</title></head><body>` + s内容 + `</body></html>`
	w3mopenpage.Show(osexec.NewOsCommand().WithDebug(), page网页)
}
