FROM golang:1.11 as builder

WORKDIR /go/src/go-micro/vessel-service

COPY . .

RUN ls

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init -v

RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -o vessel-service -a -installsuffix cgo main.go repository.go handler.go datastore.go


FROM debian:latest

# RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/go-micro/vessel-service/vessel-service .

CMD ["./vessel-service"]
