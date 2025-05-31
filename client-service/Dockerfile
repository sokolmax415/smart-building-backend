FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY config/config.yaml ./config/config.yaml
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/main ./cmd

EXPOSE 8080

CMD ["/bin/main"]
