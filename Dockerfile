FROM alpine:3.20
COPY ansdoc /usr/bin/ansdoc
ENTRYPOINT ["/usr/bin/ansdoc"]