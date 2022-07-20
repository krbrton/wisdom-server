FROM golang:1.18.3

RUN mkdir /wisdom-client
COPY . /wisdom-client
WORKDIR /wisdom-client

RUN mkdir ~/.wisdom-server
RUN cp docker/client-config.yaml ~/.wisdom-server/config.yaml
RUN cp docker/quote.txt ~/.wisdom-server/quote.txt
RUN rm -rf docker

RUN go build ./cmd/client/main.go
RUN mv main /usr/bin/wisdom-client

ENTRYPOINT /usr/bin/wisdom-client
