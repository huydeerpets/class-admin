#!/bin/sh
OLD_GOPATH=$GOPATH

scriptPath=$(cd `dirname $0`; pwd)
cd $scriptPath/../../

NEW_GOPATH=$('pwd')
export GOPATH=$OLD_GOPATH:$NEW_GOPATH

if [ -f ".server-ca.pid" ]; then
  kill  `cat .server-ca.pid`
  rm -rf ".server-ca.pid"
fi

if [ ! -d "logs" ]; then
  mkdir logs
fi

go install  class-admin

export GO_MODE="test"

binName=class-admin
nohup bin/$binName > logs/$binName.out 2>&1 &

export GOPATH=$OLD_GOPATH
