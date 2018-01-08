#!/bin/bash

echo $HEROKU_API_KEY

set -e # Abort script at first error
set -u # Disallow unset variables

# Only run when not part of a pull request and on the master branch
if [ $TRAVIS_PULL_REQUEST != "false" -o $TRAVIS_BRANCH != "auto-deploy" ]
then
    echo "Skipping deployment on branch=$TRAVIS_BRANCH, PR=$TRAVIS_PULL_REQUEST"
    exit 0;
fi

# Install the toolbelt, and the required plugin.
npm install -g heroku-cli
heroku plugins:install heroku-container-registry --force

# Build and release the application.
# To give access to your Heroku apps, you
# need to set the HEROKU_API_KEY environment variable.
docker login --username=_ --password=$HEROKU_API_KEY registry.heroku.com
heroku container:push web --app getbuffalo
