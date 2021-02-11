module github.com/zenoss/event-management-service

go 1.14

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/onsi/ginkgo v1.14.0
	github.com/onsi/gomega v1.10.1
	github.com/pkg/errors v0.9.1
	github.com/zenoss/event-context-svc v0.0.0-20210208232835-5442223405f6
	github.com/zenoss/zenkit/v5 v5.0.9
	github.com/zenoss/zing-proto/v11 v11.9.3
	google.golang.org/grpc v1.31.0
)

replace github.com/zenoss/zing-proto/v11 => ../zing-proto
