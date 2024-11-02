package main

import (
	"log"
	"os"
	"os/signal"
	cache "storage-service/internal/cache/redis-adapter"
	database "storage-service/internal/database/postgres-adapter"
	queue "storage-service/internal/queue/nats-adapter"
	"storage-service/internal/service"
	"syscall"
)

func main() {
	log.Println("storage service")

	queueAdapter, err := queue.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := queueAdapter.Close(); err != nil {
			log.Println(err)
		}
	}()

	databaseAdapter, err := database.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := databaseAdapter.Close(); err != nil {
			log.Println(err)
		}
	}()

	cacheAdapter, err := cache.New(nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := cacheAdapter.Close(); err != nil {
			log.Println(err)
		}
	}()

	serviceInstance, err := service.New(nil, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	if err := serviceInstance.Start(); err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := serviceInstance.Shutdown(); err != nil {
			log.Println(err)
		}
	}()

	waitForInterrupt()
}

// waitForInterrupt ...
func waitForInterrupt() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
}
