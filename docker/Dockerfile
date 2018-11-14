FROM golang:1.11.2 as builder
RUN go get github.com/MrManz/AWS-Dev-Challenge
RUN go get github.com/labstack/echo/...
RUN go get github.com/lib/pq
WORKDIR /go/src/github.com/MrManz/AWS-Dev-Challenge
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/MrManz/AWS-Dev-Challenge .
CMD ["./app"]
EXPOSE 3000