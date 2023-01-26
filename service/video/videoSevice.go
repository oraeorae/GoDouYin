package video

import (
	"fmt"
	"go_douyin/global/variable"
	"io"
	"mime/multipart"
	"os"
)

// 保存文件
func SaveFile(file multipart.File, header *multipart.FileHeader) bool {
	filename := header.Filename
	fmt.Println(header.Filename)
	out, err := os.Create(variable.BasePath + "/config/" + filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
