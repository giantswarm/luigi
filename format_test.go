package main

import (
	"testing"
)

func Test_format(t *testing.T) {
	testCases := []struct {
		name        string
		text        string
		expectedOut string
	}{
		{
			name: "case 0",
			text: `{"caller":"github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:174","event":"update","function":"UpdateFunc","level":"error","message":"stop framework reconciliation due to error","object":"/apis/provider.giantswarm.io/v1alpha1/namespaces/default/awsconfigs/3cnwh","stack":"[{/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:420: } {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/metricsresource/resource.go:73: } {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/retryresource/resource.go:87: } {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/retryresource/resource.go:75: } {/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v5/resource/loadbalancer/current.go:26: } {/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v5/resource/loadbalancer/current.go:54: } {ValidationError: 1 validation error detected: Value '[f1ugn-api, f1ugn-etcd, f1ugn-ingress, pqtq7-api, pqtq7-etcd, pqtq7-ingress, y22nx-api, y22nx-etcd, y22nx-ingress, es-vpn-l7-router-lb, 20j3h-api, 20j3h-etcd, 20j3h-ingress, rue99-api, rue99-etcd, rue99-ingress, 2g0sc-api, 2g0sc-etcd, 2g0sc-ingress, g11pa-api, g11pa-etcd, g11pa-ingress, 2g0sc-ingress-internal-apix-dxl, g11pa-ingress-internal-apix-gig, 6u2fa-api, 6u2fa-etcd, 6u2fa-ingress, 0gpmx-api, 0gpmx-etcd, 0gpmx-ingress, a91dab3e2df5d11e797fc02d8eb724b1, a085c7ba1e6f111e792ea026496c10d9, acbffaa65f61511e79916024a6fdab53, yq237-ingress, yq237-api, 85avm-ingress, 85avm-api, 2sq4i-ingress, 2sq4i-api, a2dfac1e10db611e88ea9025ad9ed0da, do8cv-ingress, do8cv-api, a92958c09133511e88e3802539833dd4, 9gvzw-ingress, 9gvzw-api, a30365b16171611e88ea9025ad9ed0da, a20ad4a69159611e88e3802539833dd4, aaab9c44717cc11e88ea9025ad9ed0da, a2033409f159611e88e3802539833dd4, 5i746-ingress, 5i746-api, ac248d803194111e8a1eb02374d0af04, ac267b5d1194111e8a1eb02374d0af04, ac2a46479194111e8a1eb02374d0af04, a302acebf1ad311e8896402d8eb724b1, es-vpn-inbound-l7-router-lb]' at 'loadBalancerNames' failed to satisfy constraint: Member must have length less than or equal to 20\n\tstatus code: 400, request id: 2184a219-1b89-11e8-a0c2-21a5ccc5b186}]","time":"2018-02-27 06:40:42.169"}`,
			expectedOut: `E 18-02-27 06:40:42 /apis/provider.giantswarm.io/v1alpha1/namespaces/default/awsconfigs/3cnwh stop framework reconciliation due to error | operatorkit/framework/framework.go:174 | event=update | function=UpdateFunc
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:420: 
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/metricsresource/resource.go:73: 
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/retryresource/resource.go:87: 
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/retryresource/resource.go:75: 
	/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v5/resource/loadbalancer/current.go:26: 
	/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v5/resource/loadbalancer/current.go:54: 
	ValidationError: 1 validation error detected: Value '[f1ugn-api, f1ugn-etcd, f1ugn-ingress, pqtq7-api, pqtq7-etcd, pqtq7-ingress, y22nx-api, y22nx-etcd, y22nx-ingress, es-vpn-l7-router-lb, 20j3h-api, 20j3h-etcd, 20j3h-ingress, rue99-api, rue99-etcd, rue99-ingress, 2g0sc-api, 2g0sc-etcd, 2g0sc-ingress, g11pa-api, g11pa-etcd, g11pa-ingress, 2g0sc-ingress-internal-apix-dxl, g11pa-ingress-internal-apix-gig, 6u2fa-api, 6u2fa-etcd, 6u2fa-ingress, 0gpmx-api, 0gpmx-etcd, 0gpmx-ingress, a91dab3e2df5d11e797fc02d8eb724b1, a085c7ba1e6f111e792ea026496c10d9, acbffaa65f61511e79916024a6fdab53, yq237-ingress, yq237-api, 85avm-ingress, 85avm-api, 2sq4i-ingress, 2sq4i-api, a2dfac1e10db611e88ea9025ad9ed0da, do8cv-ingress, do8cv-api, a92958c09133511e88e3802539833dd4, 9gvzw-ingress, 9gvzw-api, a30365b16171611e88ea9025ad9ed0da, a20ad4a69159611e88e3802539833dd4, aaab9c44717cc11e88ea9025ad9ed0da, a2033409f159611e88e3802539833dd4, 5i746-ingress, 5i746-api, ac248d803194111e8a1eb02374d0af04, ac267b5d1194111e8a1eb02374d0af04, ac2a46479194111e8a1eb02374d0af04, a302acebf1ad311e8896402d8eb724b1, es-vpn-inbound-l7-router-lb]' at 'loadBalancerNames' failed to satisfy constraint: Member must have length less than or equal to 20
	status code: 400, request id: 2184a219-1b89-11e8-a0c2-21a5ccc5b186`,
		},
		{
			name: "case 1",
			text: `{"caller":"github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:174","event":"update","function":"UpdateFunc","level":"error","message":"stop framework reconciliation due to error","object":"/apis/provider.giantswarm.io/v1alpha1/namespaces/default/awsconfigs/unoo8","stack":"[{/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:462: } {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/metricsresource/resource.go:111: } {/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v1/resource/legacy/resource.go:178: } {/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v1/resource/legacy/resource.go:287: could not get keys from secrets: '[{/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/service.go:73: } {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/service.go:142: timed out waiting for secrets} {/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/error.go:14: secrets retreival failed}]'} {execution failed}]","time":"2018-02-27 14:26:24.612"}`,
			expectedOut: `E 18-02-27 14:26:24 /apis/provider.giantswarm.io/v1alpha1/namespaces/default/awsconfigs/unoo8 stop framework reconciliation due to error | operatorkit/framework/framework.go:174 | event=update | function=UpdateFunc
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/framework.go:462: 
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/operatorkit/framework/resource/metricsresource/resource.go:111: 
	/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v1/resource/legacy/resource.go:178: 
	/go/src/github.com/giantswarm/aws-operator/service/awsconfig/v1/resource/legacy/resource.go:287: could not get keys from secrets: '[{/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/service.go:73: 
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/service.go:142: timed out waiting for secrets
	/go/src/github.com/giantswarm/aws-operator/vendor/github.com/giantswarm/randomkeytpr/error.go:14: secrets retreival failed}]'
	execution failed`,
		},
	}

	disableColors(true)

	for i, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := format([]byte(tc.text))
			if out != tc.expectedOut {
				t.Errorf("test %d:\n%q\nwant:\n%q", i, out, tc.expectedOut)
			}
		})
	}
}
