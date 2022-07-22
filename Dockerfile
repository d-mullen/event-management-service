ARG ZENKIT_BUILD_VERSION=1.12.0

FROM zenoss/zenkit-build:${ZENKIT_BUILD_VERSION} as stage0

COPY . /go/src/github.com/zenoss/event-management-service
WORKDIR /go/src/github.com/zenoss/event-management-service

FROM stage0 as dev
RUN go build -buildvcs=false -mod vendor -o /bin/event-management-service

FROM stage0 as builder
RUN go build -mod vendor -o /bin/event-management-service

FROM alpine as production
RUN apk add --no-cache curl

COPY --from=builder /bin/event-management-service /bin/event-management-service
RUN addgroup -S zing && adduser -S -G zing -u 512 512
USER 512

EXPOSE 8080

CMD ["/bin/event-management-service"]
