// +build solaris

package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"syscall"

	"github.com/docker/docker/libcontainerd"
	"github.com/docker/docker/pkg/system"
)

const defaultDaemonConfigFile = ""

// currentUserIsOwner checks whether the current user is the owner of the given
// file.
func currentUserIsOwner(f string) bool {
	if fileInfo, err := system.Stat(f); err == nil && fileInfo != nil {
		if int(fileInfo.UID()) == os.Getuid() {
			return true
		}
	}
	return false
}

// setDefaultUmask sets the umask to 0022 to avoid problems
// caused by custom umask
func setDefaultUmask() error {
	desiredUmask := 0022
	syscall.Umask(desiredUmask)
	if umask := syscall.Umask(desiredUmask); umask != desiredUmask {
		return fmt.Errorf("failed to set umask: expected %#o, got %#o", desiredUmask, umask)
	}

	return nil
}

func getDaemonConfDir() string {
	return "/etc/docker"
}

// setupConfigReloadTrap configures the USR2 signal to reload the configuration.
func (cli *DaemonCli) setupConfigReloadTrap() {
}

// notifySystem sends a message to the host when the server is ready to be used
func notifySystem() {
}

func (cli *DaemonCli) getPlatformRemoteOptions() []libcontainerd.RemoteOption {
	opts := []libcontainerd.RemoteOption{}
	return opts
}

// getLibcontainerdRoot gets the root directory for libcontainerd/containerd to
// store their state.
func (cli *DaemonCli) getLibcontainerdRoot() string {
	return filepath.Join(cli.Config.ExecRoot, "libcontainerd")
}

func allocateDaemonPort(addr string) error {
	return nil
}

// notifyShutdown is called after the daemon shuts down but before the process exits.
func notifyShutdown(err error) {
}

func wrapListeners(proto string, ls []net.Listener) []net.Listener {
	return ls
}
