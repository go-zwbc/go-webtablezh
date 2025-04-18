package w3mdebugzh

import (
	"github.com/go-xlan/go-webpage/w3mopenpage"
	"github.com/yyle88/osexec"
)

func Open打开链接(cmd命令配置 *osexec.OsCommand, link链接 string) {
	w3mopenpage.Open(cmd命令配置, link链接)
}

func Show渲染网页(cmd命令配置 *osexec.OsCommand, page网页 string) {
	w3mopenpage.Show(cmd命令配置, page网页)
}
