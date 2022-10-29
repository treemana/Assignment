FROM golang:1.19

COPY . /root/assignment

WORKDIR /root/assignment

RUN go env -w GOPROXY=https://goproxy.io,direct && go build -v -o /root/fetch

WORKDIR /root

RUN rm -rf /root/assignment/
