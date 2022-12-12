package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"good/helpers"

	"github.com/spf13/cobra"
	"golang.org/x/sync/semaphore"
)

// @see https://gist.github.com/picatz/9c0028efd7b3ced3329f7322f41b16e1#file-port_scanner-go
// this approach is amazing

const TOP = 65535
const OUT = "256" // hardcoded for convenience (`ulimit -n` might be significantly higher)

var portScanCmd *cobra.Command

type PortScanner struct {
	ip   string
	lock *semaphore.Weighted
}

func sanitizeHost(host string) string {
	if helpers.IsIp(host) || host == "localhost" {
		return host
	}

	return ""
}

func runPortScanCmd(cmd *cobra.Command, args []string) error {
	ip, _ := cmd.Flags().GetString("ip")
	host := sanitizeHost(ip)

	if host == "" {
		return errors.New("you provided an incorrect IP")
	}

	ps := &PortScanner{
		ip:   host,
		lock: semaphore.NewWeighted(ulimit()),
	}
	ps.start(1, TOP, 500*time.Millisecond)
	return nil
}

func ulimit() int64 {
	i, err := strconv.ParseInt(OUT, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func scanPort(ip string, port int, timeout time.Duration) int {
	target := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.DialTimeout("tcp", target, timeout)

	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(timeout)
			scanPort(ip, port, timeout)
		}
		return 0
	}

	conn.Close()
	fmt.Printf("%d open\n\n", port)
	return port
}

func (ps *PortScanner) start(f, l int, timeout time.Duration) {
	var i int
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for port := f; port <= l; port++ {
		ps.lock.Acquire(context.TODO(), 1)
		wg.Add(1)
		go func(port int) {
			defer ps.lock.Release(1)
			defer wg.Done()
			p := scanPort(ps.ip, port, timeout)
			if p != 0 {
				i++
			}
		}(port)
	}
	if i == 0 {
		fmt.Print("There's no open port\n\n")
	}
}

func init() {
	portScanCmd = &cobra.Command{
		Use:   "ports --ip=[IP]",
		Short: "Enumerate open ports",
		RunE:  runPortScanCmd,
	}
	hackCmd.AddCommand(portScanCmd)
	portScanCmd.PersistentFlags().String("ip", "", "The IP to scan ('localhost' is accepted)")
	portScanCmd.MarkPersistentFlagRequired("ip")
}
