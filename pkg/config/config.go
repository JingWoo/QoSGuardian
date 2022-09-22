package config

import (
	"fmt"
	"io/ioutil"
	"time"

	yaml "gopkg.in/yaml.v3"
	log "k8s.io/klog/v2"
)

type NodeMetricsCollectorConfig struct {
	CollectInterval time.Duration `yaml:"collect_interval"`
}

type PodMetricsCollectorConfig struct {
	CollectInterval time.Duration `yaml:"collect_interval"`
}

type MetricsCollectorConfig struct {
	Node NodeMetricsCollectorConfig `yaml:"node"`
	Pod  PodMetricsCollectorConfig  `yaml:"pod"`
}

type QoSGuardianConfig struct {
	MetricsCollectorConfig MetricsCollectorConfig `yaml:"metrics_collector"`
}

func ParseQoSGuardianConfig(file string) (*QoSGuardianConfig, error) {
	rawConfigFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to read config file: %v", file)
	}

	var qosGuardianConfig QoSGuardianConfig
	if err := yaml.Unmarshal(rawConfigFile, &qosGuardianConfig); err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("failed to unmarshal config file: %v", file)
	}

	return &qosGuardianConfig, nil
}
