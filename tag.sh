#!/bin/bash
echo "用法: ./tag.sh v1.0.27 feat:新增xx功能"
git add .
git commit -m $2
git push origin master
git tag -a $1 -m $2
git push origin $1