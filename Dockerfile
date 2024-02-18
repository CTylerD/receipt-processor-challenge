FROM golang:latest

WORKDIR /go/receipt-processor-challenge

RUN go mod init receipt_manager

RUN go get -u github.com/google/uuid
RUN go get -u github.com/gorilla/mux

COPY ./src .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]