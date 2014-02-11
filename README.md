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

The current implementations of [wandi.Image][] are listed below:

- [SDL][] backend
   - [texture][sdl/texture]: not yet implemented.
- [SFML][] backend
   - [texture][sfml/texture]: handles hardware accelerated image drawing.

[wandi.Image]: http://godoc.org/github.com/mewmew/wandi#Image
[SDL]: http://www.libsdl.org/
[SFML]: http://www.sfml-dev.org/
[sdl/texture]: http://godoc.org/github.com/mewmew/sdl/texture
[sfml/texture]: http://godoc.org/github.com/mewmew/sfml/texture

The current implementations of [wandi.Window][] are listed below:

- [SDL][] backend
   - [window][sdl/window]: not yet implemented.
- [SFML][] backend
   - [window][sfml/window]: handles window creation, drawing and events.

[wandi.Window]: http://godoc.org/github.com/mewmew/wandi#Window
[sdl/window]: http://godoc.org/github.com/mewmew/sdl/window
[sfml/window]: http://godoc.org/github.com/mewmew/sfml/window

public domain
-------------

This code is hereby released into the *[public domain][]*.

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
