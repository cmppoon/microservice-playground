FROM golang:latest

RUN mkdir -p /go/src/productsApi

WORKDIR /go/src/productsApi

COPY . /go/src/productsApi

RUN go install productsApi

CMD /go/bin/productsApi

EXPOSE 9090