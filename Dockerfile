FROM klakegg/hugo:0.95.0-ext-ubuntu-onbuild as builder

WORKDIR /site
ADD . .

RUN hugo -b "https://buffalodocs.herokuapp.com"

FROM nginx:alpine
COPY --from=builder /site/public /usr/share/nginx/html
COPY nginx.conf /etc/nginx/templates/nginx.conf.template

ENV NGINX_HOST buffalodocs.herokuapp.com
ENV NGINX_PORT $PORT

# CMD sed -i -e 's/$PORT/'"$PORT"'/g' /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'

