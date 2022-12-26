SHELL := /bin/bash

build_director:
	pushd talapas-api-director; docker build . -t talapas-api-director; popd

build_app1:
	pushd talapas-api-app1; docker build . -t talapas-api-app1; popd

build_app2:
	pushd talapas-api-app2; docker build . -t talapas-api-app2; popd

build: build_director build_app1 build_app2

MY_IP := $(shell ifconfig | grep inet | grep -v 127.0.0.1 | grep -v inet6 | awk '{print $$2}')

director: build_director
	docker run -it \
		-e TALAPAS_API_DIRECTOR_LISTEN_PORT=8080 \
		-e TALAPAS_API_DIRECTOR_APP1_HOST=$(MY_IP) \
		-e TALAPAS_API_DIRECTOR_APP1_PORT=8681 \
		-e TALAPAS_API_DIRECTOR_APP2_HOST=$(MY_IP) \
		-e TALAPAS_API_DIRECTOR_APP2_PORT=8682 \
		-p 8080:8080 \
		talapas-api-director:latest

app1: build_app1
	docker run -it \
		-e TALAPAS_API_APP1_PORT=8681 \
		-p 8681:8681 \
		talapas-api-app1:latest

app2: build_app2
	docker run -it \
		-e TALAPAS_API_APP2_PORT=8682 \
		-p 8682:8682 \
		talapas-api-app2:latest
