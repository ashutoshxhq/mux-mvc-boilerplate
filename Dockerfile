FROM golang:1.13-alpine
RUN apk update && apk upgrade
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["./main"]
EXPOSE 8080