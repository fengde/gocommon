#!/bin/bash

git add .
git commit -m $1
git push origin master
git tag -a $1 -m $2
git push origin $1