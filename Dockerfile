FROM golang:1.19-alpine

WORKDIR /app
COPY . ./

ENV GIN_MODE=release
RUN go mod download
RUN go build -o stella ./cmd

EXPOSE 4000
CMD [ "/app/stella" ]