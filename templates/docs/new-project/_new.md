```bash
$ buffalo new coke
Buffalo version <%= version %>

      create  .buffalo.dev.yml
      create  assets/images/logo.svg
      create  assets/css/application.scss
      create  assets/images/favicon.ico
      create  assets/js/application.js
      create  .babelrc
      create  package.json
      create  public/assets/.keep
      create  webpack.config.js
         run  yarn install --no-progress --save
yarn install v0.27.5
info No lockfile found.
[1/4] Resolving packages...
[2/4] Fetching packages...
[3/4] Linking dependencies...
[4/4] Building fresh packages...
success Saved lockfile.
Done in 11.71s.
      create  models/models.go
      create  models/models_test.go
      create  grifts/db.go
         run  go get github.com/gobuffalo/pop/...
      create  ./database.yml
         run  goimports -w coke/grifts/db.go coke/models/models.go coke/models/models_test.go
      create  Dockerfile
      create  .dockerignore
         run  go get -u golang.org/x/tools/cmd/goimports
      create  README.md
      create  actions/actions_test.go
      create  actions/app.go
      create  actions/home.go
      create  actions/home_test.go
      create  actions/render.go
      create  .codeclimate.yml
      create  .env
      create  grifts/init.go
      create  inflections.json
      create  locales/all.en-us.yaml
      create  main.go
      create  public/robots.txt
      create  templates/_flash.html
      create  templates/application.html
      create  templates/index.html
         run  go get -t ./...
         run  goimports -w actions/actions_test.go actions/app.go actions/home.go actions/home_test.go actions/render.go grifts/db.go grifts/init.go main.go models/models.go models/models_test.go
      create  .gitignore
         run  git init
Initialized empty Git repository in /Users/markbates/Dropbox/development/gocode/src/github.com/markbates/coke/.git/
         run  git add .
         run  git commit -q -m Initial Commit
INFO[0055] Congratulations! Your application, coke, has been successfully built!

 
INFO[0055] You can find your new application at:
/Users/markbates/Dropbox/development/gocode/src/github.com/markbates/coke 
INFO[0055] 
Please read the README.md file in your new application for next steps on running your application.
```
