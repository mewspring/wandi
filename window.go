// Package wandi specifies common interfaces used for window creation, event
// handling and image drawing.
package wandi

import (
	"image/color"

	"github.com/mewmew/we"
)

// A Window represents a graphical window capable of handling draw operations
// and window events.
type Window interface {
	// Close closes the window.
	Close()
	// SetTitle sets the title of the window.
	SetTitle(title string)
	// The Image interface is implemented by the window.
	Image
	// Update displays window rendering updates on the screen.
	Update() (err error)
	// Clear clears the window and fills it with the provided color.
	Clear(c color.Color)
	// PollEvent returns a pending event from the event queue or nil if the queue
	// was empty. Note that more than one event may be present in the event
	// queue.
	PollEvent() (event we.Event)
}
