package ensureSendUtil

import (
	"fmt"
	"time"

	"github.com/Jordanzuo/Framework/goroutineMgr"
	"github.com/Jordanzuo/Framework/monitorMgr"
	"github.com/Jordanzuo/goutil/debugUtil"
	"github.com/Jordanzuo/goutil/logUtil"
)

// 负责发送数据的协程
func sendLoop(s sender, closeSignal chan struct{}) {
	name := "ensureSendUtil.send.sendLoop"
	goroutineMgr.MonitorZero(name)
	defer goroutineMgr.ReleaseMonitor(name)

	for {
		select {
		case <-s.Done():
			closeSignal <- struct{}{}
			return
		case v := <-s.Data():
			if err := s.Send(v); err != nil {
				// 发送失败存入缓存
				s.Cache() <- v
			}
		}
	}
}

// 定时重发失败的数据
func resendLoop(s sender, folder string, closeSignal chan struct{}) {
	name := "ensureSendUtil.send.resendLoop"
	goroutineMgr.MonitorZero(name)
	defer goroutineMgr.ReleaseMonitor(name)

	// debug模式每秒重试1次
	var delay time.Duration
	if debugUtil.IsDebug() {
		delay = time.Second
	} else {
		delay = time.Minute * 5
	}

	// 定时重发失败数据
	for {
		select {
		case <-s.Done():
			closeSignal <- struct{}{}
			return
		case <-time.After(delay):
			sendCacheData(s, folder)
			loadData(s.(EnsureSender), folder)
		}
	}
}

// 从sender获取失败数据重发
func sendCacheData(s sender, folder string) {
	failed := make([]dataItem, 0)
	length := len(s.Cache())

	defer func() {
		// 用于记录多次失败后放弃发送的数据
		giveUpItems := make(chan dataItem, len(failed))

		for _, v := range failed {
			if v.Count() >= 3 {
				// 失败次数太多的数据准备存放到磁盘中
				giveUpItems <- v
			} else {
				s.Cache() <- v
			}
		}

		giveUpLen := len(giveUpItems)
		if giveUpLen > 0 {
			// 将多次失败的数据保存到文件中
			if folder[len(folder)-1] == '/' {
				folder = folder[:len(folder)-1]
			}
			saveData(giveUpItems, folder+"_giveup")

			if giveUpLen >= 5 {
				log := fmt.Sprintf("ensureSendUtil: 有%d条数据多次发送失败", giveUpLen)
				logUtil.NormalLog(log, logUtil.Error)
				monitorMgr.Report(log)
			}
		}

		// 输出信息
		log := fmt.Sprintf("ensureSendUtil: 重发%d条数据，失败%d条，存盘%d条\n", length, len(failed), giveUpLen)
		logUtil.NormalLog(log, logUtil.Info)
	}()

	for {
		select {
		case v := <-s.Cache():
			// 重发数据
			if e := s.Send(v); e != nil {
				// 记录失败的数据
				failed = append(failed, v)
			}
		default:
			return
		}
	}
}
