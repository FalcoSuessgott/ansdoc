FROM alpine:3.19
COPY ansdoc /usr/bin/ansdoc
ENTRYPOINT ["/usr/bin/ansdoc"]