package hardwarediscover

type CPUArchitecture string
type CPUCoreType string
type MemoryType string
type DiskType string
type DiskBusType string

const (
	CPUArchitectureX86_64 CPUArchitecture = "x86_64"
	CPUArchitectureARM64  CPUArchitecture = "aarch64"
)

const (
	CPUCoreTypePerformance CPUCoreType = "performance"
	CPUCoreTypeEfficiency  CPUCoreType = "efficiency"
	CPUCoreTypeUnknown     CPUCoreType = "standard"
)

const (
	MemoryTypeDDR3 MemoryType = "DDR3"
	MemoryTypeDDR4 MemoryType = "DDR4"
	MemoryTypeDDR5 MemoryType = "DDR5"
)

const (
	DiskTypeHDD  DiskType = "HDD"
	DiskTypeSSD  DiskType = "SSD"
	DiskTypeNVMe DiskType = "NVMe"
)

const (
	DiskBusTypeSATA DiskBusType = "SATA"
	DiskBusTypeNVMe DiskBusType = "PCIe"
	DiskBusTypeSAS  DiskBusType = "SAS"
)
