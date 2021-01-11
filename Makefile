build:
	docker image build -t nrece:v1 .
exec:
	docker exec -it nrece /bin/sh
clean:
	./scripts/clean.sh
run:
	docker run --rm -dit --name nrece -p 80:80 nrece:v1
stop:
	docker stop nrece

go-build:
	./scripts/build.sh
go-lint:
	./scripts/lint.sh
