package ensureSendUtil

/*
ensureSendUtil 用于推送数据
支持TCP和HTTP两种形式，在发送失败时会缓存数据，并在一定时间间隔后重试

通过NewTCPSender和NewHTTPSender两个接口分别创建TCP和HTTP模式的EnsureSender

type EnsureSender interface {
    // 用于发送数据
    Write(string) error

    // 用于停止发送，此时会自动保存未发送数据
    Close() error
}

// 创建一个tcp数据发送器
// 参数：
//      _dataFolder  数据存放目录
//      _address     连接地址
func NewTCPSender(_dataFolder, _address string) (EnsureSender, error) {


// 创建一个http数据发送器
// 参数：
//      _dataFolder  数据存放目录
//      _url         发送地址
func NewHTTPSender(_dataFolder, _url string) (EnsureSender, error) {
*/
