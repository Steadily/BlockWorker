package ethclient

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/prometheus/client_golang/prometheus"
)

type Route int

const (
	RouteRandom Route = -1
)

var (
	RequestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "web3",
		Subsystem: "ethclient",
		Name:      "request_duration",
		Help:      "Request duration",
	}, []string{
		"method",
		"endpoint",
	})
)

type Error struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ErrorMessage string `json:"Error"`
}
