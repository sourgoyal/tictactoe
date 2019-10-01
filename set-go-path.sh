#!/bin/bash

export GOPATH=`pwd`

current_path=`pwd`
cd src/tictactoe/
export GOBIN=`pwd`
cd $current_path
