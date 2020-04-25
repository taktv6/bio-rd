package server

import (
	"fmt"
	"sync"
)

type netIfaManager struct {
	srv           *Server
	netIfas       map[string]*netIfa
	netIfasMu     sync.Mutex
	useMockTicker bool
}

func newNetIfaManager(srv *Server) *netIfaManager {
	return &netIfaManager{
		srv:     srv,
		netIfas: make(map[string]*netIfa),
	}
}

// AddInterface adds an interface to the ISIS server
func (s *Server) AddInterface(cfg *InterfaceConfig) error {
	return s.netIfaManager.addInterface(cfg)
}

func (nima *netIfaManager) addInterface(cfg *InterfaceConfig) error {
	nima.netIfasMu.Lock()
	defer nima.netIfasMu.Unlock()

	if _, exists := nima.netIfas[cfg.Name]; exists {
		return fmt.Errorf("ISIS is enabled on that interface already. Updating config is not supported yet")
	}

	ifa := newNetIfa(nima.srv, cfg)
	nima.netIfas[cfg.Name] = ifa

	return nil
}

func (nima *netIfaManager) getInterface(name string) *netIfa {
	nima.netIfasMu.Lock()
	defer nima.netIfasMu.Unlock()

	if _, found := nima.netIfas[name]; !found {
		return nil
	}

	return nima.netIfas[name]
}