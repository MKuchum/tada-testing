<h3>Дополнительные комментарии</h3>

1) Дополнительно было определено что n > 0
2) Существует версия TribonacciStorageInMemory
3) ```
   flag.StringVar(&redisAddr, "redis_addr", "localhost:6379", "Redis address in format 'host:port', by default 'localhost:6739'")
   flag.StringVar(&redisPassword, "redis_password", "", "Redis password, by default empty")
   flag.IntVar(&redisDB, "redis_db", 0, "Redis database, by default 0 (redis default DB)")
   flag.StringVar(&serverAddr, "addr", "localhost:8080", "Server address in format 'host:port', by default 'localhost:8080'")
   ```
	
