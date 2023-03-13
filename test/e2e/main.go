package e2e

import (
	ginkgo "github.com/onsi/ginkgo/v2"
	gomega "github.com/onsi/gomega"
	"go.uber.org/zap/zapcore"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"time"
)

//var (
//	cfg       *rest.Config
//	k8sClient client.Client
//	testEnv   *envtest.Environment
//	ctx       context.Context
//	cancel    context.CancelFunc
//)

var _ = ginkgo.BeforeSuite(func() {
	//ctx, cancel = context.WithCancel(context.TODO())

	// Initialize logger
	opts := zap.Options{
		Development: true,
		TimeEncoder: zapcore.TimeEncoderOfLayout(time.RFC3339),
	}
	logf.SetLogger(zap.New(zap.WriteTo(ginkgo.GinkgoWriter), zap.UseFlagOptions(&opts)))
	logf.Log.Info(gomega.GOMEGA_VERSION)
})
