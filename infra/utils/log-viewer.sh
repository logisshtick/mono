#!/bin/sh

ls -r "$1" | fzf --preview "cat ${1}{} | less" 
