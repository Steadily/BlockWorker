package worker

import (
	"context"
	"fmt"
)

type forwardEvmTask struct {
	*evmpoller.ForwardEvmTask
}
