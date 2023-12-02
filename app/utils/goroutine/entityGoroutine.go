package goroutine

import (
	"fmt"
	"github.com/kanelinweihao/lwhFrameGo/app/utils/err"
)

type EntityChannel struct {
	ChannelBase  chan interface{}
	ChannelWrite chan<- interface{}
	ChannelRead  <-chan interface{}
}

func (self *EntityChannel) WriteOnce(dataOnce interface{}) {
	err.ThrowError()
	channelWrite := self.ChannelWrite
	channelWrite <- dataOnce
	close(channelWrite)
	channelRead := self.ChannelRead
	_, ok := <-channelRead
	if ok {
		msgError := fmt.Sprintf(
			"Fail to close channel |%s|",
			"interface{}")
		err.ErrPanic(msgError)
	}
	return
}

func (self *EntityChannel) ReadOnce() (dataOnce interface{}) {
	err.ThrowError()
	channelRead := self.ChannelRead
	dataOnce = <-channelRead
	return dataOnce
}
