FROM golang:alpine as dev

WORKDIR /app

COPY . .

CMD [ "air" ]