package main

import (
	"fmt"

	diskutil "github.com/shirou/gopsutil/disk"
)

func main() {
	var diskIOReadBytes uint64
	var diskIOWriteBytes uint64
	var diskTotal uint64
	var diskUsed uint64

	diskio, err := diskutil.IOCounters([]string{}...)
	if err != nil {
		fmt.Println(" error: ", err)
	}

	for _, stat := range diskio {
		diskIOReadBytes += stat.ReadBytes
		diskIOWriteBytes += stat.WriteBytes
	}

	usage, err := diskutil.Usage("/")
	// usage, err := ipt.usage(partition.Mountpoint)
	if err != nil {
		fmt.Println("error", err)
		// return 0, 0, fmt.Errorf("error getting disk usage info: %w", err)
	}
	fmt.Println(usage.Total, usage.Used)
	diskTotal += usage.Total
	diskUsed += usage.Used
	fmt.Printf("ipt.usage : %d %d", usage.Total, usage.Used)

	partitions, err := diskutil.Partitions(false) // false: only physical partition
	if err != nil {
		fmt.Println(" error: ", err)
	}

	for i, partition := range partitions {
		// usage, err := ipt.usage("/")
		usage, err := diskutil.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println(" error: ", err)
		} else {
			fmt.Printf("ipt.usage :%d %d %d %s\n", i, usage.Total, usage.Used, partition.Mountpoint)
			diskTotal += usage.Total
			diskUsed += usage.Used
		}

	}

}
