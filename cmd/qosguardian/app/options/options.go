package options

import "github.com/spf13/pflag"

const defaultQoSGuardianConfig = "/etc/sysconfig/QoSGuardian.json"

type QoSGuardianServer struct {
	Config string
}

func NewQoSGuardianServer() *QoSGuardianServer {
	return &QoSGuardianServer{}
}

func (opts *QoSGuardianServer) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&opts.Config, "config", defaultQoSGuardianConfig, "Path to config file of QosGuardian Server ")
}
