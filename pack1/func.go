package pack1

import (
	"runtime"
	"strconv"
	"strings"
)

const nullRid int64 = -1

// 获取线程id
func getRId() int64 {
	var stackInfo [64]byte
	// 获取栈信息 n == 长度
	n := runtime.Stack(stackInfo[:], false)
	// 定位id，把id放在第一位
	rIds := strings.TrimPrefix(string(stackInfo[:n]), "goroutine")
	// 第一个就是值
	rId := strings.Fields(rIds)[0]
	// 转为64位整数
	id, _ := strconv.Atoi(rId)
	return int64(id)
}
