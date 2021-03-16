# Event Management Service (event-management-service)
Manages user interactions for cloud events

## Purpose
The service provides an API for users to set event status, acknowledge events and annotate them within Zenoss cloud.

## Metrics
TODO: Describe the metrics exposed by this service, and what they indicate.

## API
For a given tenant the API allows users to
- set status for event(s)
- annotate event(s)

## Configuration
* `EVENT_MANAGEMENT_SERVICE_LOG_LEVEL`: Log level. Defaults to "info".
* `EVENT_MANAGEMENT_SERVICE_LOG_STACKDRIVER`: Whether to format logs for Stackdriver. Defaults to true.
* `EVENT_MANAGEMENT_SERVICE_AUTH_DISABLED`: Whether authentication is enforced. If true, middleware is used that injects an admin identity into unauthenticated requests.
* `EVENT_MANAGEMENT_SERVICE_AUTH_DEV_TENANT`: When auth is disabled, the tenant name to use as the identity. Defaults to "ACME".
* `EVENT_MANAGEMENT_SERVICE_AUTH_DEV_USER`: When auth is disabled, the user id to use as the identity. Defaults to "zcuser@acme.example.com".
* `EVENT_MANAGEMENT_SERVICE_TRACING_ENABLED`: Whether request tracing is enabled.
* `EVENT_MANAGEMENT_SERVICE_GRPC_LISTEN_ADDR`: The address on which the gRPC server should listen. Defaults to ":8080".
