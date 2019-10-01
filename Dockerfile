
# Prepare binary
FROM golang:latest AS builder

# Install dep
RUN go get -u github.com/golang/dep/...

WORKDIR /opt/workspace

RUN mkdir -p tictactoe/bin && mkdir -p tictactoe/pkg && mkdir -p tictactoe/src

COPY . tictactoe/

ENV GOPATH /opt/workspace/tictactoe
ENV GOBIN /opt/workspace/tictactoe/bin

WORKDIR /opt/workspace/tictactoe/src/tictactoe

RUN dep ensure -v

RUN CGO_ENABLED=0 go install main.go

#########
# To obtain a small image, copy just binary into alpine image and use it
FROM alpine

COPY --from=builder /opt/workspace/tictactoe/bin/main /main
RUN chmod 755 /main

EXPOSE 8080

ENTRYPOINT /main --port 8080 --host 0.0.0.0 
