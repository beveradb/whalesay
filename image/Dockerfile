FROM golang:alpine as build

WORKDIR /whalesay
COPY ../image .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o whalesay .

FROM alpine:edge
LABEL maintainer "Niko Virtala <niko@nikovirtala.io>"

RUN adduser -D -g '' -h /whalesay whalesay
USER whalesay
WORKDIR /whalesay

COPY --from=build /whalesay .
EXPOSE 8080
ENTRYPOINT ["./whalesay"]
