#!/bin/sh

export CGO_CFLAGS="$(php-config --includes) -DLOGLEVEL=error"
go build  -buildmode=c-shared -o sse2_strlen.so sse2_strlen.go
