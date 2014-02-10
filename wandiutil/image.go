// Package wandiutil provides utility interfaces which facilitates interface
// composition of wandi images.
package wandiutil

import (
	"image"
	"image/color"
)

// Clearer is the interface that wraps the Clear method.
type Clearer interface {
	// Clear clears the image and fills it with the provided color.
	Clear(c color.Color)
}

// Freer is the interface that wraps the Free method.
type Freer interface {
	// Free frees the image.
	Free()
}

// Imager is the interface that wraps the Image method.
type Imager interface {
	// Image returns an image.Image representation of the image.
	Image() (img image.Image, err error)
}
