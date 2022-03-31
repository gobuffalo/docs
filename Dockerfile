FROM klakegg/hugo:0.95.0-ext-ubuntu-onbuild as builder

WORKDIR /site
ADD . .

RUN hugo 

FROM nginx:alpine
COPY --from=builder /site/public /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

CMD sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'

