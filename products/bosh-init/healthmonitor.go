package boshinit

import "github.com/enaml-ops/omg-cli/plugins/products/bosh-init/enaml-gen/health_monitor"

func NewHealthMonitor(resurrectorEnabled bool) health_monitor.Hm {
	return health_monitor.Hm{
		DirectorAccount: &health_monitor.DirectorAccount{
			User:     "hm",
			Password: "hm-password",
		},
		ResurrectorEnabled: resurrectorEnabled,
	}
}
