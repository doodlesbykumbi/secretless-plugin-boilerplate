build_example:
	go build -buildmode=plugin -o ./plugins/example.so plugin_example.go
run_secretless: build_example
	secretless-broker -f ./secretless.yml -p ./plugins -debug
client_without_secretless:
	curl -v httpbin.org/anything
client_with_secretless:
	http_proxy=http://localhost:6175 curl -v httpbin.org/anything
