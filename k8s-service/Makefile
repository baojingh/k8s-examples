SHELL := /bin/bash
VERSION := 1.0

all: service 

run:
	go run main.go

service:
	docker build \
		-f zarf/docker/dockerfile \
		-t service-amd64:$(VERSION) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
		.



# ==============================================================================
# Running from within k8s/kind
KIND_CLUSTER := ardan-starter-cluster
 
kind-up:
	kind create cluster  \
			--name ${KIND_CLUSTER} \
			--config zarf/k8s/kind/kind-config.yaml
	kubectl config set-context --current \
					--namespace=service-system

kind-down:
	kind delete cluster \
			--name ${KIND_CLUSTER}


kind-status:
	kubectl get deployment 	-owide
	kubectl get svc 		-owide
	kubectl get nodes 		-owide
	kubectl get pod 		-owide

# the docker iamge in host cannot be found by k8s(kind), load is must
kind-load:
	kind load docker-image service-amd64:${VERSION} \
				--name ${KIND_CLUSTER}

kind-apply:
	cat zarf/k8s/base/service-pod/base-service.yaml | kubectl apply -f -

kind-logs:
	kubectl logs -l app=service --all-containers=true \
			-f --tail=100

kind-restart:
	kubectl rollout restart deployment service-pod
