package rismirror

import (
	"net"

	"github.com/bio-routing/bio-rd/routingtable/vrf"
)

// Router represents a router
type Router struct {
	name        string
	address     net.IP
	vrfRegistry *vrf.VRFRegistry
}

func newRouter(name string, address net.IP) *Router {
	return &Router{
		name:        name,
		address:     address,
		vrfRegistry: vrf.NewVRFRegistry(),
	}
}

// Name gets the routers name
func (r *Router) Name() string {
	return r.name
}

// Address gets a routers address
func (r *Router) Address() net.IP {
	return r.address
}

// GetVRF gets a VRF by its ID
func (r *Router) GetVRF(vrfID uint64) *vrf.VRF {
	return r.vrfRegistry.GetVRFByRD(vrfID)
}

// GetVRFs gets all VRFs
func (r *Router) GetVRFs() []*vrf.VRF {
	return r.vrfRegistry.List()
}
