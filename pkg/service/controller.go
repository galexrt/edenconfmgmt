/*
Copyright 2019 Alexander Trost. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package service

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/go-systemd/daemon"
	"go.uber.org/zap"
)

// ReloadableFunc
type ReloadableFunc func() error

// Controller keeps .
type Controller struct {
	stopCh         chan struct{}
	logger         *zap.Logger
	lastReload     time.Time
	reloadables    map[string]ReloadableFunc
	stopMiddleware []func() error
	stopChs        map[string]chan struct{}
}

// New return new Controller.
func New(logger *zap.Logger) *Controller {
	return &Controller{
		stopCh:         make(chan struct{}),
		logger:         logger,
		lastReload:     time.Time{},
		stopMiddleware: []func() error{},
		stopChs:        map[string]chan struct{}{},
	}
}

func (c *Controller) systemdWatchdog() {
	interval, err := daemon.SdWatchdogEnabled(false)
	if err != nil {
		c.logger.Warn("not starting systemd watchdog due to error", zap.Error(err))
		return
	}
	if interval <= 0 {
		c.logger.Debug("not starting systemd watchdog due to interval being zero or less", zap.Duration("interval", interval))
		return
	}
	for {
		select {
		case <-time.After(interval / 2):
			daemon.SdNotify(false, daemon.SdNotifyWatchdog)
		case <-c.stopCh:
			return
		}
	}
}

// GetStopCh return a `chan struct{}` which can be used to check if the application has been stopped.
func (c *Controller) GetStopCh(name string) chan struct{} {
	c.stopChs[name] = make(chan struct{})
	return c.stopChs[name]
}

// Start start the service controller logic.
// The systemd watchdog loop is started, if enabled/possible, as part of this.
func (c *Controller) Start(stopCh chan struct{}) {
	go c.systemdWatchdog()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGHUP)

	select {
	// Reload all given "reloadables" on SIGHUP signal
	case <-sigCh:
		names := []string{}
		for key := range c.reloadables {
			names = append(names, key)
		}
		c.Reload(names...)
		// Sleep 500 Milliseconds for everything to settle
		time.Sleep(500 * time.Millisecond)
	// On stopCh close, close all stopChs that have been given out
	case <-stopCh:
		daemon.SdNotify(false, daemon.SdNotifyStopping)
		for _, ch := range c.stopChs {
			close(ch)
		}
		// Tell systemd a second time that we are stopping
		daemon.SdNotify(false, daemon.SdNotifyStopping)
	}
}

// Reload reloads one or more services using given "reloadabales" function.
func (c *Controller) Reload(names ...string) {
	daemon.SdNotify(false, daemon.SdNotifyReloading)
	for name := range names {
		fmt.Printf("TEST: %+v\n", name)
	}
	c.lastReload = time.Now()
	daemon.SdNotify(false, daemon.SdNotifyReady)
}
