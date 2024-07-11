package ruoyiFile

import (
	"fmt"
	"os"
	"testing"
)

func Test_IsFileExist(t *testing.T) {
	FileInfo, err := os.Stat("./ruoyiFile") // 获取文件信息
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("FileInfo", FileInfo)
}
