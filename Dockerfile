FROM quay.io/vektorcloud/go:1.13

RUN apk add --no-cache make

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
RUN make build && \
    mkdir -p /go/bin && \
    mv -v xiny /go/bin/

FROM scratch
ENV TERM=linux
COPY --from=0 /go/bin/xiny /xiny
ENTRYPOINT ["/xiny"]
CMD ["-i"]
