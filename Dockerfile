FROM gobuffalo/buffalo:development

RUN mkdir -p $GOPATH/src/github.com/gobuffalo/gobuffalo
WORKDIR $GOPATH/src/github.com/gobuffalo/gobuffalo

# this will cache the npm install step, unless package.json changes
ADD package.json .
ADD yarn.lock .
RUN yarn install --no-progress
ADD . .
RUN dep ensure -v
RUN buffalo build --static -o /bin/app

EXPOSE 3000

# Comment out to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD exec /bin/app
