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

/*
Write
*/

func (self *EntityChannel) WriteOnce(dataOnce interface{}) {
	err.ThrowError()
	channelWrite := self.getChannelWrite()
	channelWrite <- dataOnce
	self.closeChannel()
	return
}

func (self *EntityChannel) getChannelWrite() (channelWrite chan<- interface{}) {
	channelWrite = self.ChannelWrite
	return channelWrite
}

func (self *EntityChannel) closeChannel() {
	channelWrite := self.ChannelWrite
	close(channelWrite)
	self.checkIsClosed()
	return
}

func (self *EntityChannel) checkIsClosed() {
	channelRead := self.ChannelRead
	_, ok := <-channelRead
	if ok {
		msgError := fmt.Sprintf(
			"Fail to close channel",
			1)
		err.ErrPanic(msgError)
	}
}

/*
Read
*/

func (self *EntityChannel) ReadOnce() (dataOnce interface{}) {
	err.ThrowError()
	channelRead := self.getChannelRead()
	dataOnce = <-channelRead
	return dataOnce
}

func (self *EntityChannel) getChannelRead() (channelRead <-chan interface{}) {
	channelRead = self.ChannelRead
	return channelRead
}
