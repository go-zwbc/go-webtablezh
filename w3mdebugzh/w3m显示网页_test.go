package w3mdebugzh_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/go-zwbc/go-webtablezh/w3mdebugzh"
	"github.com/go-zwbc/go-webtablezh/webtablezh"
	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

type Task struct {
	ID      string    // 任务 ID
	T标题     string    // 任务标题
	D描述     string    // 描述
	S状态     string    // 状态
	A负责人的ID string    // 负责人 ID
	D截止日期   time.Time // 截止日期
	C创建时间   time.Time // 创建时间
}

func newTask() *Task {
	task := &Task{}
	must.Done(gofakeit.Struct(task))
	return task
}

func TestShow渲染网页(t *testing.T) {
	rows := []*Task{
		newTask(),
		newTask(),
		newTask(),
		newTask(),
	}

	//接下来定义各个列的名称
	var I序号 = webtablezh.NewC表格单列("序号").H样().L向左对齐().D样().R向右对齐()
	var T标题 = webtablezh.NewC表格单列("标题").H样().C居中对齐().D样().C居中对齐()
	var D描述 = webtablezh.NewC表格单列("描述").H样().R向右对齐().D样().L向左对齐()
	var S状态 = webtablezh.NewC表格单列("状态").H样().R向右对齐().D样().L向左对齐()
	var A负责人的ID = webtablezh.NewC表格单列("负责人的ID").H样().R向右对齐().D样().L向左对齐()
	var D截止日期 = webtablezh.NewC表格单列("截止日期").H样().R向右对齐().D样().L向左对齐()
	var C创建时间 = webtablezh.NewC表格单列("创建时间").H样().R向右对齐().D样().L向左对齐()
	tab := webtablezh.T表格行列{
		I序号, T标题, D描述, S状态, A负责人的ID, D截止日期, C创建时间,
	}
	for _, row := range rows {
		I序号.V值(row.ID)
		T标题.V值(row.T标题)
		D描述.V值(row.D描述)
		S状态.V值(row.S状态)
		A负责人的ID.V值(row.A负责人的ID)
		D截止日期.V值(row.D截止日期.Format(time.RFC3339))
		C创建时间.V值(row.C创建时间.Format(time.RFC3339))
	}
	w3mdebugzh.Show渲染网页(osexec.NewOsCommand().WithDebug(), tab.Gen网页表格())
}
