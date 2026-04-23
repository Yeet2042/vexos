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
	fiberserver "github.com/Yeet2042/vexos/pkg/fiber-server"
)

func main() {
	// context with signal handling for graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// ----- Initialize Config
	config, err := cfg.NewConfig[config.VEXOSConfig]("config/vexos-core/config.yml")
	if err != nil {
		log.Printf("[main]: Failed to load config: %v", err)
		return
	}

	// ----- Initialize Package
	// database, err := database.NewSurrealdb(ctx, &database.SurrealdbConfig{
	// 	Path:      config.Database.Path,
	// 	Namespace: config.Database.Namespace,
	// 	Database:  config.Database.Database,
	// })
	// if err != nil {
	// 	log.Printf("[cmd/apss_main]: failed to connect to database: %v", err)
	// 	return
	// }
	// defer database.Close()

	fiber, err := fiberserver.NewFiber(&fiberserver.FiberConfig{
		Port: config.Server.Port,
	})
	if err != nil {
		log.Printf("[main]: Failed to initialize Fiber server: %v", err)
		return
	}
	defer fiber.Close()

	// ----- Initialize Module

	// ----- Initialize Core Service
	service, err := vexosservice.NewV1(
		config,
		fiber,
	)
	if err != nil {
		log.Printf("[main]: Failed to initialize Core Service: %v", err)
		return
	}

	// ----- Start Core Service
	log.Println("[main]: Starting Core Service...")
	err = service.Start(ctx)
	if err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("[main]: Failed to start Core Service: %v", err)
		return
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
