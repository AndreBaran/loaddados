FROM golang:1.18
WORKDIR /go/src/github.com/postgres-go
COPY . .
RUN go get -u github.com/lib/pq
RUN go build -o main .
RUN go test ./...
CMD ["./main"]