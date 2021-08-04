FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY *.html ./
COPY *.css ./
COPY *.yml ./

ENV DB_PORT 8000
RUN go build -o /docker-gs-ping
CMD [ "/docker-gs-ping" ]

