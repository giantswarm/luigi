[![CircleCI](https://circleci.com/gh/giantswarm/luigi.svg?style=shield)](https://circleci.com/gh/giantswarm/luigi)
# luigi

## Installation

```
GO111MODULE="on" go get github.com/giantswarm/luigi
```

## Usage

```
kubectl log -f POD_NAME | luigi
```

It turns this:

```
$ opsctl deploy -i ginger aws-operator@k8scloudconfig-0.6.3
Requesting the creation of deployment using "App" method

    chart           aws-operator
    repository      aws-operator
    ref             k8scloudconfig-0.6.3
    version         8.7.0-903edaa40367d1b0a81fbe324391b083c5cc6f97
    installation    ginger

Requested the creation of deployment using "App" method

Waiting for deployment
Failed.
{"caller":"github.com/giantswarm/micrologger@v0.3.1/activation_logger.go:93","level":"error","stack":{"annotation":"execution failed error: deployment failed (status: `FAILED`, reason: `Release \"aws-operator-k8scloudconfig-0.6.3\" failed: admission webhook \"validation.gatekeeper.sh\" denied the request: [denied by unique-operator-version] label app.kubernetes.io/version=8.7.1-dev is not unique in resources of kind \"Deployment\"`)","kind":"unknown","stack":[{"file":"/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/cmd/deploy/appdeploy/deployer.go","line":338},{"file":"/Users/kopiczko/go/pkg/mod/github.com/giantswarm/backoff@v0.2.0/retry.go","line":13},{"file":"/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/cmd/deploy/appdeploy/deployer.go","line":347},{"file":"/Users/kopiczko/go/src/github.com/giantswarm/opsctl/command/deploy/command.go","line":346},{"file":"/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/output/funcs.go","line":33},{"file":"/Users/kopiczko/go/src/github.com/giantswarm/opsctl/command/deploy/command.go","line":356}]},"time":"2020-07-01T12:21:29.465397+00:00","verbosity":0}
```

Into this:

```
$ opsctl deploy -i ginger aws-operator@k8scloudconfig-0.6.3 | luigi
Requesting the creation of deployment using "App" method

    chart           aws-operator
    repository      aws-operator
    ref             k8scloudconfig-0.6.3
    version         8.7.0-903edaa40367d1b0a81fbe324391b083c5cc6f97
    installation    ginger

Requested the creation of deployment using "App" method

Waiting for deployment
Failed.
E 07/01 12:21:29 | micrologger@v0.3.1/activation_logger.go:93 | verbosity=0
	unknown: execution failed error: deployment failed (status: `FAILED`, reason: `Release "aws-operator-k8scloudconfig-0.6.3" failed: admission webhook "validation.gatekeeper.sh" denied the request: [denied by unique-operator-version] label app.kubernetes.io/version=8.7.1-dev is not unique in resources of kind "Deployment"`)
	/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/cmd/deploy/appdeploy/deployer.go:338
	/Users/kopiczko/go/pkg/mod/github.com/giantswarm/backoff@v0.2.0/retry.go:13
	/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/cmd/deploy/appdeploy/deployer.go:347
	/Users/kopiczko/go/src/github.com/giantswarm/opsctl/command/deploy/command.go:346
	/Users/kopiczko/go/src/github.com/giantswarm/opsctl/pkg/output/funcs.go:33
	/Users/kopiczko/go/src/github.com/giantswarm/opsctl/command/deploy/command.go:356
```
