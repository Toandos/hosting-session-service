# Hosting Session Service
This repository belongs to the [hosting product](https://github.com/Toandos/hosting-gitops) and contains a go implementation of the [session service definition](https://github.com/Toandos/apis/blob/main/proto/toando/hosting/v1alpha1/session.proto) in the Toando APIs repository. It's based on [ConnectRPC](https://connectrpc.com) and the auto-generated [client library for go](https://github.com/Toandos/apis-go).

## Deployment
The recommended way of deployment is using the helm chart which is automatically published to ```oci://ghcr.io/toandos/charts/hosting-session-service```.