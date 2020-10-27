package retryManager

import (
	"log"
	"time"
)

type RetryManager interface {
	AddHandler(handler RetryHandler)
}

type RetryManagerImpl struct {
	RetriesChannel     chan RetryHandler
	RetryTimeoutSecond time.Duration
	MaxRetryCount      int
	logger             log.Logger
}

type RetryHandler struct {
	Execute       func() error
	RetryErrorLog string
	retryCount    int
}

func (r *RetryManagerImpl) AddHandler(handler RetryHandler) {
	r.RetriesChannel <- handler
}

func (r *RetryManagerImpl) initialize(channel chan RetryHandler) {
	go func() {
		for {
			payload := <-channel
			errResponse := payload.Execute()
			if errResponse != nil {
				r.retry(payload)
			}
		}
	}()
}

func (r *RetryManagerImpl) retry(payload RetryHandler) {
	if payload.retryCount <= r.MaxRetryCount {
		time.AfterFunc(r.RetryTimeoutSecond, func() {
			payload.retryCount += 1
			err := payload.Execute()
			if err != nil {
				r.retry(payload)
			}
		})
	} else {
		r.logger.Print(payload.RetryErrorLog)
	}
}

func NewRetryManager(channel chan RetryHandler, retryTimeoutSecond time.Duration, maxRetryCount int, logger log.Logger) RetryManager {
	retryManager := &RetryManagerImpl{
		RetriesChannel:     channel,
		RetryTimeoutSecond: retryTimeoutSecond,
		MaxRetryCount:      maxRetryCount,
		logger:             logger,
	}
	retryManager.initialize(channel)
	return retryManager
}