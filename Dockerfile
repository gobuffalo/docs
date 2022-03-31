FROM node:alpine as node

WORKDIR /deps

ADD package.json .
ADD package-lock.json .

RUN npm i

FROM klakegg/hugo:0.95.0-ext-ubuntu-onbuild as builder

WORKDIR /site
ADD . .
COPY --from=node /deps/node_modules .

RUN hugo -b "https://buffalodocs.herokuapp.com"

FROM nginx:alpine

COPY --from=builder /site/public /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

CMD sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'
