.PHONY: test


test:
	docker-compose run elit go test -cover ./...
