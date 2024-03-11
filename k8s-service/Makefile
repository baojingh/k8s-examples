SHELL := /bin/bash
VERSION := 1.0
NAMESPACE := sales-system

all: service 

run:
	go run app/services/sales-api/main.go

sales:
	# docker rmi $(docker images | grep none | awk '{print $3}')
	
	docker build \
		-f zarf/docker/dockerfile.sales-api \
		-t sales-api-amd64:$(VERSION) \
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
					--namespace=${NAMESPACE}

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
	kind load docker-image sales-api-amd64:${VERSION} \
				--name ${KIND_CLUSTER}

kind-apply:
	kustomize build zarf/k8s/kind/sales-pod/ | kubectl apply -f -

kind-logs:
	kubectl logs -l app=sales --all-containers=true -f --tail=100

kind-restart:
	kubectl rollout restart deployment sales-pod

kind-des:
	kubectl describe pod -l app=sales


kind-update: all kind-load kind-restart

kind-update-apply: all kind-load kind-apply


# ==============================================================================

tidy:
	go mod tidy
