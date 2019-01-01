// Copyright © 2018 Mark Spicer

package provisioner

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// NoobsZipPrefix is the NOOBs zip prefix to use when downloading new
	// versions of NOOBs.
	NoobsZipPrefix = "noobs-"

	// NoobsZipExt is the NOOBs file extension to use for the downloaded zips.
	NoobsZipExt = ".zip"
)

// BasicProvisioner is an implementation of the Provisioner interface. I don't
// imagine there will be other implementations of this interface, but wanted a
// way to mock out the provisioner so that consumers of this package do not need
// to interact with the filesystem or deal with remote HTTP calls.
type BasicProvisioner struct {
	Config *Config
}

// NewBasicProvisioner provides an initialized BasicProvsioner with the provided
// configuration.
func NewBasicProvisioner(config *Config) *BasicProvisioner {
	return &BasicProvisioner{
		Config: config,
	}
}

// CleanNoobsCache will purge the local NOOBs cache to free local storage space.
func (prov *BasicProvisioner) CleanNoobsCache() error {
	log.WithFields(log.Fields{
		"cachePath": prov.Config.NoobsCachePath,
	}).Info("cleaning noobs cache")

	return os.RemoveAll(prov.Config.NoobsCachePath)
}

// UpdateNoobsCache will fetch a fresh version of NOOBs from the configured URL
// and store it in the NOOBs cache.
func (prov *BasicProvisioner) UpdateNoobsCache() error {
	log.WithFields(log.Fields{
		"downloadURL": prov.Config.NoobsDownloadURL,
		"cachePath":   prov.Config.NoobsCachePath,
	}).Info("updating NOOBs cache")

	err := prov.ensureNoobsCacheDirExists()
	if err != nil {
		return err
	}

	err = prov.downloadNoobsZip()
	if err != nil {
		return err
	}

	return nil
}

func (prov *BasicProvisioner) ensureNoobsCacheDirExists() error {
	return os.MkdirAll(prov.Config.NoobsCachePath, 0755)
}

func (prov *BasicProvisioner) generateNoobsFilename() string {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return prov.Config.NoobsCachePath + "/" + NoobsZipPrefix + timestamp + NoobsZipExt
}

func (prov *BasicProvisioner) downloadNoobsZip() error {
	zip, err := os.Create(prov.generateNoobsFilename())
	if err != nil {
		return err
	}
	defer zip.Close()

	log.WithFields(log.Fields{
		"downloadURL":  prov.Config.NoobsDownloadURL,
		"downloadPath": zip.Name(),
	}).Info("dowloading latest NOOBs zip, this will take a few minutes")

	resp, err := http.Get(prov.Config.NoobsDownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(zip, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
