package goroutine

func MakeEntityChannel() (entityChannel *EntityChannel) {
	entityChannel = new(EntityChannel)
	entityChannel.Init()
	return entityChannel
}

func (self *EntityChannel) Init() *EntityChannel {
	channelBase := make(chan interface{})
	self.ChannelBase = channelBase
	self.ChannelWrite = channelBase
	self.ChannelRead = channelBase
	return self
}
