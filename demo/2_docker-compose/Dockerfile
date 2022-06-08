FROM golang:alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/appserver /app/main.go

FROM alpine
WORKDIR /usr/bin
EXPOSE 8000
ENV PORT=8000
COPY --from=builder /app/appserver .
CMD ["appserver"]