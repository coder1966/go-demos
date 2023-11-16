package main

import (
	"math"
	"os"
	"strings"

	//nolint
	"github.com/shirou/gopsutil/disk"
)

var diskStats PSDiskStats

func main() {
	diskStats = &PSDisk{}

	disks, partitions, err := diskStats.FilterUsage()

	_, _, _ = disks, partitions, err

	return
}

type PSDiskStats interface {
	Usage(path string) (*disk.UsageStat, error)
	FilterUsage() ([]*disk.UsageStat, []*disk.PartitionStat, error)
	OSGetenv(key string) string
	Partitions(all bool) ([]disk.PartitionStat, error)
}

type PSDisk struct {
	// ipt *Input
}

func (dk *PSDisk) Usage(path string) (*disk.UsageStat, error) {
	return disk.Usage(path)
}

func (dk *PSDisk) OSGetenv(key string) string {
	return os.Getenv(key)
}

func (dk *PSDisk) Partitions(all bool) ([]disk.PartitionStat, error) {
	return disk.Partitions(all)
}

func (dk *PSDisk) FilterUsage() ([]*disk.UsageStat, []*disk.PartitionStat, error) {
	parts, err := dk.Partitions(true)
	for _, v := range parts {
		println(v.String())
	}
	if err != nil {
		return nil, nil, err
	}

	// excluded := func(x string, arr []string) bool {
	// 	for _, fs := range arr {
	// 		if strings.EqualFold(x, fs) {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }

	var usage []*disk.UsageStat
	var partitions []*disk.PartitionStat

	for i := range parts {
		p := parts[i]

		if strings.HasPrefix(p.Device, "/rootfs/dev/") {
			p.Device = p.Device[7:]
		}

		// l.Debugf("disk---fstype:%s ,device:%s ,mountpoint:%s ", p.Fstype, p.Device, p.Mountpoint)
		// // nolint
		// if !strings.HasPrefix(p.Device, "/dev/") && runtime.GOOS != datakit.OSWindows && !excluded(p.Device, dk.ipt.ExtraDevice) {
		// 	continue // ignore the partition
		// }

		// if excluded(p.Device, dk.ipt.ExcludeDevice) {
		// 	continue
		// }

		mergerFlag := false
		// merger device
		for _, cont := range partitions {
			if cont.Device == p.Device {
				mergerFlag = true
				break
			}
		}

		if mergerFlag {
			continue
		}

		du, err := dk.Usage(p.Mountpoint)
		// for _, v := range du {
		// 	println(v.String())
		// }
		if err != nil {
			continue
		}

		du.Fstype = p.Fstype

		usage = append(usage, du)
		partitions = append(partitions, &p)
	}

	return usage, partitions, nil

	return nil, nil, nil
}

type MountOptions []string

func (opts MountOptions) Mode() string {
	switch {
	case opts.exists("rw"):
		return "rw"
	case opts.exists("ro"):
		return "ro"
	default:
		return "unknown"
	}
}

func (opts MountOptions) exists(opt string) bool {
	for _, o := range opts {
		if o == opt {
			return true
		}
	}
	return false
}

func wrapUint64(x uint64) int64 {
	if x > uint64(math.MaxInt64) {
		return -1
	}
	return int64(x)
}

func unique(strSlice []string) []string {
	keys := make(map[string]interface{})
	var list []string
	for _, entry := range strSlice {
		if _, ok := keys[entry]; !ok {
			keys[entry] = nil
			list = append(list, entry)
		}
	}
	return list
}
