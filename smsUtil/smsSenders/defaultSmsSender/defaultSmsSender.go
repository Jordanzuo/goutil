package defaultSmsSender

import (
    "github.com/Jordanzuo/goutil/smsUtil/dataSender"
    "github.com/Jordanzuo/goutil/smsUtil/dataSenders/httpSender"
)


type defaultSmsSender struct {
    sender dataSender.Sender
}

func (this *defaultSmsSender) Send(msg dataSender.Requester) (bool, error) {
    rspn, err := this.sender.Send(msg)
    if err != nil {
        return false, err
    }

    ok, err := msg.ParseResponse(rspn)

    return ok, err
}

func New() (*defaultSmsSender) {
    return &defaultSmsSender{sender: httpSender.New()}
}
