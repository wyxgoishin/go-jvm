#!/bin/zsh
srcPwd=$(pwd)
dstPwd=$1
class=$2
cd $dstPwd
~/jdk1.8.0_332.jdk/Contents/Home/bin/java $2
cd $srcPwd