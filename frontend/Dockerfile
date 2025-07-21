FROM node:18-alpine AS builder
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend .
COPY ../.env .env
RUN npm run build

FROM nginx:stable-alpine
COPY ./frontend/nginx.conf.template /etc/nginx/templates/nginx.conf.template
COPY --from=builder /app/build /usr/share/nginx/html
COPY --from=builder /app/.env .env 
EXPOSE 80
CMD /bin/sh -c "envsubst '\$BACKEND_CONTAINER_NAME \$BACKEND_API' < /etc/nginx/templates/nginx.conf.template > /etc/nginx/conf.d/default.conf && nginx -g 'daemon off;'"
