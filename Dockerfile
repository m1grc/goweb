FROM alpine:3.16
COPY serveur-web-golang /
RUN apk update && apk add libc6-compat
CMD ["/serveur-web-golang"]