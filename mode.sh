#!/usr/bin/bash
# Author: Tommy Chu
# Toggle release and debug mode for Gin.

if [ "$GIN_MODE" = "debug" ]
then
    export GIN_MODE="release"
    echo "[GIN MODE] RELEASE"
else
    export GIN_MODE="debug"
    echo "[GIN MODE] DEBUG"
fi
