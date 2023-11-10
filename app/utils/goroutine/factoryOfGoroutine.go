package goroutine

/*
Init
*/

func MakeEntityOfGoroutine() (entityChannel *EntityChannel) {
	entityChannel = initEntityChannel()
	return entityChannel
}

func initEntityChannel() (entityChannel *EntityChannel) {
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
