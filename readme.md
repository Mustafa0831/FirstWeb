# Description

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of ascii-art project.

# Authors

Mustafa(mus11110), Zhangir(zhangir11)

# Usage

In terminal type and run command (in root path):
``` go run .
```
To go main page, type in Browser URL:"localhost:8001", and press enter.

# Implementation details

Server listens 8001's port. Handles only two requests "/" and "ascii-art".
When client make request "ascii-art" by method Post, program gets text and fontStyle from page elements and gives to GetASCII func.
This func will return string "text" in ASCII-art
Used Languages: HTML, Golang

#