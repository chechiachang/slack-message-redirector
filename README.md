sendgrid-event-redirector
===

A api server which redirect http request to slack message. Use cases:
- Integration service which has no-slack http webhook
  - Sendgrid

---

# Run

```
export SLACK_TOKEN=[slackAppToken or slackPersonalToken]
export SLACK_CHANNEL=[slackChannelID]
make run
```

# Test
```
make test
make run
make message
```

# Run with Docker

```
export SLACK_TOKEN=[slackAppToken or slackPersonalToken]
export SLACK_CHANNEL=[slackChannelID]
make docker
make docker-push
```

---

Run on Kubernetes
===

# Deploy

Set credentials
```
SLACK_TOKEN=
SLACK_CHANNEL_ID=

kubectl create secret generic slack-message-redirector-credentials \
--from-literal=SLACK_TOKEN=${SLACK_TOKEN} \
--from-literal=SLACK_CHANNEL_ID=${SLACK_CHANNEL_ID}
```

```
kubectl apply -f kubernetes/
```

Wait for loadbalancer provisoning

Access
```
APP_IP=$(kubectl get svc --selector='app=slack-message-redirector' -o jsonpath='{.items[0].status.loadBalancer.ingress[0].ip}')
make remote
```
