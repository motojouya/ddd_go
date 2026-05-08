package pkg

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	databaseRepository "github.com/motojouya/ddd_go/pkg/database/repository"
	localRepository "github.com/motojouya/ddd_go/pkg/local/repository"
	qRepository "github.com/motojouya/ddd_go/pkg/queue/repository"
	qController "github.com/motojouya/ddd_go/pkg/queue/controller"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	qStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type WorkCmd struct {
	WorkerName  string `arg:"" name:"worker-name" help:"worker name"`
	KeepWorking bool   `name:"keep-working" default:"true" negatable:"" help:"if true, work as daemon (default). use --no-keep-working for one-shot drain"`
}

// FIXME gopsutilで同じworkerを複数process起動できるか制御できる https://mikoto2000.blogspot.com/2024/05/go-gopsutil.html
func (wrk *WorkCmd) Run() error {
	route := queueCore.NewJobRouter()
	err := RegisterProcedure(route)
	if err != nil {
		log.Fatal(err)
	}

	// db connectionを取得してserver停止時にclose
	dbGetter := databaseRepository.NewDatabaseGet()
	db, err := dbGetter.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	localer := localRepository.CreateLocal()
	qStr := qStore.NewQueueStore(db)
	qBhv := qRepository.NewQueueRepository(qStr, localer)

	var wg sync.WaitGroup
	// SIGINT/SIGTERMで停止する
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// FIXME SIGTERMしても長ければ終わらないのでtimeout設定があったほうがいいが、そのときはSIGKILLでいい気もする
	wg.Add(1)
	go func() {
		fmt.Println("worker start!")
		err := qController.ExecuteWorker(qBhv, route, ctx, wrk.WorkerName, wrk.KeepWorking)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
	return nil
}

func RegisterProcedure(route queueCore.JobRouter) error {
	// jobを追加していく
	return nil
}
