package main

import (
	"fmt"
	"os"
	"strings"
)

var Ai *AiClient

func main() {
	var tools = []string{
		"[q 'question'] 任意问题",
		"[fy 'text'] 翻译中文",
		"[fyy 'text'] 翻译英文",
		"[mm 'desc'] 变量命名",
		"[go] 推荐golang仓库",
		"[js] 推荐javascript仓库",
		"[dart] 推荐dart仓库",
		"[vue] 推荐js vue仓库",
		"[bug 'error'] 帮助解决程序报错",
		"[sql 'desc'] 根据描述生成sql语句",
	}
	if len(os.Args) < 2 {
		fmt.Println("请提供命令,目前支持:\n", strings.Join(tools, "\n"))
		return
	}
	var err error
	Ai, err = NewClient()
	if err != nil {
		fmt.Println("初始化openai客户端失败:", err)
		return
	}

	switch os.Args[1] {
	case "q":
		Q(os.Args[2:])
	case "fy":
		FY(os.Args[2:])
	case "fyy":
		FYY(os.Args[2:])
	case "mm":
		MM(os.Args[2:])
	case "go":
		Top(os.Args[2:], "go")
	case "js":
		Top(os.Args[2:], "javascript")
	case "dart":
		Top(os.Args[2:], "dart")
	case "vue":
		Top(os.Args[2:], "vue")
	case "bug":
		Debug(os.Args[2:])
	case "sql":
		Sql(os.Args[2:])
	default:
		fmt.Println("未能识别的命令:"+os.Args[1]+",目前支持:", tools)
	}
}
