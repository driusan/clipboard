[![Build Status](https://travis-ci.org/atotto/clipboard.svg?branch=master)](https://travis-ci.org/atotto/clipboard) [![Build Status](https://drone.io/github.com/atotto/clipboard/status.png)](https://drone.io/github.com/atotto/clipboard/latest) 

[![GoDoc](https://godoc.org/github.com/atotto/clipboard?status.svg)](http://godoc.org/github.com/atotto/clipboard)

# Clipboard for Go

This is a fork of atotto/clipboard which adds support for DragonFly and Plan9.

(atotto/clipboard does not appear to be maintained, so this is just here so I can
use for de. I don't expect any significant work on this other than adding new OS
support as needed.)

Provide copying and pasting to the Clipboard for Go.

Download shell commands at https://drone.io/github.com/atotto/clipboard/files

Build:

    $ go get github.com/driusan/clipboard

Platforms:

* OSX
* Windows 7 (probably work on other Windows)
* Linux, Unix (requires 'xclip' or 'xsel' command to be installed)
* Plan 9


Document: 

* http://godoc.org/github.com/driusan/clipboard

Notes:

* Text string only
* UTF-8 text encoding only (no conversion)

TODO:

## Commands:

paste shell command:

    $ go get github.com/atotto/clipboard/cmd/gopaste
    $ # example:
    $ gopaste > document.txt

copy shell command:

    $ go get github.com/atotto/clipboard/cmd/gocopy
    $ # example:
    $ cat document.txt | gocopy



