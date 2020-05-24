/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// apiserver is the main api server and master for the cluster.
// it is responsible for serving the cluster management API.
/**
## 作用：
 1. 整个集群管理的 API 接口：所有对集群进行的查询和管理都要通过 API 来进行
 2. 集群内部各个模块之间通信的枢纽：所有模块之间并不会直接互相调用，而是通过和 API Server 打交道来完成自己那部分的工作
 3. 集群安全控制：API Server 提供的验证和授权保证了整个集群的安全
 */
package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-apiserver/app"
	_ "k8s.io/kubernetes/pkg/util/prometheusclientgo" // load all the prometheus client-go plugins
	_ "k8s.io/kubernetes/pkg/version/prometheus"      // for version metric registration
)

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewAPIServerCommand()

	// TODO: once we switch everything over to Cobra commands, we can go back to calling
	// utilflag.InitFlags() (by removing its pflag.Parse() call). For now, we have to set the
	// normalize func and add the go flag set by hand.
	// utilflag.InitFlags()
	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
