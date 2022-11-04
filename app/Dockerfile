FROM golang:alpine

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www

ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64

RUN go build -o app

EXPOSE 3200
ENTRYPOINT ["/var/www/app"]