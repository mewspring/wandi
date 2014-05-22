package wandi

import (
	"image"
)

// A Drawable is a mutable collection of pixels.
type Drawable interface {
	// Draw draws the entire src image onto the dst image starting at the
	// destination point dp.
	Draw(dp image.Point, src Image) (err error)
	// DrawRect draws a subset of the src image, as defined by the source
	// rectangle sr, onto the dst image starting at the destination point dp.
	DrawRect(dp image.Point, src Image, sr image.Rectangle) (err error)
}
