http_benchmark:
	wrk -t1 -c50 -d20 http://0.0.0.0:8081/plaintext

godebug:
	GODEBUG=gctrace=1 ./no1_benchmark/http_slow/http_slow

godebug_cli:
	GODEBUG=gctrace=1 hey -z 10s http://127.0.0.1:8080/plaintext

pprof:
	open http://127.0.0.1:8080/debug/pprof

cpu:
	go tool pprof -http=:32001 http://127.0.0.1:8081/debug/pprof/profile?seconds=10

mem:
	go tool pprof -http=:33000 http://127.0.0.1:8081/debug/pprof/heap

trace:
	curl -o trace.out http://127.0.0.1:8081/debug/pprof/trace?seconds=2
	go tool trace trace.out