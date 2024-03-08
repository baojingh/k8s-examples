SHELL := /bin/bash
VERSION := 1.0

run:
	go run main.go

service:
	docker build \
		-f zarf/docker/dockerfile \
		-t salse-api-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.


dev-status:
	kubectl get pod 		-owide
	kubectl get deployment 	-owide
	kubectl get svc 		-owide

# ==============================================================================
# Running from within k8s/kind

KIND_CLUSTER := ardan-starter-cluster

kind-up:
	kind create cluster --name ${KIND_CLUSTER}
	kubectl config set-context --current --namespace=sales-system

kind-down:
	kind delete cluster --name ${KIND_CLUSTER}
