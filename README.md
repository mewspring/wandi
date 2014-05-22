WIP
---

This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

wandi
=====

This package specifies common interfaces used for window creation, event
handling and image drawing. It is heavily inspired by [go.wde][].

[go.wde]: https://github.com/skelterjohn/go.wde

Documentation
-------------

Documentation provided by GoDoc.

- [wandi][]: specifies common interfaces used for window creation, drawing and
event handling.
   - [wandiutil][wandi/wandiutil]: provides utility interfaces which facilitates
   the composition of wandi interfaces.

[wandi]: http://godoc.org/github.com/mewmew/wandi
[wandi/wandiutil]: http://godoc.org/github.com/mewmew/wandi/wandiutil

Implementations
---------------

The current implementations of [wandi.Drawable][] are listed below:

- SDL backend
   - [texture.Drawable][sdl/texture#Drawable]: not yet implemented.
- SFML backend
   - [texture.Drawable][sfml/texture#Drawable]: drawable GPU texture.

[wandi.Drawable]: http://godoc.org/github.com/mewmew/wandi#Drawable
[sdl/texture#Drawable]: http://godoc.org/github.com/mewmew/sdl/texture#Drawable
[sfml/texture#Drawable]: http://godoc.org/github.com/mewmew/sfml/texture#Drawable

The current implementations of [wandi.Image][] are listed below:

- SDL backend
   - [texture.Image][sdl/texture#Image]: not yet implemented.
- SFML backend
   - [texture.Image][sfml/texture#Image]: read-only GPU texture.

[wandi.Image]: http://godoc.org/github.com/mewmew/wandi#Image
[sdl/texture#Image]: http://godoc.org/github.com/mewmew/sdl/texture#Image
[sfml/texture#Image]: http://godoc.org/github.com/mewmew/sfml/texture#Image

The current implementations of [wandi.Window][] are listed below:

- SDL backend
   - [window.Window][sdl/window#Window]: not yet implemented.
- SFML backend
   - [window.Window][sfml/window#Window]: graphical window capable of handling
   draw operations and window events.

[wandi.Window]: http://godoc.org/github.com/mewmew/wandi#Window
[sdl/window#Window]: http://godoc.org/github.com/mewmew/sdl/window#Window
[sfml/window#Window]: http://godoc.org/github.com/mewmew/sfml/window#Window

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
