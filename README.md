namespaces-go
=================

Utilizes linux namespaces to sandbox a given rootfs and open a shell.

![namespaces-go](https://github.com/bxffour/namespaces-go/blob/main/assets/gif/ns.gif)

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
