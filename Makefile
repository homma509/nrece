.PHONY: build exec clean run stop go-lint

build:
	docker image build -t nrece:1.0.0 .
exec:
	docker exec -it nrece /bin/sh
clean:
	./scripts/clean.sh
run:
	docker run --rm -dit --name nrece -p 80:80 nrece:1.0.0
stop:
	docker stop nrece

go-build:
	./scripts/build.sh
go-lint:
	./scripts/lint.sh
