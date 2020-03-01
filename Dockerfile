FROM golang:1.12-alpine as builder
WORKDIR /go/src/cinema-admin
COPY . .
RUN apk add --update git make
RUN go get -u github.com/Masterminds/glide
RUN glide install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /
COPY --from=builder /go/src/cinema-admin/cinema-admin ./cinema-admin
COPY --from=builder /go/src/cinema-admin/keys ./keys
COPY --from=builder /go/src/cinema-admin/template ./template
COPY --from=builder /go/src/cinema-admin/admin ./admin
COPY --from=builder /go/src/cinema-admin/app/views ./app/views
COPY --from=builder /go/src/cinema-admin/vendor/github.com/qor ./vendor/github.com/qor
CMD ["/cinema-admin"]
EXPOSE 4000 