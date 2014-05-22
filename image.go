package wandi

// An Image is a rectangular collection of pixels.
type Image interface {
	// Width returns the width of the image.
	Width() int
	// Height returns the height of the image.
	Height() int
}
