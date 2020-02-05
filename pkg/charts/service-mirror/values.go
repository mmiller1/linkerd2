package service_mirror

import (
	"fmt"
	"github.com/linkerd/linkerd2/pkg/charts"
	"k8s.io/helm/pkg/chartutil"
	"sigs.k8s.io/yaml"
)

const (
	helmDefaultServiceMirrorChartDir = "linkerd2-service-mirror"
)

type Values struct {
	Namespace            string `json:"namespace"`
	ServiceMirrorVersion string `json:"serviceMirrorVersion"`
}

// NewValues returns a new instance of the Values type.
func NewValues() (*Values, error) {
	chartDir := fmt.Sprintf("%s/", helmDefaultServiceMirrorChartDir)
	v, err := readDefaults(chartDir)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// readDefaults read all the default variables from the values.yaml file.
// chartDir is the root directory of the Helm chart where values.yaml is.
func readDefaults(chartDir string) (*Values, error) {
	file := &chartutil.BufferedFile{
		Name: chartutil.ValuesfileName,
	}
	if err := charts.ReadFile(chartDir, file); err != nil {
		return nil, err
	}
	values := Values{}
	if err := yaml.Unmarshal(file.Data, &values); err != nil {
		return nil, err
	}
	return &values, nil
}
