FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
# copy the source files
COPY . .

RUN go build -o ./out/urls ./cmd/app
ENTRYPOINT ["./out/urls"]
