package main

import (
	"context"
	heimdallr "git.llsapp.com/long.qin01/go-heimdallr"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"git.llsapp.com/platform/pkg/datacollect"
	"github.com/gin-gonic/gin"
)

type Test struct {
}

func (t *Test) GetConsumeFunc() heimdallr.ConsumeFunc{
	return func(records []heimdallr.Record) error {
		log.Println("consume batch records")
		_, err := http.Get("https://www.baidu.com")
		return err
	}
}

func (t *Test) GetProduceFunc() heimdallr.ProduceFunc{
	return func(i *gin.Context) heimdallr.Record {
		return map[string]interface{}{
			"status": i.Writer.Status(),
		}
	}
}

func (t *Test) GetRecordThresh() heimdallr.Thresh {
	return heimdallr.Thresh{
		Queue:    50,
		CronTime: 30,
	}
}

func main() {
	datacollect.BaseURL = "https://alidev-flm.thellsapi.com/data_collect"

	r := gin.Default()
	stub := heimdallr.NewStub()
	//heimdallrConfig := heimdallr.Config{
	//	AppName:       "rosa",
	//	EventType:     11,
	//	CommonData:    map[string]string{ "category": "test" },
	//	RecordThresh:  heimdallr.Thresh{ Queue: 5, CronTime: 5, Payload: 5 * 1024 },
	//	Ignore:        nil,
	//	CustomPayload: nil,
	//}
	r.Use(heimdallr.GinMiddleware(stub, &Test{}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	stub.Close(ctx)
}