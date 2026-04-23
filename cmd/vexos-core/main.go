package main

func main() {

}

// eBPF load example
// func main() {
// 	spec, err := ebpf.LoadCollectionSpec("tools/vexos-ebpf/target/bpfel-unknown-none/release/libvexos_ebpf.so")
// 	if err != nil {
// 		log.Fatalf("Failed to load ELF: %v", err)
// 	}

// 	for name := range spec.Programs {
// 		log.Printf("Found eBPF Program in ELF: '%s'", name)
// 	}

// 	var objs struct {
// 		XdpFirewall *ebpf.Program `ebpf:"xdp_firewall"`
// 	}
// 	if err := spec.LoadAndAssign(&objs, nil); err != nil {
// 		log.Fatalf("Failed to prepare eBPF objects: %v", err)
// 	}
// 	defer objs.XdpFirewall.Close()

// 	ifaceName := "lo"
// 	iface, err := net.InterfaceByName(ifaceName)
// 	if err != nil {
// 		log.Fatalf("Failed to find interface %s: %v", ifaceName, err)
// 	}

// 	l, err := link.AttachXDP(link.XDPOptions{
// 		Program:   objs.XdpFirewall,
// 		Interface: iface.Index,
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to attach XDP: %v", err)
// 	}
// 	defer l.Close()

// 	log.Printf("Successfully attached XDP program to interface %s", ifaceName)

// 	stopper := make(chan os.Signal, 1)
// 	signal.Notify(stopper, os.Interrupt, syscall.SIGTERM)
// 	<-stopper
// }
