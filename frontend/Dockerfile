FROM golang:1.9

WORKDIR /go/src/github.com/kubernetes-playgrounds/kubernetes-envoy/frontend
COPY . .
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init
RUN make build-linux
EXPOSE 8080
CMD ["/go/src/github.com/kubernetes-playgrounds/kubernetes-envoy/frontend/envoy_frontend_unix"]