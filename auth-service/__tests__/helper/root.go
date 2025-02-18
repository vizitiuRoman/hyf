package helper

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func getRootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	const module = "__tests__"

	if !strings.Contains(dir, module) {
		panic(errors.New(fmt.Sprintf("dir %s does not contain module %s", dir, module)))
	}

	root := string([]rune(dir)[:strings.Index(dir, module)+len(module)])
	return root
}
