from golang:1.15.0-alpine3.12
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["GoConway", "server"]
EXPOSE 4000