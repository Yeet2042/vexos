package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/Yeet2042/vexos/config/vexos-core"
	vexosservice "github.com/Yeet2042/vexos/internal/core/vexos-core"
	cfg "github.com/Yeet2042/vexos/pkg/cfg"
)

func main() {
	// context with signal handling for graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// ----- Initialize Config
	config, err := cfg.NewConfig[config.VEXOSConfig]("config/vexos-core/config.yml")
	if err != nil {
		log.Fatalf("[main]: Failed to load config: %v", err)
	}

	// ----- Initialize Package

	// ----- Initialize Module

	// ----- Initialize Core Service
	service, err := vexosservice.NewV1(
		config,
	)
	if err != nil {
		log.Fatalf("[main]: Failed to initialize Core Service: %v", err)
	}

	// ----- Start Core Service
	log.Println("[main]: Starting Core Service...")
	err = service.Start(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Fatalf("[main]: Failed to start Core Service: %v", err)
	}

	log.Println("[main]: Graceful shutdown complete.")
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
