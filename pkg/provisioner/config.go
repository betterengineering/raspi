// Copyright © 2018 Mark Spicer

package provisioner

import (
	homedir "github.com/mitchellh/go-homedir"
)

// Config is a configuration struct for the provisioner. This can be used to
// change things like the local NOOBs cache location or which URL to
// use for the NOOBs download.
type Config struct {
	// NoobsCachePath is the local path to store NOOBs zips.
	NoobsCachePath string

	// NoobsDownloadURL is the NOOBs download URL to use when updating the NOOBs
	// cache.
	NoobsDownloadURL string
}

// DefaultConfig returns a configuration struct populated with the defaults.
func DefaultConfig() (*Config, error) {
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	return &Config{
		NoobsCachePath:   home + "/.raspi/noobs-cache",
		NoobsDownloadURL: "https://downloads.raspberrypi.org/NOOBS_latest",
	}, nil
}
