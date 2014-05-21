// Package wandiutil provides utility interfaces which facilitates the
// composition of wandi interfaces.
package wandiutil

import (
	"image"
	"image/color"
)

// Filler is the interface that wraps the Fill method.
type Filler interface {
	// Fill fills the entire image with the provided color.
	Fill(c color.Color)
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
