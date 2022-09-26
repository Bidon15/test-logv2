package main

import (
	"context"
	"time"

	logging "github.com/ipfs/go-log"
	"github.com/testground/sdk-go/run"
	"github.com/testground/sdk-go/runtime"
	"github.com/testground/sdk-go/sync"
)

func main() {
	run.InvokeMap(map[string]interface{}{
		"smallbrain": run.InitializedTestCaseFn(test),
	})
}

func test(runenv *runtime.RunEnv, initCtx *run.InitContext) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	syncclient := initCtx.SyncClient
	state := sync.State("ready")
	seq := syncclient.MustSignalAndWait(ctx, state, runenv.TestInstanceCount)
	var log = logging.Logger("test")
	lvl, err := logging.LevelFromString("info")
	if err != nil {
		return err
	}
	logging.SetAllLoggers(lvl)
	time.Sleep(1 * time.Second)
	runenv.RecordMessage("I am seq %d", seq)
	log.Infof("I am seq %d", seq)

	runenv.RecordSuccess()

	return nil
}
