#!/bin/bash

set -ex

git checkout -b build
GOOS=linux buffalo build
git add .
git commit -a -m "binary commit"
git push heroku build:master --force
git checkout master
git branch -D build
rm -rf bin/
