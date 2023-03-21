#!/bin/bash

pid=$(ps x | grep 'pkgsite' | grep -v grep | awk '{print $1}')
kill $pid

echo 'Godash pkgsite is: http://localhost:6060/github.com/rbtyang/godash'
pkgsite -http=localhost:6060