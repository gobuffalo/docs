<%= code("text") { %>
$ buffalo new coke
Buffalo version <%= version %>

--> go get -u golang.org/x/tools/cmd/goimports
--> go install golang.org/x/tools/cmd/goimports
--> go get -u github.com/golang/dep
--> go install github.com/golang/dep
--> go get -u github.com/motemen/gore
--> go install github.com/motemen/gore
--> $GOPATH/src/github.com/markbates/coke/README.md
--> $GOPATH/src/github.com/markbates/coke/actions/actions_test.go
--> $GOPATH/src/github.com/markbates/coke/actions/app.go
--> $GOPATH/src/github.com/markbates/coke/actions/home.go
--> $GOPATH/src/github.com/markbates/coke/actions/home_test.go
--> $GOPATH/src/github.com/markbates/coke/actions/render.go
--> $GOPATH/src/github.com/markbates/coke/.codeclimate.yml
--> $GOPATH/src/github.com/markbates/coke/.gitignore
--> $GOPATH/src/github.com/markbates/coke/grifts/routes.go
--> $GOPATH/src/github.com/markbates/coke/locales/all.en-us.yaml
--> $GOPATH/src/github.com/markbates/coke/main.go
--> $GOPATH/src/github.com/markbates/coke/templates/_flash.html
--> $GOPATH/src/github.com/markbates/coke/templates/application.html
--> $GOPATH/src/github.com/markbates/coke/templates/index.html
--> $GOPATH/src/github.com/markbates/coke/.buffalo.dev.yml
--> $GOPATH/src/github.com/markbates/coke/assets/images/logo.svg
--> $GOPATH/src/github.com/markbates/coke/assets/css/application.scss
--> $GOPATH/src/github.com/markbates/coke/assets/js/application.js
--> $GOPATH/src/github.com/markbates/coke/public/assets/.gitignore
--> $GOPATH/src/github.com/markbates/coke/webpack.config.js
--> npm init -y
--> $GOPATH/src/github.com/markbates/coke/models/models.go
--> $GOPATH/src/github.com/markbates/coke/models/models_test.go
--> $GOPATH/src/github.com/markbates/coke/grifts/seed.go
--> go get github.com/markbates/pop/...
--> go install github.com/markbates/pop/soda
--> database.yml
--> $GOPATH/src/github.com/markbates/coke/Dockerfile
--> $GOPATH/src/github.com/markbates/coke/.dockerignore
--> dep init
--> goimports -w coke/actions/actions_test.go coke/actions/app.go coke/actions/home.go coke/actions/home_test.go coke/actions/render.go coke/grifts/routes.go coke/grifts/seed.go coke/main.go coke/models/models.go coke/models/models_test.go
Congratulations! Your application, coke, has been successfully built!

You can find your new application at:
$GOPATH/src/github.com/markbates/coke

Please read the README.md file in your new application for next steps on running your application.
<% } %>
