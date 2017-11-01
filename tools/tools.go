package tools

import (
	"fmt"
	"os"
	logs "github.com/Suenaa/agenda-go/logs"
)

// Report 输出错误
func Report(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		logs.Log(err)
		os.Exit(2)
	}
}
