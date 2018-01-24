# https://devcenter.heroku.com/articles/container-registry-and-runtime
FROM gobuffalo/buffalo:development as builder
ENV BP=$GOPATH/src/github.com/gobuffalo/gobuffalo

RUN mkdir -p $BP
WORKDIR $BP

ADD package.json .
ADD yarn.lock .
RUN yarn install

ADD . .
RUN dep ensure

RUN buffalo build --static -o /bin/app --environment=production -d

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /bin/

ENV GO_ENV=production
COPY --from=builder /bin/app .
ENV ADDR=0.0.0.0
EXPOSE 3000

CMD exec /bin/app
