package ensureSendUtil

import (
	"fmt"

	"github.com/Jordanzuo/goutil/fileUtil"
	"github.com/Jordanzuo/goutil/logUtil"
	"github.com/Jordanzuo/goutil/stringUtil"
)

// 从目录加载缓存数据并发送
func loadData(s EnsureSender, folder string) error {
	if fileList, err := fileUtil.GetFileList(folder); err != nil {
		return err
	} else {
		for _, filename := range fileList {
			// 读取发送内容
			if fileContent, err := fileUtil.ReadFileContent(filename); err != nil {
				// 打印错误
				log := fmt.Sprintf("ensureSendUtil.loadData: Failed To Read File: %s %s\n", err, filename)
				logUtil.NormalLog(log, logUtil.Error)
			} else if err = fileUtil.DeleteFile(filename); err != nil {
				// 删除文件，如果成功则将内容添加到通道中，否则不处理
				log := fmt.Sprintf("ensureSendUtil.loadData: Failed To Delete File: %s %s", err, filename)
				logUtil.NormalLog(log, logUtil.Error)
			} else {
				// 发送数据
				s.Write(fileContent)
			}
		}
	}

	return nil
}

// 保存数据到文件中(通常在退出时调用)
func saveData(datas <-chan dataItem, folder string) (failed []dataItem, err error) {
	defer func() {
		if len(failed) > 0 {
			err = fmt.Errorf("保存数据时有%d个失败数据", len(failed))
		}
	}()

	for {
		select {
		case v := <-datas:
			filename := stringUtil.GetNewGUID()
			if e := fileUtil.WriteFile(folder, filename, false, v.String()); e != nil {
				failed = append(failed, v)
				log := fmt.Sprintf("ensureSendUtil.saveData: 写入错误\n目录：%s，文件：%s，错误信息为：%s, Data:%s",
					folder, filename, err, v.String())
				logUtil.NormalLog(log, logUtil.Error)
			}
		default:
			return
		}
	}
}
