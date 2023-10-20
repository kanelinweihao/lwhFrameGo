package goroutine

import (
	"fmt"
	_ "go.lwh.com/linweihao/lwhFrameGo/app/utils/dd"
	"go.lwh.com/linweihao/lwhFrameGo/app/utils/err"
)

type EntityChannel struct {
	ChannelBase  chan interface{}
	ChannelWrite chan<- interface{}
	ChannelRead  <-chan interface{}
}

////
// Init
////

func InitEntityChannel() (entityChannel *EntityChannel) {
	channelBase := make(chan interface{})
	var channelWrite chan<- interface{}
	channelWrite = channelBase
	var channelRead <-chan interface{}
	channelRead = channelBase
	entityChannel = &EntityChannel{
		ChannelBase:  channelBase,
		ChannelWrite: channelWrite,
		ChannelRead:  channelRead,
	}
	return entityChannel
}

////
// Exec
////

func (self *EntityChannel) WriteOnce(dataOnce interface{}) {
	channelWrite := self.getChannelWrite()
	channelWrite <- dataOnce
	self.closeChannel()
	return
}

func (self *EntityChannel) ReadOnce() (dataOnce interface{}) {
	channelRead := self.getChannelRead()
	dataOnce = <-channelRead
	return dataOnce
}

func (self *EntityChannel) getChannelWrite() (channelWrite chan<- interface{}) {
	channelWrite = self.ChannelWrite
	return channelWrite
}

func (self *EntityChannel) getChannelRead() (channelRead <-chan interface{}) {
	channelRead = self.ChannelRead
	return channelRead
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
