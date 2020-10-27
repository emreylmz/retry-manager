# Retry Manager
A golang package that helps you run a function with repeatable .

It behaves in a similar way to Javascript's `setInterval` function.


##Usage

````go
	//initialize retry manager your main go send as a parameter for singleton usage.
	errors := make(chan redisCache.RetryHandler)
	retryManager := redisCache.NewRetryManager(errors, 5 * time.Second,5,logger);


    //use retry-manager your error cases etc.
    retry := manager.RetryHandler {
        Execute: func() error {
            return errors.New("sadasdsa") // service.SetTicketToCache(request);
        },
        RetryErrorLog: fmt.Sprintf("Key : %s",key),
    }
    service.retryManager.AddHandler(retry)
````