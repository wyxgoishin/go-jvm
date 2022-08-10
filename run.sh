#!/bin/zsh
target=$1
output=$2
./main --cp /Users/yixinwu/go-jvm/resource --Xjre /Users/yixinwu/jdk1.8.0_332.jdk/Contents/Home/jre --verbose $target > output