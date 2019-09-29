run:
	test $(CMD)
	go run . $(CMD)

test:
	go test -cover -v
