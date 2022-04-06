FROM klakegg/hugo:ext-alpine-ci as builder

ADD package-lock.json .
ADD package.json .
RUN npm install .

ADD . .
ENV NODE_ENV production
RUN hugo -b "https://gobuffalo.io"

FROM nginx:alpine

COPY --from=builder /src/public /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

CMD sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'
