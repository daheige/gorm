package gorm

import (
	"log"
	"testing"

	"gorm.io/gorm/utils"
)

func TestReplace(t *testing.T) {
	log.Println("FileWithLineNum: ", utils.FileWithLineNum())
}

/**
2021/05/31 21:01:20 gormSourceDir:  /Users/heige/web/go/gorm
=== RUN   TestReplace
2021/05/31 21:01:20 file: /Users/heige/web/go/gorm/utils_filenum_test.go  line:  11 ok:  true
2021/05/31 21:01:20 FileWithLineNum:  /Users/heige/web/go/gorm/utils_filenum_test.go:11
--- PASS: TestReplace (0.00s)
PASS
*/
