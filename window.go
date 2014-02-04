// Package wde specifies common interfaces used for window creation, drawing and
// event handling.
package wde

import (
	"image/color"
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
	// Update displays window updates on the screen.
	Update() (err error)
	// Clear clears the screen and fills it with the provided color.
	Clear(c color.Color) (err error)
	// PollEvent returns a pending event from the event queue or nil if the queue
	// was empty.
	PollEvent() (event Event)
}
