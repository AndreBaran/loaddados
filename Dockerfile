FROM golang:1.18
WORKDIR /go/src/neoway
COPY . .
RUN mv .env.example .env

# go requirements
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# building
RUN go build -o main .

# go test
RUN go test ./...

#go run
CMD ["./main"]