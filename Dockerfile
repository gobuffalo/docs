# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:development as builder

RUN go version
RUN which go
RUN echo $PATH
ENV GOPROXY="https://proxy.golang.org"
ENV GO111MODULE="on"
ARG GITHUB_TOKEN=local
ENV GITHUB_TOKEN ${GITHUB_TOKEN}

# this will cache the npm install step, unless package.json changes
ADD package.json .
ADD package-lock.json .
RUN npm install --no-progress
ADD . .
RUN buffalo build --static -o /bin/app -v --skip-template-validation

# Comment out to run the binary in "production" mode:
# ENV GO_ENV=production

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 3000

# Comment out to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD exec /bin/app

