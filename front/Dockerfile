FROM node:20-alpine as builder

COPY package*.json /

RUN npm install --ignore-scripts

COPY . .

RUN npm run build

FROM nginx:1.25-alpine
EXPOSE 80

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /dist /var/www/front
