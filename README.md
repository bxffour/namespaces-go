namespaces-go
=================

namespaces-go used namespaces to run a shell in a sandboxed environment. This is a purely
educational project as a result of my pursuit to understanding to fundamental cogs that 
make containers possible.

Demo
--------------

![namespaces-go](https://vhs.charm.sh/vhs-3cg3qwjrdVCdHXrylLKizb.gif)

Installation
--------------

    $ git clone https://github.com/bxffour/namespaces-go.git
    $ cd namespaces-go/

Build the binary

    $ make build

Preparing rootfs

    $ make assets/setup

Usage
--------------

Sandbox hardcoded rootfs

    $ ./bin/ns

Sandbox rootfs provided through cmdline args

    $ ./bin/ns /path/to/rootfs

Credit
---------------

1. Namespaces in Go blog by Ed King - https://medium.com/@teddyking/namespaces-in-go-basics-
2. Liz Rice's demo on containers - https://youtu.be/8fi7uSYlOdc
