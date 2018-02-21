FROM quay.io/vektorcloud/go:1.9

RUN apk add --no-cache make

COPY Gopkg.* /go/src/github.com/bcicen/xiny/
WORKDIR /go/src/github.com/bcicen/xiny/
RUN dep ensure -vendor-only

COPY . /go/src/github.com/bcicen/xiny
RUN make build && \
    mkdir -p /go/bin && \
    mv -v xiny /go/bin/

FROM scratch
COPY --from=0 /go/bin/xiny /xiny

ENTRYPOINT ["/xiny"]
CMD ["-i"]
