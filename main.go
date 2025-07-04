package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	BuildVersion   = ""
	BuildGoVersion = ""
	BuildTime      = ""
)

func init() {
	if BuildTime == "" {
		BuildTime = time.Now().String()
	}
	if BuildVersion == "" {
		BuildVersion = "just_build_for_develop"
	}
	fmt.Printf("BuildTime: %s\nBuildGoVersion: %s\nBuildGoVersion: %s\n", BuildTime, BuildGoVersion, BuildVersion)
}
