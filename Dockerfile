FROM alpine
WORKDIR /go/bin
COPY dictionaries.json .
COPY postpars .
CMD ["export LISTEN_PORT=8080"]
CMD ["/go/bin/postpars"]
