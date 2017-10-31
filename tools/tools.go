package tools

import (
	"fmt"
	"os"
)

// Report 输出错误
func Report(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}
}
