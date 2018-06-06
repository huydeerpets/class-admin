#!/bin/sh

if [ -f ".server_ca.pid" ]; then
  kill  `cat .server_ca.pid`
  rm -rf ".server_ca.pid"
fi
