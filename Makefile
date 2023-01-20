all: build

build:
	docker build -t mheers/dex-local-discovery:latest .

push:
	docker push mheers/dex-local-discovery:latest
