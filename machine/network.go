package machine

import (
	"fmt"
	"github.com/iesreza/netconfig"
	"net"
)

var netcfg = netconfig.GetNetworkConfig()

func ActiveNetIFace() (*net.Interface, error) {

	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.HardwareAddr.String() == netcfg.HardwareAddress.String() {
			return &iface, nil
		}
	}

	return nil, fmt.Errorf("unable to find active device")
}

func NetworkConfig() *netconfig.Network {
	return netcfg
}
