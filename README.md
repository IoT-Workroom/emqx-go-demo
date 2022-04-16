# EMQX Go demo

This is an example of a client that publishes some data for the EMQX project.

The repository is a part of the
article: [How to run own self-hosted MQTT broker?](https://medium.com/@yehors/how-to-run-own-self-hosted-mqtt-broker-b59a15b2487c)

## Running

You must run an EMQX instance via Docker Compose before running below command.

```bash
go run send_events.go
```
