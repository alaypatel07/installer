package main

import (
	"go.etcd.io/etcd/pkg/mock/mockserver"
	"testing"
)

//func Test_waitForEtcdCluster(t *testing.T) {
//
//	replicas := new(int64)
//	dummyAsset := &installconfig.InstallConfig{
//		Config: &types.InstallConfig{
//			ObjectMeta: v1.ObjectMeta{
//				Name: "foo",
//			},
//			BaseDomain: "bar.com",
//			ControlPlane: &types.MachinePool{
//				Name:     "dummy-pool",
//				Replicas: replicas,
//			},
//		},
//	}
//	type args struct {
//		ctx       context.Context
//		asset     asset.WritableAsset
//		directory string
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		{
//			name: "valid testcase",
//			args: args{
//				ctx:   nil,
//				asset: dummyAsset,
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := waitForEtcdCluster(tt.args.ctx, tt.args.asset, tt.args.directory); (err != nil) != tt.wantErr {
//				t.Errorf("waitForEtcdCluster() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

func Test_isEtcdHealthy(t *testing.T) {
	ms, err := mockserver.StartMockServers(3)
	if err != nil {
		t.Fatal(err)
	}
	defer ms.Stop()

	ep := []string{}

	for _, s := range ms.Servers {
		ep = append(ep, s.Address)
	}

	type args struct {
		endpoints []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "cluster is ready",
			args: args{endpoints: ep},
			want: true,
		},
		{
			name: "cluster is unready",
			args: args{endpoints: ep},
			want: false,
		},
		{
			name: "cluster is ready again",
			args: args{endpoints: ep},
			want: true,
		},
	}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := isEtcdHealthy(tt.args.endpoints); got != tt.want {
	//			t.Errorf("isEtcdHealthy() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}

	if got := isEtcdHealthy(tests[0].args.endpoints); got != tests[0].want {
		t.Errorf("isEtcdHealthy() = %v, want %v", got, tests[0].want)
	}
	ms.StopAt(1)

	if got := isEtcdHealthy(tests[1].args.endpoints); got != tests[1].want {
		t.Errorf("isEtcdHealthy() = %v, want %v", got, tests[1].want)
	}

	err = ms.StartAt(1)
	if err != nil {
		t.Errorf("Error starting etcd mock server %v\n", err)
	}

	if got := isEtcdHealthy(tests[2].args.endpoints); got != tests[2].want {
		t.Errorf("isEtcdHealthy() = %v, want %v", got, tests[2].want)
	}
}
