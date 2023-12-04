FROM golang:alpine

WORKDIR /src/blogpost

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]