package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/providers"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/tfvars"
	"github.com/ppc64le-cloud/kubetest2-plugins/pkg/utils"
	"github.com/spf13/pflag"
)

const (
	Name = "common"
)

var _ providers.Provider = &Provider{}

var CommonProvider = &Provider{}

type Provider struct {
	tfvars.TFVars
}

func (p *Provider) BindFlags(flags *pflag.FlagSet) {
	flags.StringVar(
		&p.ReleaseMarker, "release-marker", "ci/latest", "Kubernetes Release Marker",
	)
	flags.StringVar(
		&p.BuildVersion, "build-version", "", "Kubernetes Build Version",
	)
	flags.StringVar(
		&p.ClusterName, "cluster-name", "", "Kubernetes Cluster Name, this will used for creating the nodes and directories etc(Default: autogenerated with k8s-cluster-<6letters>",
	)
	flags.IntVar(
		&p.ApiServerPort, "apiserver-port", 992, "API Server Port Address",
	)
	flags.IntVar(
		&p.WorkersCount, "workers-count", 1, "Numbers of workers in the k8s cluster",
	)
	flags.StringVar(
		&p.BootstrapToken, "bootstrap-token", "", "Kubeadm bootstrap token used for installing and joining the cluster(default: random generated token in [a-z0-9]{6}\\.[a-z0-9]{16} format)",
	)
	flags.StringVar(
		&p.KubeconfigPath, "kubeconfig-path", "", "File path to write the kubeconfig content for the deployed cluster(default: data folder where terraform files copied)",
	)
}

func (p *Provider) DumpConfig(dir string) error {
	filename := path.Join(dir, Name+".auto.tfvars.json")

	config, err := json.MarshalIndent(p.TFVars, "", "  ")
	if err != nil {
		return fmt.Errorf("errored file converting config to json: %v", err)
	}

	err = ioutil.WriteFile(filename, config, 0644)
	if err != nil {
		return fmt.Errorf("failed to dump the json config to: %s, err: %v", filename, err)
	}

	return nil
}

func (p *Provider) Initialize() error {
	if p.ClusterName == "" {
		randPostFix, err := utils.RandString(6)
		if err != nil {
			return fmt.Errorf("failed to generate a random string, error: %v", err)
		}
		p.ClusterName = "k8s-cluster-" + randPostFix
	}
	if p.BootstrapToken == "" {
		bootstrapToken, err := utils.GenerateBootstrapToken()
		if err != nil {
			return fmt.Errorf("failed to generate a random string, error: %v", err)
		}
		p.BootstrapToken = bootstrapToken
	}
	if p.KubeconfigPath == "" {
		p.KubeconfigPath = path.Join(p.ClusterName, "kubeconfig")
	}
	return nil
}
