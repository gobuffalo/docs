# https://devcenter.heroku.com/articles/container-registry-and-runtime
FROM gobuffalo/buffalo:development
ENV BP=$GOPATH/src/github.com/gobuffalo/gobuffalo

RUN mkdir -p $BP
WORKDIR $BP
ADD . .
RUN npm install

RUN buffalo build -o bin/gobuffalo

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=0 /go/src/github.com/gobuffalo/gobuffalo/bin/gobuffalo .

CMD ./gobuffalo
