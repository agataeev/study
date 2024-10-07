package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
	"unpack-text/internal/commands"
)

const (
	UnpackCommand = "unpack"
	PackCommand   = "pack"
)

func runAsDaemon() {
	execPath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	cmd := exec.Command(execPath, "--mode", "daemon")
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start as daemon: %v", err)
	}

	log.Printf("Daemon started with PID: %d", cmd.Process.Pid)
	os.Exit(0)
}

func daemonUnpackMode() {
	fmt.Println("Введите строку для распаковки (Ctrl^C для завершения):")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Введите строку: ")
		if scanner.Scan() {
			input := scanner.Text()

			// Unpack the string
			result, err := commands.Unpack(input)
			if err != nil {
				fmt.Printf("Ошибка распаковки: %v\n", err)
			} else {
				fmt.Printf("Распаковано: %s\n", result)
			}
		}
	}
}

func main() {
	mode := flag.String("mode", "", "Mode: unpack or pack or daemon")
	input := flag.String("input", "", "Input text to pack or unpack")
	daemon := flag.Bool("daemon", false, "Run in daemon mode")
	flag.Parse()

	if *daemon {
		runAsDaemon()
		return
	}

	if *mode == "daemon" && *mode == UnpackCommand {
		daemonUnpackMode()
		return
	}

	// Packing and unpacking logic
	switch *mode {
	case UnpackCommand:
		unpacked, err := commands.Unpack(*input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Unpacked: %s\n", unpacked)
		}
	case PackCommand:
		packed := commands.Pack(*input)
		fmt.Printf("Packed: %s\n", packed)
	default:
		fmt.Println("Invalid mode. Use 'pack', 'unpack', or 'daemon'.")
	}
}
