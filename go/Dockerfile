FROM golang:1.16-alpine as dev

WORKDIR /app

FROM golang:1.16-alpine as build

WORKDIR /server

COPY . ./

RUN go build -o server

FROM alpine as runtime
COPY --from=build /server /

EXPOSE 80 

CMD ./server
