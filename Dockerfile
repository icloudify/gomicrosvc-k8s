FROM golang:1.15-alpine

MAINTAINER ravindra
ENV SOURCES /go/src/github.com/ravindra031/gomicrosvc-k8s/
COPY . ${SOURCES}

RUN cd ${SOURCES} && go install
ENV PORT 8080
EXPOSE 8080

ENTRYPOINT gomicrosvc-k8s