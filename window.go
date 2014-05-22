// Package wandi specifies common interfaces used for window creation, event
// handling and image drawing.
package wandi

import (
	"github.com/mewmew/we"
)

// A Window represents a graphical window capable of handling draw operations
// and window events.
type Window interface {
	// Close closes the window.
	Close()
	// SetTitle sets the title of the window.
	SetTitle(title string)
	// Width returns the width of the window.
	Width() int
	// Height returns the height of the window.
	Height() int
	// The Drawable interface is implemented by the window.
	Drawable
	// Display displays what has been rendered so far to the window.
	Display()
	// PollEvent returns a pending event from the event queue or nil if the queue
	// was empty. Note that more than one event may be present in the event
	// queue.
	PollEvent() (event we.Event)
}
