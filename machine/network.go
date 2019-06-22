package machine

import (
	"fmt"
	"github.com/iesreza/netconfig"
	"net"
)

func ActiveNetIFace() (*net.Interface, error) {
	cfg := netconfig.GetNetworkConfig()
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	fmt.Println()
	for _, iface := range ifaces {
		if iface.HardwareAddr.String() == cfg.HardwareAddress.String() {
			return &iface, nil
		}
	}

	return nil, fmt.Errorf("unable to find active device")
}
