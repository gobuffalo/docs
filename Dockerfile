FROM klakegg/hugo:0.91.2-ext-alpine

COPY package.json package-lock.json ./

RUN npm ci
