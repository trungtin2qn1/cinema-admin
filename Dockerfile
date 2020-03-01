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
CMD ["/cinema-admin"]
EXPOSE 4000 