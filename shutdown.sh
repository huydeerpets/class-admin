#!/bin/sh

if [ -f ".class-admin.pid" ]; then
  kill  `cat .class-ca.pid`
  rm -rf ".class-ca.pid"
fi
