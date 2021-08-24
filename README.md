# wandi

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mewspring/wandi)

The wandi project specifies common interfaces used for window creation, event handling and image drawing. It is heavily inspired by [go.wde](https://github.com/skelterjohn/go.wde).

## Implementations

The current implementations of [wandi.Drawable](http://pkg.go.dev/github.com/mewspring/wandi#Drawable) are listed below:

- SDL backend
   - [texture.Drawable](http://pkg.go.dev/github.com/mewspring/sdl/texture#Drawable): not yet implemented.
- SFML backend
   - [texture.Drawable](http://pkg.go.dev/github.com/mewspring/sfml/texture#Drawable): drawable GPU texture.

The current implementations of [wandi.Image](http://pkg.go.dev/github.com/mewspring/wandi#Image) are listed below:

- SDL backend
   - [texture.Image](http://pkg.go.dev/github.com/mewspring/sdl/texture#Image): not yet implemented.
- SFML backend
   - [texture.Image](http://pkg.go.dev/github.com/mewspring/sfml/texture#Image): read-only GPU texture.

The current implementations of [wandi.Window](http://pkg.go.dev/github.com/mewspring/wandi#Window) are listed below:

- SDL backend
   - [window.Window](http://pkg.go.dev/github.com/mewspring/sdl/window#Window): not yet implemented.
- SFML backend
   - [window.Window](http://pkg.go.dev/github.com/mewspring/sfml/window#Window): graphical window capable of handling
   draw operations and window events.
