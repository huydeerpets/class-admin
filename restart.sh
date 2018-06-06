#!/bin/sh
OLD_GOPATH=$GOPATH

path=$(cd `dirname $0`; pwd)
cd $scriptPath/../../

NEW_GOPATH=$('pwd')
export GOPATH=$OLD_GOPATH:$NEW_GOPATH

cd src/class-admin

if [ -f ".class-admin.pid" ]; then
  kill  `cat .class-admin.pid`
  rm -rf ".class-admin.pid"
fi

if [ ! -d "logs" ]; then 
  mkdir logs
fi

go install

export GO_MODE="test"

binName=class-admin
nohup bin/$binName > logs/$binName.out 2>&1 &

export GOPATH=$OLD_GOPATH
