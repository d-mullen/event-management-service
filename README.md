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
- set status for event(s) - acknowledge and status
- annotate event(s)

<!-- markdown-swagger -->
 Endpoint                        | Method | Auth? | Description
 ------------------------------- | ------ | ----- | -----------
 `/v1/event-management/status`   | POST   | Yes   | Acknowledge and/or set status         
 `/v1/event-management/annotate` | POST   | Yes   | Add/edit an annotation for the event
<!-- /markdown-swagger -->

## How to use the API

  1. Create a Zenoss API Key in one of the whitelisted tenants for a stack. Currently: (zing-dev: cladhoc; zing-testing; e2e-long; staging: qa-staging; prod: guidedtour)
  2. Make requests to the event-management-service via the API endpoint of the tenant for which you made an API key.

### List of user-trials API endpoints by GCP project

- staging: <https://ingress-zcloud-qa.zenoss.io/v1/event-management>
- production: <https://api.zenoss.io/v1/event-management>

Any HTTP request to the  event-management-service should include the header zenoss-api-key for authentication.
Note: The examples below use an event id under the dev tenant and used the api key for the dev tenant.
Data can be verified in the event console under event details or in firestore under EventContextTenants> <your tenant> > Events > <event-id> > Occurrences....

Example 1:
command line:

$ curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@ack.json  https://api.zing.soy/v1/event-management/status

Given a ack.json file that has 

{
"statuses": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "acknowledge": false
    }]
}

Example 2:
command line:

$ curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@stat.json  https://api.zing.soy/v1/event-management/status

Given a stat.json file that has 

{
"statuses": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "statusWrapper": { "status": "EM_STATUS_SUPPRESSED"}
    }]
}

Example 3:
command line:

$ curl  -H "zenoss-api-key:<your-key>" -k -S -X POST -d@annot.json  https://api.zing.soy/v1/event-management/annotate

Given a annot.json file that has 

{
"annotations": [{
        "eventId": "AAAABWRDWOPER3lzOFI2tSCG49g=",
        "occurrenceId": "AAAABWRDWOPER3lzOFI2tSCG49g=:1",
        "annotation":  "annotation from a curl api request"
    }]
}

