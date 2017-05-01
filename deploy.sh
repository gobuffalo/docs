#!/bin/bash

set -ex

heroku container:push web
# heroku run bin/gobuffalo migrate
