// Copyright © 2018 Mark Spicer

// Package provisioner provides an interface to provision a new Raspberry Pi
// with NOOBs, wireless configuration, and discovery.
package provisioner

// Provisioner is an interface used to provision a new Raspberry Pi with NOOBs,
// wireless configuration, and discovery.
type Provisioner interface {
	// CleanNoobsCache will clean the local NOOBs cache to free storage.
	CleanNoobsCache() error

	// UpdateNoobsCache will download the latest version of NOOBs to the local
	// cache.
	UpdateNoobsCache() error

	PartitionDisk(device string) error
}
