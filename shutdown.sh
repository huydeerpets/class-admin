#!/bin/sh

if [ -f ".class-admin.pid" ]; then
  kill  `cat .class-admin.pid`
  rm -rf ".class-admin.pid"
fi
