package main

import (
	"bufio"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

var (
	local = flag.Bool("local", false, "Run locally")
)

func main() {
	flag.Parse()
	cmdline := "/usr/share/bcc/tools/execsnoop" // test, alpine
	if *local {
		cmdline = "execsnoop-bpfcc" // my local (Pop!_OS)
	}

	// cmd := exec.Command("oomkill-bpfcc")
	cmd := exec.Command(cmdline)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	outpipe, err := cmd.StdoutPipe()
	if err != nil {
		slog.Error("StdoutPipe failed:", "err", err)
		return
	}

	errpipe, err := cmd.StderrPipe()
	if err != nil {
		slog.Error("StdoutPipe failed:", "err", err)
		return
	}

	var wg sync.WaitGroup
	err = cmd.Start()
	if err != nil {
		slog.Error("Start failed:", "err", err)
		return
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		s := make(chan os.Signal)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		sig := fmt.Sprintf("%s", <-s)
		slog.Info("terminated, cleaning up...", "sig", sig)

		err := cmd.Process.Signal(syscall.SIGTERM)
		if err != nil {
			slog.Info("failed to terminate process, force kill...")
			cmd.Process.Signal(syscall.SIGKILL)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		outscan := bufio.NewScanner(outpipe)
		for {
			chk := outscan.Scan()
			if !chk {
				break
			}

			stxt := outscan.Text()
			slog.Info("stdout:", "line", stxt)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		errscan := bufio.NewScanner(errpipe)
		for {
			chk := errscan.Scan()
			if !chk {
				break
			}

			stxt := errscan.Text()
			slog.Info("stderr:", "line", stxt)
		}
	}()

	cmd.Wait()
	wg.Wait()
}
