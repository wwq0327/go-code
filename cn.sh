#!/bin/bash 
# 统计代码行数
find . -name '*.go' |xargs cat| wc -l
