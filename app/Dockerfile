FROM golang:alpine

ENV GO111MODULE=on \
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64

RUN go build -o app

RUN mkdir -p /var/www

WORKDIR /var/www

COPY . /var/www


EXPOSE 3200
ENTRYPOINT ["/var/www/app"]