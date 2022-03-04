ARG ZENKIT_BUILD_VERSION=1.9.1

FROM zenoss/zenkit-build:${ZENKIT_BUILD_VERSION}

COPY . /go/src/github.com/zenoss/event-management-service
WORKDIR /go/src/github.com/zenoss/event-management-service
RUN go build -mod vendor -o /bin/event-management-service

FROM alpine
RUN apk add --no-cache curl

COPY --from=0 /bin/event-management-service /bin/event-management-service
RUN addgroup -S zing && adduser -S -G zing -u 512 512
USER 512

EXPOSE 8080

CMD ["/bin/event-management-service"]
