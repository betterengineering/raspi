// Copyright © 2018 Mark Spicer

package provisioner_test

import (
	"testing"

	"github.com/lodge93/raspi/pkg/provisioner"
	homedir "github.com/mitchellh/go-homedir"
)

const (
	DefaultNoobsCachePath   = "/.raspi/noobs-cache"
	DefaultNoobsDownloadURL = "https://downloads.raspberrypi.org/NOOBS_latest"
)

func TestDeafultConfig(t *testing.T) {
	config, err := provisioner.DefaultConfig()
	if err != nil {
		t.Fatal("could not instatiate default configuration:", err)
	}

	home, err := homedir.Dir()
	if err != nil {
		t.Fatal("could not locate home dir:", err)
	}

	if config.NoobsCachePath != (home + DefaultNoobsCachePath) {
		t.Error("NoobsCachePath does not match default.")
	}

	if config.NoobsDownloadURL != DefaultNoobsDownloadURL {
		t.Error("NoobsDownloadURL does not match default.")
	}
}
