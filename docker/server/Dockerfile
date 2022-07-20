FROM golang:1.18.3

RUN mkdir /wisdom-server
COPY . /wisdom-server
WORKDIR /wisdom-server

RUN mkdir ~/.wisdom-server
RUN cp docker/server-config.yaml ~/.wisdom-server/config.yaml
RUN cp docker/quote.txt ~/.wisdom-server/quote.txt
RUN rm -rf docker

RUN go build ./cmd/server/main.go
RUN mv main /usr/bin/wisdom-server

ENTRYPOINT /usr/bin/wisdom-server
