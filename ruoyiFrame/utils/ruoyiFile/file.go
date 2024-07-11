package ruoyiFile

import "os"

func IsFileExist(addr string) bool {
	_, err := os.Stat(addr) //获取文件信息
	if err != nil {
		return false
	}
	return true
}
