#!/bin/sh
if [ $dev = "1" ];
then
    echo Development mode.
    go get github.com/codegangsta/gin
    gin run
else
    ./app
fi
