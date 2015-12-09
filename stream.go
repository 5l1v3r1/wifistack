package wifistack

import "github.com/unixpickle/gofi"

// A Stream is an abstract interface which can transfer 802.11 frames.
type Stream interface {
	// Incoming returns the channel to which incoming packets are delivered.
	// When the stream is closed, this channel will be closed as well.
	Incoming() <-chan gofi.RadioPacket

	// Outgoing returns the channel to which outgoing packets should be sent.
	// You can close() this channel to close the stream.
	Outgoing() chan<- gofi.Frame

	// SupportedChannels returns a list of all supported WLAN channels
	// in no particular order.
	SupportedChannels() []gofi.Channel

	// Channel returns the WLAN channel to which the stream is tuned.
	Channel() gofi.Channel

	// SetChannel tunes the stream into a WLAN channel.
	// Some types of Stream may not support channel hopping, in which case
	// this will return an error.
	SetChannel(c gofi.Channel) error

	// FirstError returns the first read or write error that this Stream
	// encountered.
	// If the stream was closed without the stream's user explicitly
	// closing it, this must return a non-nil value.
	FirstError() error
}
