package examples

import (
	"fmt"
	"os"
	"strings"
)

// CheckArgs should be used to ensure the right command line arguments are passed before executing an example.
// 应使用 CheckArgs 确保在执行示例之前传递正确的命令行参数。
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		Warning("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
// 如果错误不为零，则应使用 CheckIfError 来天真地发生恐慌。
func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
// 信息应用于描述即将运行的示例命令。
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
// warning应该用来显示警告
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}
