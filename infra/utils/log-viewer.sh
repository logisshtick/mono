#!/bin/sh

ls "$1" | fzf --preview "cat ${1}{} | less" --layout=reverse
