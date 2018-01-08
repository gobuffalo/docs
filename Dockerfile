# https://devcenter.heroku.com/articles/container-registry-and-runtime
FROM gobuffalo/buffalo:development as builder
ENV BP=$GOPATH/src/github.com/gobuffalo/gobuffalo

RUN mkdir -p $BP
WORKDIR $BP

RUN dep ensure
ADD package.json .
ADD yarn.lock .
RUN yarn install

ADD . .

RUN buffalo build --static -o /bin/app

FROM alpine
RUN apk add --no-cache bash

WORKDIR /bin/

COPY --from=builder /bin/app .

EXPOSE 3000

CMD /bin/app
