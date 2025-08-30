FROM node:18-alpine AS builder
WORKDIR /app
COPY ./frontend/package*.json ./
RUN npm install
COPY ./frontend .
COPY ../.env .env
RUN npm run build

FROM nginx:stable-alpine
COPY ./frontend/nginx.conf /etc/nginx/templates/nginx.conf.template
COPY --from=builder /app/build /usr/share/nginx/html
COPY ./frontend/nginx.conf /etc/nginx/templates/default.conf.template
EXPOSE 80:80
CMD ["nginx", "-g", "daemon off;"]