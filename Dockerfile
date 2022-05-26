FROM alpine:3.12
WORKDIR /workspace
RUN apk update && apk add git-daemon && rm -rf /var/cache/apk
ADD sgits .
CMD ["./sgits"]
