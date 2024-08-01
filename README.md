# TerminalBot


## Introduction
terminalBot 是一个终端机器人， 可以将chatGPT接入你的终端，以最快的方式提问。

## User Guide
1. 安装依赖编译 `go build *.go`
2. 添加 ai 二进制文件到环境变量
3. 设置环境变量`OPENAI_API_KEY`和`OPENAI_API_URL`
4. 终端输入 `ai  q "你的问题"`,除了q 外，还可以输入 `ai  mm 课程` 帮我们编程取变量名，还有其他的工具自行探索。

## cmd list
- [q 'question'] 任意问题
- [fy 'text'] 翻译中文
- [fyy 'text'] 翻译英文
- [mm 'desc'] 变量命名
- [go] 推荐golang仓库
- [js] 推荐javascript仓库
- [dart] 推荐dart仓库
- [vue] 推荐js vue仓库
- [bug 'error'] 帮助解决程序报错
- [sql 'desc'] 根据描述生成sql语句


## Contribution
Feel free to open issues or pull requests if you have any suggestions or found any bugs.

## License
This project is licensed under the MIT License.
