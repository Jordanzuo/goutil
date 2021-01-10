/*
This package is designed to simplify the use of csp model.
In csp model, we use channel to communicate among goroutines.
But it's not easy and error prone to implement the logic every time.
So, this package is designed as a template to use.
*/
package cspBaseModel

import (
	"context"
	"fmt"
)

// Request type is defined as the message exchanged between higher layer and the base model.
type Request struct {
	Name      string      // The identity of a request. Which is used to find the callback function.
	Parameter interface{} // The data sent to the callback function.
	ReturnCh  interface{} // A channel used to pass message between the request method and the callback function.
}

func NewRequest(name string, parameter, returnCh interface{}) *Request {
	return &Request{
		Name:      name,
		Parameter: parameter,
		ReturnCh:  returnCh,
	}
}

// BaseModel type is the core type
type BaseModel struct {
	callbackMap    map[string]func(*Request) // This callbackMap stores all the name and callback function pairs.
	RequestChannel chan *Request             // This channel is used to receive all the request from higher layer functions.
}

// Higher layer modules can register name and callback function pairs to base model.
func (this *BaseModel) Register(name string, callback func(*Request)) {
	if _, exists := this.callbackMap[name]; exists {
		panic(fmt.Sprintf("Item with name:%s has already existed.", name))
	}
	this.callbackMap[name] = callback
}

// Higher layer modules start base model to handle all the requests.
// ctx is used to stop this goroutine, to avoid goroutine leak.
func (this *BaseModel) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case requestObj := <-this.RequestChannel:
				if callback, exists := this.callbackMap[requestObj.Name]; !exists {
					panic(fmt.Sprintf("There is no callback related to %s", requestObj.Name))
				} else {
					callback(requestObj)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

func NewBaseModel() *BaseModel {
	return &BaseModel{
		callbackMap:    make(map[string]func(*Request), 16),
		RequestChannel: make(chan *Request, 64),
	}
}
