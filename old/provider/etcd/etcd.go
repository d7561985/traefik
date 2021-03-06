package etcd

import (
	"fmt"

	"github.com/abronan/valkeyrie/store"
	etcdv3 "github.com/abronan/valkeyrie/store/etcd/v3"
	"github.com/containous/traefik/old/provider"
	"github.com/containous/traefik/old/provider/kv"
	"github.com/containous/traefik/old/types"
	"github.com/containous/traefik/safe"
)

var _ provider.Provider = (*Provider)(nil)

// Provider holds configurations of the provider.
type Provider struct {
	kv.Provider `mapstructure:",squash" export:"true"`
}

// Init the provider
func (p *Provider) Init(constraints types.Constraints) error {
	err := p.Provider.Init(constraints)
	if err != nil {
		return err
	}

	store, err := p.CreateStore()
	if err != nil {
		return fmt.Errorf("failed to Connect to KV store: %v", err)
	}

	p.SetKVClient(store)
	return nil
}

// Provide allows the etcd provider to Provide configurations to traefik
// using the given configuration channel.
func (p *Provider) Provide(configurationChan chan<- types.ConfigMessage, pool *safe.Pool) error {
	return p.Provider.Provide(configurationChan, pool)
}

// CreateStore creates the KV store
func (p *Provider) CreateStore() (store.Store, error) {
	etcdv3.Register()
	p.SetStoreType(store.ETCDV3)
	return p.Provider.CreateStore()
}
