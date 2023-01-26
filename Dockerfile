FROM golang:1.19.0
WORKDIR /usr/src/app
RUN go install github.com/cosmtrek/air@latest
COPY . .
WORKDIR /usr/src/app/domain
RUN go mod tidy
