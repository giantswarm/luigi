[![CircleCI](https://circleci.com/gh/giantswarm/luigi.svg?style=shield)](https://circleci.com/gh/giantswarm/luigi)
# luigi

## Installation

```
go get github.com/giantswarm/luigi
```

## Usage

```
kubectl log -f POD_NAME | luigi
```

It's under development, but it will turn your logs from this:

```
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/current.go:26","event":"update","function":"GetCurrentState","level":"debug","message":"looking for the secret in the Kubernetes API","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:48.880691+00:00"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/current.go:32","event":"update","function":"GetCurrentState","level":"debug","message":"did not find the secret in the Kubernetes API","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:48.882435+00:00"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/desired.go:20","event":"update","function":"GetDesiredState","level":"debug","message":"computing the desired secret","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:48.882509+00:00"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/desired.go:47","event":"update","function":"GetDesiredState","level":"debug","message":"computed the desired secret","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:48.882566+00:00"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:48","event":"update","function":"NewUpdatePatch","level":"debug","message":"finding out if the secret has to be created","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:48.882597+00:00"}
{"caller":"github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:122","event":"update","function":"NewUpdatePatch","level":"warning","message":"retrying due to error","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","stack":"[{/go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:115: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/update.go:41: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:54: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/resource.go:116: } {/go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/vaultcrt/create.go:21: } {Error making API request.\n\nURL: PUT http://vault.default.svc.cluster.local:8200/v1/pki-ci-wip-7fa7f/issue/role-org-aad8f23c2fa06d46c52ae4e06c8d4bdac6074b8f\nCode: 400. Errors:\n\n* cannot satisfy request, as TTL would result in notAfter 2018-12-01T15:34:48.886539753Z that is beyond the expiration of the CA certificate at 2018-12-01T15:34:35Z}]","time":"2018-10-02T15:34:48.886924+00:00","underlyingResource":"vaultcrtv2"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:48","event":"update","function":"NewUpdatePatch","level":"debug","message":"finding out if the secret has to be created","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:49.964022+00:00"}
{"caller":"github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:122","event":"update","function":"NewUpdatePatch","level":"warning","message":"retrying due to error","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","stack":"[{/go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:115: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/update.go:41: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:54: } {/go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/resource.go:116: } {/go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/vaultcrt/create.go:21: } {Error making API request.\n\nURL: PUT http://vault.default.svc.cluster.local:8200/v1/pki-ci-wip-7fa7f/issue/role-org-aad8f23c2fa06d46c52ae4e06c8d4bdac6074b8f\nCode: 400. Errors:\n\n* cannot satisfy request, as TTL would result in notAfter 2018-12-01T15:34:49.972891568Z that is beyond the expiration of the CA certificate at 2018-12-01T15:34:35Z}]","time":"2018-10-02T15:34:49.973578+00:00","underlyingResource":"vaultcrtv2"}
{"caller":"github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:48","event":"update","function":"NewUpdatePatch","level":"debug","message":"finding out if the secret has to be created","object":"/apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus","resource":"vaultcrtv2","time":"2018-10-02T15:34:50.973841+00:00"}
```

To this:

```
D 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.GetCurrentState looking for the secret in the Kubernetes API | cert-operator/service/controller/v2/resources/vaultcrt/current.go:26 | event=update
D 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.GetCurrentState did not find the secret in the Kubernetes API | cert-operator/service/controller/v2/resources/vaultcrt/current.go:32 | event=update
D 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.GetDesiredState computing the desired secret | cert-operator/service/controller/v2/resources/vaultcrt/desired.go:20 | event=update
D 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.GetDesiredState computed the desired secret | cert-operator/service/controller/v2/resources/vaultcrt/desired.go:47 | event=update
D 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.NewUpdatePatch finding out if the secret has to be created | cert-operator/service/controller/v2/resources/vaultcrt/create.go:48 | event=update
W 10/02 15:34:48 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.NewUpdatePatch retrying due to error | operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:122 | event=update | underlyingResource=vaultcrtv2
        /go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:115:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/update.go:41:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:54:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/resource.go:116:
        /go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/vaultcrt/create.go:21:
        Error making API request.

URL: PUT http://vault.default.svc.cluster.local:8200/v1/pki-ci-wip-7fa7f/issue/role-org-aad8f23c2fa06d46c52ae4e06c8d4bdac6074b8f
Code: 400. Errors:

* cannot satisfy request, as TTL would result in notAfter 2018-12-01T15:34:48.886539753Z that is beyond the expiration of the CA certificate at 2018-12-01T15:34:35Z
D 10/02 15:34:49 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.NewUpdatePatch finding out if the secret has to be created | cert-operator/service/controller/v2/resources/vaultcrt/create.go:48 | event=update
W 10/02 15:34:49 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.NewUpdatePatch retrying due to error | operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:122 | event=update | underlyingResource=vaultcrtv2
        /go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/operatorkit/controller/resource/retryresource/crud_resource_ops_wrapper.go:115:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/update.go:41:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/create.go:54:
        /go/src/github.com/giantswarm/cert-operator/service/controller/v2/resources/vaultcrt/resource.go:116:
        /go/src/github.com/giantswarm/cert-operator/vendor/github.com/giantswarm/vaultcrt/create.go:21:
        Error making API request.

URL: PUT http://vault.default.svc.cluster.local:8200/v1/pki-ci-wip-7fa7f/issue/role-org-aad8f23c2fa06d46c52ae4e06c8d4bdac6074b8f
Code: 400. Errors:

* cannot satisfy request, as TTL would result in notAfter 2018-12-01T15:34:49.972891568Z that is beyond the expiration of the CA certificate at 2018-12-01T15:34:35Z
D 10/02 15:34:50 /apis/core.giantswarm.io/v1alpha1/namespaces/default/certconfigs/ci-wip-7fa7f-prometheus vaultcrtv2.NewUpdatePatch finding out if the secret has to be created | cert-operator/service/controller/v2/resources/vaultcrt/create.go:48 | event=update
```
