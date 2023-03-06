FROM golang:alpine
RUN apk update && apk add --no-cache git
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get -d -v
CMD ["bookstore"]