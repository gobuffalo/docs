#!/bin/bash
set +x

git checkout -b build
GOOS=linux buffalo build -z
git add .
git commit -a -m "binary commit"
git push heroku build:master
git checkout master
git branch -d build
rm -rf bin/
