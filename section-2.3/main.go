package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	darwin  = "darwin"
	linux   = "linux"
	windows = "windows"
)

func main() {
	var out *Platform
	var err error
	out, err = getInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(out.String())
}

func getInfo() (*Platform, error) {
	var out bytes.Buffer
	var stderr bytes.Buffer
	var cmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case darwin:
		cmd = exec.Command("uname", "-srm")
	case linux:
		cmd = exec.Command("uname", "-srio")
	case windows:
		cmd = exec.Command("cmd", "ver")
	default:
		log.Fatal("Unsupported OS.")
	}

	cmd.Stdin = strings.NewReader("some input")
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()

	osStr := strings.Replace(out.String(), "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osInfo := strings.Split(osStr, " ")

	if err != nil {
		return nil, err
	}

	p := &Platform{
		Kernel:   osInfo[0],
		Core:     osInfo[1],
		Platform: osInfo[2],
		OS:       osInfo[0],
		GoOS:     runtime.GOOS,
		CPUs:     runtime.NumCPU(),
	}
	p.Hostname, _ = os.Hostname()
	return p, nil
}

// Platform describes host system
type Platform struct {
	GoOS     string
	Kernel   string
	Core     string
	Platform string
	OS       string
	Hostname string
	CPUs     int
}

// String returns the platform details
func (p *Platform) String() string {
	return fmt.Sprintf("GoOS: %v, Kernel: %v, Core: %v, Platform: %v, OS: %v, Hostname: %v, CPUs: %v",
		p.GoOS, p.Kernel, p.Core, p.Platform, p.OS, p.Hostname, p.CPUs)
}
