package fileUtil

import (
	"io"
	"net/http"
	"os"
	"path"
)

// 下载网络文件
// netUrl：网络文件地址
// saveDir：存储位置
// saveFileName:存储的文件名
// ifTruncate:如果文件存在了，是否覆盖此文件
// 返回值：
// err:错误对象
func DownLoadNetFile(netUrl string, saveDir string, saveFileName string, ifTruncate bool) (err error) {
	resp, err := http.Get(netUrl)
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		return
	}

	// 创建文件夹
	if IsDirExists(saveDir) == false {
		os.MkdirAll(saveDir, os.ModePerm|os.ModeTemporary)
	}

	// 创建文件
	filePath := path.Join(saveDir, saveFileName)
	var fileObj *os.File
	if ifTruncate {
		fileObj, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm|os.ModeTemporary)
	} else {
		// 如果文件已经存在，则不能打开
		fileObj, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_EXCL, os.ModePerm|os.ModeTemporary)
	}
	defer func() {
		if fileObj != nil {
			fileObj.Close()
		}
	}()
	if err != nil {
		return
	}

	// 写入文件数据
	_, err = io.Copy(fileObj, resp.Body)

	return
}
