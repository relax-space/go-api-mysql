FROM golang AS builder

WORKDIR /go/src/ping-mysql
COPY . .
# disable cgo
ENV CGO_ENABLED=0
# build steps
RUN echo ">>> 1: go version" && go version
RUN echo ">>> 2: go get" && go get -v -d
RUN echo ">>> 3: go install" && go install

FROM pangpanglabs/alpine-ssl
WORKDIR /go/bin/
RUN apk add --no-cache bash
COPY --from=builder /go/bin/ping-mysql ./ping-mysql
COPY --from=builder /go/src/ping-mysql/wait-for-it.sh ./wait-for-it.sh
RUN chmod +x ./wait-for-it.sh
EXPOSE 8080
CMD ["./ping-mysql"]


