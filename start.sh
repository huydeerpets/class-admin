#!/bin/sh

export PATH="$PATH:$GOPATH/bin"

OLD_GOPATH=$GOPATH

scriptPath=$(cd `dirname $0`; pwd)
cd $scriptPath/../../

NEW_GOPATH=$('pwd')
export GOPATH=$OLD_GOPATH:$NEW_GOPATH

cd src

if [ -f ".server_ca.pid" ]; then
  kill -9 `cat .server_ca.pid`
  rm -rf ".server_ca.pid"
fi

echo "==========run=========="
if [ -f "nohup-ca.out" ]; then
  rm -rf "nohup-ca.out"
fi

nohup bee run class-wechat

export GOPATH=$OLD_GOPATH
