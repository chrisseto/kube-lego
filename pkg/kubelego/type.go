package kubelego

import (
	"net"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jetstack/kube-lego/pkg/ingress"
	"github.com/jetstack/kube-lego/pkg/kubelego_const"

	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/workqueue"
)

type KubeLego struct {
	legoURL                      string
	legoEmail                    string
	legoSecretName               string
	legoIngressNameNginx         string
	legoNamespace                string
	legoPodIP                    net.IP
	legoServiceNameNginx         string
	legoServiceNameGce           string
	legoSupportedIngressClass    []string
	legoSupportedIngressProvider []string
	legoHTTPPort                 intstr.IntOrString
	legoCheckInterval            time.Duration
	legoMinimumValidity          time.Duration
	legoDefaultIngressClass      string
	legoDefaultIngressProvider   string
	legoKubeApiURL               string
	legoWatchNamespace           string
	kubeClient                   *kubernetes.Clientset
	legoIngressSlice             []*ingress.Ingress
	legoIngressProvider          map[string]kubelego.IngressProvider
	log                          *log.Entry
	version                      string
	acmeClient                   kubelego.Acme

	// stop channel for services
	stopCh chan struct{}

	// wait group
	waitGroup sync.WaitGroup

	// work queue
	workQueue *workqueue.Type
}
