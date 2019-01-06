// Copyright © 2018 Mark Spicer

package main

import (
	"log"
	"syscall"

	"github.com/lodge93/raspi/pkg/provisioner"
)

func main() {
	// TODO: unmount partitions if mounted
	err := syscall.Unmount("/tmp/mspicer", int(1))
	if err != nil {
		log.Fatal(err)
	}

	config, err := provisioner.DefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	prov := provisioner.NewBasicProvisioner(config)
	err = prov.PartitionDisk("/dev/mmcblk0")
	if err != nil {
		log.Fatal(err)
	}

	err := syscall.Mount("/dev/mmcblk0p1", "/tmp/mspicer", int(1))
	if err != nil {
		log.Fatal(err)
	}
}
