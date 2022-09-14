# Kubexit Webhook

Kubernetes's job cannot exit successfully when there is
 an sidecar. Therefore, a number of solutions have been
 introduced in the community. One of the well-known
 terminator is called `kubexit`.

This project aims to develop the webhook server to mutate
 the pod specification to inject the `kubexit` and all
 the rest of required environment variable.
