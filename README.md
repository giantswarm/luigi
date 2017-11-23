# luigi

## Instalation

```
go get github.com/giantswarm/luigi
```

## Usage

```
kubectl log -f POD_NAME | luigi
```

It's under development, but it will turn your logs from this:

```
{"caller":"github.com/giantswarm/azure-operator/service/service.go:67","debug":"creating azure-operator with config: service.Config{Logger:(*micrologger.logger)(0xc42022f280), Flag:(*flag.Flag)(0xc420175680), Viper:(*viper.Viper)(0xc4201af770), Description:\"The azure-operator manages Kubernetes clusters on Azure.\", GitCommit:\"n/a\", Name:\"azure-operator\", Source:\"https://github.com/giantswarm/azure-operator\"}","time":"2017-11-23 04:42:06.177"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/client/k8sclient/k8sclient.go:87","debug":"creating in-cluster config","time":"2017-11-23 04:42:06.178"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/microkit/server/server.go:297","debug":"running server at http://0.0.0.0:8000","time":"2017-11-23 04:42:06.201"}
{"caller":"github.com/giantswarm/azure-operator/service/operator/service.go:156","debug":"third party resource already exists","time":"2017-11-23 04:42:06.233"}
{"caller":"github.com/giantswarm/azure-operator/service/operator/service.go:161","debug":"starting list/watch","time":"2017-11-23 04:42:06.233"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/tpr/tpr.go:219","debug":"executing the reconciler's list function","event":"list","time":"2017-11-23 04:42:06.233"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/tpr/tpr.go:235","debug":"executing the reconciler's watch function","event":"watch","time":"2017-11-23 04:42:06.237"}
{"action":"start","caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:106","component":"operatorkit","function":"ProcessCreate","time":"2017-11-23 04:42:48.446"}
{"caller":"github.com/giantswarm/azure-operator/service/resource/resourcegroup/resource.go:172","cluster":"pawel-cluster2","debug":"creating Azure resource group","resource":"resourcegroup","time":"2017-11-23 04:42:49.895"}
{"caller":"github.com/giantswarm/azure-operator/service/resource/resourcegroup/resource.go:191","cluster":"pawel-cluster2","debug":"creating Azure resource group: created","resource":"resourcegroup","time":"2017-11-23 04:42:50.853"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:110","error":"[{/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:217: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:376: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/resource.go:156: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/deployment.go:26: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:75: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:142: timed out waiting for secrets} {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/error.go:14: secrets retreival failed}]","event":"create","time":"2017-11-23 04:44:21.138"}
{"action":"start","caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:187","component":"operatorkit","function":"ProcessUpdate","time":"2017-11-23 04:44:21.138"}
{"caller":"github.com/giantswarm/azure-operator/service/resource/resourcegroup/resource.go:172","cluster":"pawel-cluster2","debug":"creating Azure resource group","resource":"resourcegroup","time":"2017-11-23 04:44:22.206"}
{"caller":"github.com/giantswarm/azure-operator/service/resource/resourcegroup/resource.go:193","cluster":"pawel-cluster2","debug":"creating Azure resource group: already created","resource":"resourcegroup","time":"2017-11-23 04:44:22.206"}
{"caller":"github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:191","error":"[{/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:376: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/resource.go:156: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/deployment.go:26: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:75: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:142: timed out waiting for secrets} {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/error.go:14: secrets retreival failed}]","event":"update","time":"2017-11-23 04:45:52.600"}
```

To this:

```
D 17-11-23 04:42:06 creating azure-operator with config: service.Config{Logger:(*micrologger.logger)(0xc42022f280), Flag:(*flag.Flag)(0xc420175680), Viper:(*viper.Viper)(0xc4201af770), Description:"The azure-operator manages Kubernetes clusters on Azure.", GitCommit:"n/a", Name:"azure-operator", Source:"https://github.com/giantswarm/azure-operator"} | azure-operator/service/service.go:67
D 17-11-23 04:42:06 creating in-cluster config | operatorkit/client/k8sclient/k8sclient.go:87
D 17-11-23 04:42:06 running server at http://0.0.0.0:8000 | microkit/server/server.go:297
D 17-11-23 04:42:06 third party resource already exists | azure-operator/service/operator/service.go:156
D 17-11-23 04:42:06 starting list/watch | azure-operator/service/operator/service.go:161
D 17-11-23 04:42:06 executing the reconciler's list function | operatorkit/tpr/tpr.go:219 | event=list
D 17-11-23 04:42:06 executing the reconciler's watch function | operatorkit/tpr/tpr.go:235 | event=watch
U 17-11-23 04:42:48 | operatorkit/framework/framework.go:106 | action=start | component=operatorkit | function=ProcessCreate
D 17-11-23 04:42:49 creating Azure resource group | azure-operator/service/resource/resourcegroup/resource.go:172 | cluster=pawel-cluster2 | resource=resourcegroup
D 17-11-23 04:42:50 creating Azure resource group: created | azure-operator/service/resource/resourcegroup/resource.go:191 | cluster=pawel-cluster2 | resource=resourcegroup
E 17-11-23 04:44:21 [{/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:217: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:376: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/resource.go:156: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/service/resource/deployment/deployment.go:26: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:75: } {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/service.go:142: timed out waiting for secrets} {/Users/pawel/go/src/github.com/giantswarm/azure-operator/vendor/github.com/giantswarm/certificatetpr/error.go:14: secrets retreival failed}] | operatorkit/framework/framework.go:110 | event=create
```
