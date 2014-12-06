nofork
======

`nofork` is a tiny go app that doesn't do much more than starting a program and
writing a pidfile. The advantage though, is that you don't have to fork the
program and then retrieve the exit code with the shell (`$!`). This allows
signals to reach the process normally and simplifies wrapping a program with a
pidfile mechanism.

The intended use case is to enable signaling processes which may take a while
to run, such as builds.

Usage
=====

`nofork -pidfile /path/to/pidfile [-remove] <command>`

The pidfile location and command are required. If `-remove` is specified, the
pidfile will be automatically removed when the program exits. By default, the
pidfile will be left on the filesystem.
