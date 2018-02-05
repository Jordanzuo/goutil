package ensureSendUtil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/Jordanzuo/goutil/debugUtil"
)

// 保存接收的数据用于校验
var http_recv_msg = make([]byte, 0)

func init() {
	debugUtil.SetDebug(true)
}

type httpHandler struct {
	cnt int
}

func (ctx *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	result, _ := ioutil.ReadAll(r.Body)

	if string(result) == "http-msg-failed" {
		http.NotFound(w, r)
	} else {
		ctx.cnt++
		// 模拟一次失败
		if ctx.cnt == 2 {
			http.NotFound(w, r)
		} else {
			http_recv_msg = append(http_recv_msg, result...)
		}
	}

}

func Test_http(t *testing.T) {
	http.Handle("/test", new(httpHandler))
	go http.ListenAndServe("127.0.0.1:9560", nil)

	httpSender, err := NewHTTPSender("./test_http", "http://127.0.0.1:9560/test")
	if err != nil {
		t.Error(err)
	}

	time.Sleep(time.Millisecond * 50)

	// 第一次应该成功
	httpSender.Write("http-msg-1")

	time.Sleep(time.Millisecond)

	// 发送消息，此数据会多次失败，被丢弃到giveup目录
	httpSender.Write("http-msg-failed")

	time.Sleep(time.Second * 4)

	// 第二次应该失败
	httpSender.Write("http-msg-2")

	time.Sleep(time.Millisecond)

	// 保存数据
	httpSender.Close()

	// 重启之后应该会重发数据
	httpSender, err = NewHTTPSender("./test_http", "http://127.0.0.1:9560/test")
	if err != nil {
		t.Error(err)
	}
	time.Sleep(time.Second * 2)

	httpSender.Close()

	if string(http_recv_msg) != "http-msg-1http-msg-2" {
		t.Error("message error. got " + string(http_recv_msg))
	} else {
		fmt.Println("HTTP OK")
	}
}
