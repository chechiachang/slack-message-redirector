sendgrid-event-redirector
===

A api server which redirect http request to slack message. Use cases:
- Integration service which has no-slack http webhook
  - Sendgrid

---

# Run

```
export SLACK_URL=[slackUrl]
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
export SLACK_URL=[slackUrl]
make docker
make docker-push
```

---

Run on Kubernetes
===

# Deploy

Set credentials
```
export AUTH_TOKEN=
export SLACK_URL=
export SLACK_CHANNEL=
export SLACK_USERNAME=
export SLACK_ICON_EMOJI=

kubectl create secret generic slack-message-redirector-credentials \
--from-literal=AUTH_TOKEN=${AUTH_TOKEN} \
--from-literal=SLACK_URL=${SLACK_URL} \
--from-literal=SLACK_CHANNEL=${SLACK_CHANNEL} \
--from-literal=SLACK_USERNAME=${SLACK_USERNAME} \
--from-literal=SLACK_ICON_EMOJI=${SLACK_ICON_EMOJI}
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
