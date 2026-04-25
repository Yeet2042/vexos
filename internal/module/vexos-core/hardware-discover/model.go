package hardwarediscover

type System struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	BIOSVersion  string `json:"bios_version"`
	BIOSDate     string `json:"bios_date"`
	MachineID    string `json:"machine_id"`
}

type CPU struct {
	Vendor         string          `json:"vendor"`
	Model          string          `json:"model"`
	Architecture   CPUArchitecture `json:"architecture"`
	Sockets        int             `json:"sockets"`
	CoresPerSocket int             `json:"cores_per_socket"`
	ThreadsPerCore int             `json:"threads_per_core"`
	TotalThreads   int             `json:"total_threads"`
	BaseFrequency  float64         `json:"base_frequency"`
	MaxFrequency   float64         `json:"max_frequency"`
	NumaNodes      int             `json:"numa_nodes"`
	L3CacheBytes   uint64          `json:"l3_cache_bytes"`
	Flags          []string        `json:"flags"`
	HasPMU         bool            `json:"has_pmu"`
	NumaCPUMap     map[int][]int   `json:"numa_cpu_map"`
}

type Memory struct {
	TotalBytes  uint64         `json:"total_bytes"`
	Type        MemoryType     `json:"type"`
	SpeedMTs    int            `json:"speed_mt_s"`
	ECC         bool           `json:"ecc"`
	Modules     []DIMM         `json:"modules"`
	NumaMapping map[int]uint64 `json:"numa_mapping"`
}

type DIMM struct {
	Locator      string `json:"locator"`
	SizeBytes    uint64 `json:"size_bytes"`
	SpeedMTs     int    `json:"speed_mt_s"`
	DataWidth    int    `json:"data_width"`
	Manufacturer string `json:"manufacturer"`
	PartNumber   string `json:"part_number"`
}

type NIC struct {
	Device          string   `json:"device"`
	MACAddress      string   `json:"mac_address"`
	SpeedMbps       int      `json:"speed_mbps"`
	LinkUp          bool     `json:"link_up"`
	Driver          string   `json:"driver"`
	FirmwareVersion string   `json:"firmware_version"`
	PCIAddress      string   `json:"pci_address"`
	NumaNode        int      `json:"numa_node"`
	IOMMUGroup      int      `json:"iommu_group"`
	MaxVFs          int      `json:"max_vfs"`
	Features        []string `json:"features"`
}

type Disk struct {
	Device           string      `json:"device"`
	Model            string      `json:"model"`
	Type             string      `json:"type"`
	IsRotational     bool        `json:"is_rotational"`
	CapacityBytes    uint64      `json:"capacity_bytes"`
	LogicalBlockSize int         `json:"logical_block_size"`
	Serial           string      `json:"serial"`
	WorldWideName    string      `json:"world_wide_name"`
	BusType          DiskBusType `json:"bus_type"`
	PCIAddress       string      `json:"pci_address"`
	NumaNode         int         `json:"numa_node"`
	SmartHealthOK    bool        `json:"smart_health_ok"`
}

type Accelerator struct {
	Device      string `json:"device"`
	Vendor      string `json:"vendor"`
	Model       string `json:"model"`
	MemoryBytes uint64 `json:"memory_bytes"`
	PCIAddress  string `json:"pci_address"`
	NumaNode    int    `json:"numa_node"`
	IOMMUGroup  int    `json:"iommu_group"`
}
