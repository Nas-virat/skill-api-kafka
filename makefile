
IMAGE_NAME_API = api
IMAGE_NAME_CONSUMER = consumer
IMAGE_TAG ?= latest
REGISTRY = registry.gitlab.com/nas-virat-arise/skill-api-kafka
USERNAME = nas-virat-arise

.PHONY: e2e
e2e:
	@echo "Running e2e tests..."
	./scripts/e2e-script.sh

.PHONY: push
push: build 
	docker push $(REGISTRY)/$(IMAGE_NAME_API):$(IMAGE_TAG)
	docker push $(REGISTRY)/$(IMAGE_NAME_CONSUMER):$(IMAGE_TAG)

.PHONY: login
login:
	docker login registry.gitlab.com -u $(USERNAME) -p $(GITLAB_TOKEN)

.PHONY: build
build: api-build consumer-build
	

.PHONY: api-build
api-build:
	@echo "Building api"
	docker build -t $(REGISTRY)/$(IMAGE_NAME_API):$(IMAGE_TAG) ./api


.PHONY: consumer-build
consumer-build:
	@echo "Building consumer"
	docker build -t $(REGISTRY)/$(IMAGE_NAME_CONSUMER):$(IMAGE_TAG) ./consumer	

