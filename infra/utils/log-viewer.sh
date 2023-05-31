#!/bin/sh

less ${1}"$(ls -r "$1" | fzf --preview "cat ${1}/{} | less" --preview-window right:75%)"
