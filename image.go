package wandi

import (
	"image"
)

// An Image is a mutable collection of pixels.
type Image interface {
	// Width returns the width of the image.
	Width() int
	// Height returns the height of the image.
	Height() int
	// Draw draws the entire src image onto the dst image starting at the
	// destination point dp.
	Draw(dp image.Point, src Image) (err error)
	// DrawRect fills the destination rectangle dr of the dst image with
	// corresponding pixels from the src image starting at the source point sp.
	DrawRect(dr image.Rectangle, src Image, sp image.Point) (err error)
}
