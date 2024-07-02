FROM nginx:latest

COPY nginx.conf /etc/nginx/conf.d/default.conf

WORKDIR /usr/share/nginx/homepage

COPY src src
COPY index.html index.html 

CMD ["nginx", "-g", "daemon off;"]

EXPOSE 80