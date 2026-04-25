package hardwarediscover

type System struct {
	Manufacturer string `json:"manufacturer"`
	ProductName  string `json:"product_name"`
	BIOSVersion  string `json:"bios_version"`
	BIOSDate     string `json:"bios_date"`
	MachineID    string `json:"machine_id"`
}

// TODO: /sys/devices/system/cpu/
type CPU struct {
	Vendor       string          `json:"vendor"`
	ModelName    string          `json:"model_name"`
	Architecture CPUArchitecture `json:"architecture"`
	Microcode    string          `json:"microcode"`

	Sockets        int `json:"sockets"`
	CoresPerSocket int `json:"cores_per_socket"`
	ThreadsPerCore int `json:"threads_per_core"`
	TotalThreads   int `json:"total_threads"`
	NumaNodes      int `json:"numa_nodes"`

	BaseFrequency float64 `json:"base_frequency"`
	BogoMIPS      float64 `json:"bogomips"`

	CoreMapping []LogicalThread `json:"core_mapping"`

	Cache CPUCache `json:"cache"`

	Flags    []string `json:"flags"`
	VMXFlags []string `json:"vmx_flags"`
	Bugs     []string `json:"bugs"`
}

type LogicalThread struct {
	ProcessorID int `json:"processor_id"`
	CoreID      int `json:"core_id"`
	SocketID    int `json:"socket_id"`
	NumaNode    int `json:"numa_node"`

	Capacity int         `json:"capacity"`
	CoreType CPUCoreType `json:"core_type"`
}

type CPUCache struct {
	L1dBytes       int `json:"l1d_bytes"`
	L1iBytes       int `json:"l1i_bytes"`
	L2Bytes        int `json:"l2_bytes"`
	L3Bytes        int `json:"l3_bytes"`
	CacheLineBytes int `json:"cache_line_bytes"`
}

// dmidecode -t memory
type Memory struct {
	TotalBytes  uint64         `json:"total_bytes"`
	Type        MemoryType     `json:"type"`
	SpeedMTs    int            `json:"speed_mt_s"`
	ECC         bool           `json:"ecc"`
	NumaMapping map[int]uint64 `json:"numa_mapping"`
	NumaBalance bool           `json:"numa_balance"`
	Modules     []DIMM         `json:"modules"`
}

type DIMM struct {
	Locator            string `json:"locator"`
	BankLocator        string `json:"bank_locator"`
	SizeBytes          uint64 `json:"size_bytes"`
	SpeedMTs           int    `json:"speed_mt_s"`
	ConfiguredSpeedMTs int    `json:"configured_speed_mt_s"`
	DataWidth          int    `json:"data_width"`
	Rank               int    `json:"rank"`
	Manufacturer       string `json:"manufacturer"`
	PartNumber         string `json:"part_number"`
}

// /sys/class/net/
type NIC struct {
	Device          string   `json:"device"`
	MACAddress      string   `json:"mac_address"`
	SpeedMbps       int      `json:"speed_mbps"`
	LinkUp          bool     `json:"link_up"`
	Driver          string   `json:"driver"`
	Modalias        string   `json:"modalias"`
	FirmwareVersion string   `json:"firmware_version"`
	PCIAddress      string   `json:"pci_address"`
	NumaNode        int      `json:"numa_node"`
	IOMMUGroup      int      `json:"iommu_group"`
	MaxVFs          int      `json:"max_vfs"`
	Features        []string `json:"features"`
	VendorID        string   `json:"vendor_id"`
	DeviceID        string   `json:"device_id"`
	SubVendorID     string   `json:"sub_vendor_id"`
	SubDeviceID     string   `json:"sub_device_id"`
	PCIClass        string   `json:"pci_class"`
}

// /sys/block/
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
