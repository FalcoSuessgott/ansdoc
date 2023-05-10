FROM alpine:3.18
COPY ansdoc /usr/bin/ansdoc
ENTRYPOINT ["/usr/bin/ansdoc"]