sendgrid-event-exporter
===
A api server which redirect http request to slack message.

# Run

```
export SLACK_TOKEN=[slackAppToken or slackPersonalToken]
export SLACK_CHANNEL=[slackChannelID]
make run
```

# Test
```
make test
```

# Run with Docker

```
export SLACK_TOKEN=[slackAppToken or slackPersonalToken]
export SLACK_CHANNEL=[slackChannelID]
make docker
```

# Use cases

Integration service which has no-slack http webhook
- Sendgrid
