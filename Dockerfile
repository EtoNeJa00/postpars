FROM alpine
WORKDIR /go/bin
COPY configuraton/configuration.json configuraton/
COPY Rest_Api_TV .
CMD ["/go/bin/postpars"]
