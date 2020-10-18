FROM alpine:3.9

RUN apk --no-cache update && \
	apk --no-cache add ca-certificates tzdata && \
	rm - rf /var/cache/apk/*

RUN adduser -D -g '' appuser

COPY ./cmd/gideon/gideon /app/gideon

EXPOSE 80

ENTRYPOINT ["/app/gideon"]