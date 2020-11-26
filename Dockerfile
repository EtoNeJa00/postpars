FROM alpine
WORKDIR /go/bin
COPY dictionaries.json .
COPY postpars .
CMD ["/go/bin/postpars"]
