package main

import (
	"github.com/spf13/pflag"
)

func addCommonFlags(flagSet *pflag.FlagSet) *container {
	c := &container{}

	// please add the following flag by name in alphabetical order
	// blkio
	flagSet.Uint16Var(&c.blkioWeight, "blkio-weight", 0, "Block IO (relative weight), between 10 and 1000, or 0 to disable")
	flagSet.Var(&c.blkioWeightDevice, "blkio-weight-device", "Block IO weight (relative device weight)")
	flagSet.Var(&c.blkioDeviceReadBps, "device-read-bps", "Limit read rate (bytes per second) from a device")
	flagSet.Var(&c.blkioDeviceReadIOps, "device-read-iops", "Limit read rate (IO per second) from a device")
	flagSet.Var(&c.blkioDeviceWriteBps, "device-write-bps", "Limit write rate (bytes per second) from a device")
	flagSet.Var(&c.blkioDeviceWriteIOps, "device-write-iops", "Limit write rate (IO per second) from a device")

	// capbilities
	flagSet.StringSliceVar(&c.capAdd, "cap-add", nil, "Add Linux capabilities")
	flagSet.StringSliceVar(&c.capDrop, "cap-drop", nil, "Drop Linux capabilities")

	// cpu
	flagSet.Int64Var(&c.cpushare, "cpu-share", 0, "CPU shares (relative weight)")
	flagSet.StringVar(&c.cpusetcpus, "cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	flagSet.StringVar(&c.cpusetmems, "cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")

	flagSet.StringSliceVarP(&c.devices, "device", "", nil, "Add a host device to the container")
	flagSet.BoolVar(&c.enableLxcfs, "enableLxcfs", false, "Enable lxcfs for the container, only effective when enable-lxcfs switched on in Pouchd")
	flagSet.StringVar(&c.entrypoint, "entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	flagSet.StringSliceVarP(&c.env, "env", "e", nil, "Set environment variables for container")
	flagSet.StringVar(&c.hostname, "hostname", "", "Set container's hostname")
	flagSet.StringVar(&c.ipcMode, "ipc", "", "IPC namespace to use")
	flagSet.StringSliceVarP(&c.labels, "label", "l", nil, "Set labels for a container")

	// memory
	flagSet.Int64Var(&c.memorySwappiness, "memory-wappiness", -1, "Container memory swappiness [0, 100]")
	flagSet.StringVarP(&c.memory, "memory", "m", "", "Memory limit")
	flagSet.StringVar(&c.memorySwap, "memory-swap", "", "Swap limit equal to memory + swap, '-1' to enable unlimited swap")

	flagSet.StringVar(&c.name, "name", "", "Specify name of container")
	flagSet.StringSliceVar(&c.networks, "net", nil, "Set networks to container")
	flagSet.StringVar(&c.pidMode, "pid", "", "PID namespace to use")
	flagSet.StringVar(&c.restartPolicy, "restart", "", "Restart policy to apply when container exits")
	flagSet.StringVar(&c.runtime, "runtime", "", "OCI runtime to use for this container")
	flagSet.StringSliceVar(&c.securityOpt, "security-opt", nil, "Security Options")
	flagSet.StringSliceVar(&c.sysctls, "sysctl", nil, "Sysctl options")
	flagSet.BoolVarP(&c.tty, "tty", "t", false, "Allocate a pseudo-TTY")
	flagSet.StringVar(&c.utsMode, "uts", "", "UTS namespace to use")
	flagSet.StringSliceVarP(&c.volume, "volume", "v", nil, "Bind mount volumes to container")
	flagSet.StringVarP(&c.workdir, "workdir", "w", "", "Set the working directory in a container")

	return c
}
