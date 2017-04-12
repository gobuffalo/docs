#!/bin/bash

set -ex

rm -rfv public/assets
git checkout -b build
GOOS=linux buffalo build -z
git add .
git commit -a -m "binary commit"
git push heroku build:master --force
git checkout master
git branch -D build
rm -rf bin/
