# https://devcenter.heroku.com/articles/container-registry-and-runtime
FROM gobuffalo/buffalo:development
ENV BP=$GOPATH/src/github.com/gobuffalo/gobuffalo

RUN mkdir -p $BP
WORKDIR $BP
ADD . .
RUN npm install

RUN buffalo build -o bin/gobuffalo

CMD ./bin/gobuffalo
