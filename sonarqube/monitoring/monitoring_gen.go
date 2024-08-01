package monitoring

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

/*
MetricsRequest Return monitoring metrics in Prometheus format.
Support content type 'text/plain' (default) and 'application/openmetrics-text'.
This endpoint can be accessed using a Bearer token, which needs to be defined in sonar.properties with the 'sonar.web.systemPasscode' key.
*/
type MetricsRequest struct{}

// MetricsResponse is the response for MetricsRequest
type MetricsResponse string
