// Copyright © 2018 Mark Spicer

package provisioner_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/lodge93/raspi/pkg/provisioner"
	"github.com/sirupsen/logrus"
)

const (
	TestNoobsDownloadURL = "https://www.google.com/robots.txt"
)

func init() {
	logrus.SetLevel(logrus.PanicLevel)
}

func SetupTest() (*provisioner.BasicProvisioner, error) {
	tempDir, err := ioutil.TempDir("", "raspi-provisioner-test-dir")
	if err != nil {
		return nil, err
	}

	prov := provisioner.NewBasicProvisioner(&provisioner.Config{
		NoobsCachePath:   tempDir,
		NoobsDownloadURL: TestNoobsDownloadURL,
	})

	return prov, nil
}

func TearDownTest(prov *provisioner.BasicProvisioner) {
	os.RemoveAll(prov.Config.NoobsCachePath)
}

func TestCleanNoobsCache(t *testing.T) {
	prov, err := SetupTest()
	if err != nil {
		t.Fatal("could not setup test for basic provisioner")
	}
	defer TearDownTest(prov)

	err = prov.CleanNoobsCache()
	if err != nil {
		t.Error("CleanNoobsCache returned with error:", err)
	}

	_, err = os.Stat(prov.Config.NoobsCachePath)
	if !os.IsNotExist(err) {
		t.Error("base image dir exists after it was cleaned")
	}
}

func TestUpdateNoobsCache(t *testing.T) {
	prov, err := SetupTest()
	if err != nil {
		t.Fatal("could not setup test for basic provisioner")
	}
	defer TearDownTest(prov)

	err = prov.UpdateNoobsCache()
	if err != nil {
		t.Error("UpdateNoobsCache returned with error:", err)
	}

	pattern := prov.Config.NoobsCachePath + "/" + provisioner.NoobsZipPrefix + "*" + provisioner.NoobsZipExt

	matches, err := filepath.Glob(pattern)
	if err != nil {
		t.Fatal("could not search for file son the file system:", err)
	}

	if len(matches) != 1 {
		t.Error("could not find file downloaded")
	}
}
