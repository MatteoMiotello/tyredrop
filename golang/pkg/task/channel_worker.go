package task

// ChannelWorker worker that spawn maxWorker goroutines and consume a channel given.
type ChannelWorker[T any] struct {
	maxWorker int
	channel   chan T
}

// NewChannelWorker create a channel worker with maxWorker and channel
func NewChannelWorker[T any](maxWorker int, channel chan T) *ChannelWorker[T] {
	return &ChannelWorker[T]{
		maxWorker: maxWorker,
		channel:   channel,
	}
}

func (c *ChannelWorker[T]) Run(handler func(c T)) {
	for i := 0; i < 24; i++ {
		go func() {
			for c := range c.channel {
				handler(c)
			}
		}()
	}
}

func (c *ChannelWorker[T]) InsertToChannel(element T) {
	c.channel <- element
}
