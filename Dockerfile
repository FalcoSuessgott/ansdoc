FROM alpine:3.17
COPY ansdoc /usr/bin/ansdoc
ENTRYPOINT ["/usr/bin/ansdoc"]