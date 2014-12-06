package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	os.Exit(realMain(os.Args))
}

func realMain(args []string) int {
	var remove bool
	var pidFile string
	flag.StringVar(&pidFile, "pidfile", "", "pid file location")
	flag.BoolVar(&remove, "remove", false, "remove pidfile on exit")
	flag.Parse()

	if pidFile == "" {
		fmt.Println("Must provide -pidfile location")
		flag.Usage()
		return 1
	}

	cmdArgs := flag.Args()
	if len(cmdArgs) == 0 {
		fmt.Println("No command specified")
		flag.Usage()
		return 1
	}

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		return 1
	}

	fh, err := os.Create(pidFile)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	if remove {
		defer os.Remove(fh.Name())
	}
	if _, err := fh.WriteString(fmt.Sprintf("%d\n", cmd.Process.Pid)); err != nil {
		fmt.Println(err)
		return 1
	}
	fh.Close()

	if err := cmd.Wait(); err != nil {
		// Get the exit status. There isn't a platform-independent way to do
		// this so its a bit round-about.
		return err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus()
	}

	return 0
}
