/*
Copyright 2023.

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

package main

import (
	"flag"
	"fmt"
	"os"
	go_runtime "runtime"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	"github.com/mahdi8731/redis-cluster-operator/pkg/apis"
	redisv1alpha1 "github.com/mahdi8731/redis-cluster-operator/pkg/apis/redis/v1alpha1"
	"github.com/mahdi8731/redis-cluster-operator/pkg/controller"
	"github.com/mahdi8731/redis-cluster-operator/pkg/controller/distributedrediscluster"
	"github.com/mahdi8731/redis-cluster-operator/pkg/controller/redisclusterbackup"
	"github.com/mahdi8731/redis-cluster-operator/pkg/k8sutil"
	"github.com/mahdi8731/redis-cluster-operator/pkg/utils"
	"github.com/mahdi8731/redis-cluster-operator/version"
	"github.com/spf13/pflag"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/healthz"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	//+kubebuilder:scaffold:imports

	config2 "github.com/mahdi8731/redis-cluster-operator/pkg/config"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

var log = logf.Log.WithName("cmd")

// var (
// 	metricsHost               = "0.0.0.0"
// 	metricsPort         int32 = 8383
// 	operatorMetricsPort int32 = 8686
// )

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	//+kubebuilder:scaffold:scheme
}

func printVersion() {
	log.Info(fmt.Sprintf("Go Version: %s", go_runtime.Version()))
	log.Info(fmt.Sprintf("Go OS/Arch: %s/%s", go_runtime.GOOS, go_runtime.GOARCH))
	// log.Info(fmt.Sprintf("Version of operator-sdk: %v", sdkVersion.Version))
	log.Info(fmt.Sprintf("Version of operator: %s+%s", version.Version, version.GitSHA))
}

func main() {
	var metricsAddr string = "0.0.0.0:8383"
	var enableLeaderElection bool = true
	var probeAddr string = "0.0.0.0:8585"
	flag.StringVar(&metricsAddr, "metrics-bind-address", ":8080", "The address the metric endpoint binds to.")
	flag.StringVar(&probeAddr, "health-probe-bind-address", ":8081", "The address the probe endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "leader-elect", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	opts := zap.Options{
		Development: false,
	}
	ctrl.SetLogger(zap.New(zap.UseFlagOptions(&opts)))
	// opts.BindFlags(flag.CommandLine)

	pflag.CommandLine.AddFlagSet(distributedrediscluster.FlagSet())
	pflag.CommandLine.AddFlagSet(redisclusterbackup.FlagSet())

	// Add flags registered by imported packages (e.g. glog and
	// controller-runtime)
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	log.Info("****************test***************")

	config2.RedisConf().AddFlags(pflag.CommandLine)

	log.Info("****************test***************")

	pflag.Parse()

	log.Info("****************test***************")

	printVersion()

	namespace, err := k8sutil.GetWatchNamespace()
	if err != nil {
		log.Error(err, "Failed to get watch namespace")
		os.Exit(1)
	}

	utils.SetClusterScoped(namespace)

	// Get a config to talk to the apiserver
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// ctx := context.TODO()
	// Become the leader before proceeding
	// err = leader.Become(ctx, "redis-cluster-operator-lock")
	// if err != nil {
	// 	log.Error(err, "")
	// 	os.Exit(1)
	// }

	mgr, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     metricsAddr,
		Port:                   9443,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "redis-cluster-operator-lock",
		Namespace:              namespace,
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	//+kubebuilder:scaffold:builder

	if err := mgr.AddHealthzCheck("healthz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up health check")
		os.Exit(1)
	}
	if err := mgr.AddReadyzCheck("readyz", healthz.Ping); err != nil {
		setupLog.Error(err, "unable to set up ready check")
		os.Exit(1)
	}

	// old code

	log.Info("Registering Components.")

	// Setup Scheme for all resources
	if err := apis.AddToScheme(mgr.GetScheme()); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	// Setup all Controllers
	if err := controller.AddToManager(mgr); err != nil {
		log.Error(err, "")
		os.Exit(1)
	}

	if os.Getenv("ENABLE_WEBHOOKS") == "true" {
		log.Info("Starting the WebHook.")
		ws := mgr.GetWebhookServer()
		ws.CertDir = "/etc/webhook/certs"
		ws.Port = 7443
		if err = (&redisv1alpha1.DistributedRedisCluster{}).SetupWebhookWithManager(mgr); err != nil {
			log.Error(err, "unable to create webHook", "webHook", "DistributedRedisCluster")
			os.Exit(1)
		}
	}

	// log.Info("Starting the Cmd.")

	// // Start the Cmd
	// if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
	// 	log.Error(err, "Manager exited non-zero")
	// 	os.Exit(1)
	// }

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}

}
