package providers

import (
	"github.com/debarshibasak/go-kubeadmclient/kubeadmclient"
	"github.com/debarshibasak/kubestrike/v1alpha1/config"
	"github.com/pkg/errors"
)

type Provider string

const (
	MultipassProvider Provider = "Multipass"
	BaremetalProvider Provider = "Baremetal"
)

type Providers interface {
	Provision() ([]*kubeadmclient.MasterNode, []*kubeadmclient.WorkerNode, *kubeadmclient.HaProxyNode, error)
}

func Get(orchestrator *config.ClusterOrchestrator) ([]*kubeadmclient.MasterNode, []*kubeadmclient.WorkerNode, *kubeadmclient.HaProxyNode, error) {

	switch orchestrator.Provider {
	case MultipassProvider:
		{
			return orchestrator.Multipass.Provision()
		}
	case BaremetalProvider:
		{
			return orchestrator.BareMetal.Provision()
		}
	}

	return nil, nil, nil, errors.New("provisioner not found")
}
