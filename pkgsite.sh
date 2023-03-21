#!/bin/bash

pid=$(ps x | grep 'pkgsite' | grep -v grep | awk '{print $1}')
kill $pid

pkgsite -http=localhost:6060
#Then browse http://localhost:6060/github.com/rbtyang/godash