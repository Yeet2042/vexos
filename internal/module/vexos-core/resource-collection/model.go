package resourcecollection

type PSIMetrics struct {
	Average10S float64 `json:"average_10_s"`
	Average1M  float64 `json:"average_1_m"`
	Average5M  float64 `json:"average_5_m"`
	TotalUs    uint64  `json:"total_us"`
}

type HostPSI struct {
	CPU    *PSIMetrics `json:"cpu"`
	Memory *PSIMetrics `json:"memory"`
	IO     *PSIMetrics `json:"io"`
}

type CPUHostMetrics struct {
	UserHz         uint64  `json:"user_hz"`
	UsagePercent   float64 `json:"usage_percent"`
	IdlePercent    float64 `json:"idle_percent"`
	IOWaitPercent  float64 `json:"io_wait_percent"`
	ContextSwitch  uint64  `json:"context_switch"`
	Interrupts     uint64  `json:"interrupts"`
	LoadAverage1M  float64 `json:"load_average_1_m"`
	LoadAverage5M  float64 `json:"load_average_5_m"`
	LoadAverage15M float64 `json:"load_average_15_m"`
}

type CPUGuestMetrics struct {
	VMID              string  `json:"vmid"`
	PID               int     `json:"pid"`
	VCPUs             int     `json:"vcpus"`
	UsagePercent      float64 `json:"usage_percent"`
	RunqueueLatencyUs uint64  `json:"runqueue_latency_us"`
	StealTimeUs       uint64  `json:"steal_time_us"`
	IPC               float64 `json:"ipc"`
	CacheMisses       uint64  `json:"cache_misses"`
}

type MemoryHostMetrics struct {
	TotalBytes         uint64 `json:"total_bytes"`
	AvailableBytes     uint64 `json:"available_bytes"`
	UsedBytes          uint64 `json:"used_bytes"`
	HugePageSizeBytes  uint64 `json:"huge_page_size_bytes"`
	HugePageTotal      uint64 `json:"huge_page_total"`
	HugePageFree       uint64 `json:"huge_page_free"`
	HugePageReserved   uint64 `json:"huge_page_reserved"`
	BalloonActualBytes uint64 `json:"balloon_actual_bytes"`
	SwapUsedBytes      uint64 `json:"swap_used_bytes"`
	OOMKillCount       uint64 `json:"oom_killed_count"`
}

type MemoryGuestMetrics struct {
	VMID                    string `json:"vmid"`
	AllocatedBytes          uint64 `json:"allocated_bytes"`
	UsedBytes               uint64 `json:"used_bytes"`
	AllocationBytePerSecond uint64 `json:"allocation_byte_per_second"`
	DirtyPagePerSecond      uint64 `json:"dirty_page_per_second"`
	WorkingSet              uint64 `json:"working_set"`
}

type NetworkMetrics struct {
	Interface string `json:"interface"`
	RxBytes   uint64 `json:"rx_bytes"`
	TxBytes   uint64 `json:"tx_bytes"`
	RxPackets uint64 `json:"rx_packets"`
	TxPackets uint64 `json:"tx_packets"`
	RxErrors  uint64 `json:"rx_errors"`
	TxErrors  uint64 `json:"tx_errors"`
	RxDropped uint64 `json:"rx_dropped"`
	TxDropped uint64 `json:"tx_dropped"`
}

type HighThroughputNetworkMetrics struct {
	NetworkMetrics     NetworkMetrics `json:"network_metrics"`
	ReadLatencyUs      uint64         `json:"read_latency_us"`
	WriteLatencyUs     uint64         `json:"write_latency_us"`
	WriteThroughputBps uint64         `json:"write_throughput_bps"`
}

type DiskMetrics struct {
	Device      string `json:"device"`
	ReadBytes   uint64 `json:"read_bytes"`
	WriteBytes  uint64 `json:"write_bytes"`
	ReadIOPS    uint64 `json:"read_iops"`
	WriteIOPS   uint64 `json:"write_iops"`
	ReadErrors  uint64 `json:"read_errors"`
	WriteErrors uint64 `json:"write_errors"`
	AwaitTimeUs uint64 `json:"await_time_us"`
}

type DRBDVolumeMetrics struct {
	ResourceName       string              `json:"resource_name"`
	Role               DRBDRole            `json:"role"`
	Suspended          bool                `json:"suspended"`
	DiskState          string              `json:"disk_state"`
	ConnectionState    DRBDConnectionState `json:"connection_state"`
	OutOfSyncBytes     uint64              `json:"out_of_sync_bytes"`
	SyncPercent        float64             `json:"sync_percent"`
	PendingBytes       uint64              `json:"pending_bytes"`
	UnackedBytes       uint64              `json:"unacked_bytes"`
	ActivityLogExtents uint64              `json:"activity_log_extents"`
}
