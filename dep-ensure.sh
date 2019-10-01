#!/bin/bash

source set-go-path.sh

current_path=`pwd`
cd src/tictactoe/ 
env DEPNOLOCK=1 dep ensure -v
cd $current_path
