# Event Management Service (event-management-service)

Manages user interactions for cloud events

## Purpose

The service provides an API for users to set event status, acknowledge events and annotate them within Zenoss cloud.

## Metrics

TODO: Describe the metrics exposed by this service, and what they indicate.

## Configuration

* `EVENT_MANAGEMENT_SERVICE_LOG_LEVEL`: Log level. Defaults to "info".
* `EVENT_MANAGEMENT_SERVICE_LOG_STACKDRIVER`: Whether to format logs for Stackdriver. Defaults to true.
* `EVENT_MANAGEMENT_SERVICE_AUTH_DISABLED`: Whether authentication is enforced. If true, middleware is used that injects an admin identity into unauthenticated requests.
* `EVENT_MANAGEMENT_SERVICE_AUTH_DEV_TENANT`: When auth is disabled, the tenant name to use as the identity. Defaults to "ACME".
* `EVENT_MANAGEMENT_SERVICE_AUTH_DEV_USER`: When auth is disabled, the user id to use as the identity. Defaults to "zcuser@acme.example.com".
* `EVENT_MANAGEMENT_SERVICE_TRACING_ENABLED`: Whether request tracing is enabled.
* `EVENT_MANAGEMENT_SERVICE_GRPC_LISTEN_ADDR`: The address on which the gRPC server should listen. Defaults to ":8080".

## API

For a given tenant the API allows users to

* set status for event(s) - acknowledge and status
* annotate event(s)

<!-- markdown-swagger -->
 Endpoint                        | Method | Auth? | Description
 ------------------------------- | ------ | ----- | -----------
 `/v1/event-management/status`   | POST   | Yes   | Acknowledge and/or set status
 `/v1/event-management/annotate` | POST   | Yes   | Add/edit an annotation for the event
<!-- /markdown-swagger -->

Status request fields:

* statuses - array of EMEventStatus  messages
* EMEventStatus  contians
  * eventId - the event id
  * occurrenceId - the occurrence id
  * acknowledge - boolean value for event acknowledgement. true/false
  * statusWrapper - a wrapper for EMStatus  value
* EMStatus has the following possible values
  * EM_STATUS_DEFAULT
  * EM_STATUS_OPEN
  * EM_STATUS_SUPPRESSED
  * EM_STATUS_CLOSED

Status response fields:

* statusResponses - array of EMEventStatusResponse messages
* EMEventStatusResponse  contans
  * eventid
  * occurrenceId
  * success - true or false
  * error - details when success is false

  Annotate request fields:
* annotations - array of Annotation messages
* Annotation   contians
  * eventId - the event id
  * occurrenceId - the occurrence id
  * annotationId -  the id of the annotaion when editing one. leave out for new notes
  * annotation  - the text to be used.

Annotation response fields:

* annotationResponses  - array of AnnotationResponse messages
* AnnotationResponse contains
  * eventid
  * occurrenceId
  * annotationId - the id of the new annotation or that of the one edited.
  * success - true or false
  * error - details when success is false

## How to use the HTTP API

  1. Acquire a Zenoss Cloud API key for authenticating requests.
  2. Make requests to the event-management-service via the API endpoint of the tenant for which you made an API key.

### List of event-management-service API endpoints by GCP project

* staging: <https://api-zing-preview.zenoss.io/v1/event-management>
* production: <https://api.zenoss.io/v1/event-management>
* production: <https://api.zenoss.io/v1/event-management> (CNAME to api-zcloud-prod.zenoss.io)
* production: <https://api-zcloud-prod.zenoss.io/v1/event-management>
* production 2: <https://api2.zenoss.io/v1/event-management> (CNAME to api-zcloud-prod2.zenoss.io)
* production emea: <https://api.zenoss.eu/v1/event-management> (CNAME to api-zcloud-emea.zenoss.eu)
* production emea: <https://api-zcloud-emea.zenoss.eu/v1/event-management>
* testing: <https://api-zing-testing-200615.zenoss.io/v1/event-management>
* zing-perf: <https://api-zing-perf.zenoss.io/v1/event-management>
* zing-dev: <https://api.zing.soy/v1/event-management> (CNAME to api-zing-dev-197522.zing.soy)
* zing-dev: <https://api-zing-dev-197522.zing.soy/v1/event-management>

Any HTTP request to the event-management-service should include the header `zenoss-api-key` for authentication.
Note: The examples below use an event id under the dev tenant and used the api key for the dev tenant.
Data can be verified in the event console under event details or in firestore under EventContextTenants> <your tenant> > Events > <event-id> > Occurrences....

Example 1:
command line:

```sh
curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@ack.json  https://api.zing.soy/v1/event-management/status
```

Given a ack.json file that has

```json
{
    "statuses": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "acknowledge": false
    }]
}
```

response will be like

```json
{"statusResponses":[{"eventId":"AAAABWRDWOPER3lzOFI2tSCG49g=","occurrenceId":"AAAABWRDWOPER3lzOFI2tSCG49g=:1","success":true}]}
```

Example 2:
command line:

```sh
curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@stat.json  https://api.zing.soy/v1/event-management/status
```

Given a stat.json file that has

```json
{
    "statuses": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "statusWrapper": { "status": "EM_STATUS_SUPPRESSED"}
    }]
}
```

response will be like

```json
{"statusResponses":[{"eventId":"AAAABWRDWOPER3lzOFI2tSCG49g=","occurrenceId":"AAAABWRDWOPER3lzOFI2tSCG49g=:1","success":true}]}
```

Example 3:
command line:

```sh
curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@annot.json  https://api.zing.soy/v1/event-management/annotate
```

Given a annot.json file that has

```json
{
    "annotations": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "annotation":  "annotation from a curl api request"
    }]
}
```

response will be like:

```json
{"annotationResponses":[{"eventId":"AAAABWRDWOPER3lzOFI2tSCG49g=","occurrenceId":"AAAABWRDWOPER3lzOFI2tSCG49g=:1","annotationId":"Q0WQI5BQayOBBoeZ2LBD","success":true}]}
```
