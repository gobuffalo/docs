#!/bin/bash

echo $HEROKU_API_KEY

set -e # Abort script at first error
set -u # Disallow unset variables

# Only run when not part of a pull request and on the master branch
if [ $TRAVIS_PULL_REQUEST != "false" -o $TRAVIS_BRANCH != "master" ]
then
    echo "Skipping deployment on branch=$TRAVIS_BRANCH, PR=$TRAVIS_PULL_REQUEST"
    exit 0;
fi

# Install the toolbelt, and the required plugin.
sudo curl https://cli-assets.heroku.com/install-standalone.sh | sh
heroku plugins:install @heroku-cli/plugin-container-registry

# Build and release the application.
# To give access to your Heroku apps, you
# need to set the HEROKU_API_KEY environment variable.
echo "$HEROKU_API_KEY" | docker login --username=_ --password-stdin registry.heroku.com
heroku container:push web --app getbuffalo
