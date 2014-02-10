// Package wandiutil specifies utility interfaces for image drawing.
package wandiutil

import (
	"image/color"

	"github.com/mewmew/wandi"
)

// ImageFreer is the interface that groups the wandi.Image interface with the
// Free method.
type ImageFreer interface {
	wandi.Image
	// Free frees the image.
	Free()
}

// ImageClearer is the interface that groups the wandi.Image interface with the
// Clear method.
type ImageClearer interface {
	wandi.Image
	// Clear clears the image and fills it with the provided color.
	Clear(c color.Color)
}

// ImageClearFreer is the interface that groups the wandi.Image interface with
// the Clear and Free methods.
type ImageClearFreer interface {
	wandi.Image
	// Clear clears the image and fills it with the provided color.
	Clear(c color.Color)
	// Free frees the image.
	Free()
}
