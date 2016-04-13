FROM golang:1.5.1
#ADD docker-build-noproxy.sh /docker-build.sh
ADD https://github.com/Masterminds/glide/releases/download/0.10.1/glide-0.10.1-linux-amd64.tar.gz /glide.tgz
RUN mkdir /glide
RUN tar xvfz /glide.tgz -C /glide

ADD *.go /go/src/couchbase-cluster-go/
ADD glide.yaml /go/src/couchbase-cluster-go/
ADD glide.lock /go/src/couchbase-cluster-go/
ADD data       /go/src/couchbase-cluster-go/data/
ADD cmd        /go/src/couchbase-cluster-go/cmd/
#ADD start.sh /
#RUN chmod +x /start.sh

ENV GO15VENDOREXPERIMENT=1
WORKDIR /go/src/couchbase-cluster-go/
ENV HTTP_PROXY "proxy-us.intel.com:911"
ENV HTTPS_PROXY "proxy-us.intel.com:911"

RUN HTTP_PROXY=proxy-us.intel.com:911 HTTPS_PROXY=proxy-us.intel.com:911 /glide/linux-amd64/glide --debug --home /go/src/couchbase-cluster-go/ install

RUN echo "gopath: " $GOPATH
RUN echo "goroot: " $GOROOT
#RUN go install

#CMD ["/start.sh"]
