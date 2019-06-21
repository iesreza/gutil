package machine

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/denisbrodbeck/machineid"
	"github.com/iesreza/netconfig"
)

func UniqueID() string {
	id, _ := machineid.ID()
	cfg := netconfig.GetNetworkConfig()
	id += cfg.HardwareAddress.String()
	hasher := md5.New()
	hasher.Write([]byte(id))
	return hex.EncodeToString(hasher.Sum(nil))
}
