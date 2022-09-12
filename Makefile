SERVICE_NAME=dnetwork-backend
SERVICE_IMAGE=opas/dnetwork-backend
SERVICE_VERSION=1.0.0

export SERVICE_NAME
export SERVICE_IMAGE
export SERVICE_VERSION

hello:
	echo "Hello"

dev:
	@echo "---Start Dev $(SERVICE_NAME)---"
	go run cmd/api/main.go

build:
	GOOS=windows GOARCH=amd64 go build -o bin/dnetwork.exe cmd/v1/main.go

build-docker:
	docker build -t dnetwork-backend:1.0.0 . --build-arg VERSION=1.0.0

run:
	go run cmd/v1/main.go

# sonar:
# 	docker run --rm -e SONAR_HOST_URL="http://192.168.120.12:9000" -e SONAR_LOGIN="sqp_8bcc54bcca1728610ea0e2e15f3c211840261dfa" -v "/Users/opas/Projects/github.com/opas-keat/ect/dnetwork-phase2/backend:/usr/src" sonarsource/sonar-scanner-cli -Dsonar.sonar.projectName=dnetwork-backend

# git tag v1.0.1 
# git push --tags

# dev-up:
# 	@echo "---Start Dev $(SERVICE_NAME) Environtment---"
# 	docker-compose -p todo-api-dev -f ./deploy/docker-compose.yml -f ./deploy/docker-compose.dev.yml up -d

# dev-down:
# 	@echo "---Stop Dev $(SERVICE_NAME) Environtment---"
# 	docker-compose -p todo-api-dev -f ./deploy/docker-compose.yml -f ./deploy/docker-compose.dev.yml down

# dev:
# 	@echo "---Start Dev $(SERVICE_NAME)---"
# 	go run cmd/api/main.go

d-build:
	@echo "---Build $(SERVICE_NAME) $(SERVICE_IMAGE):$(SERVICE_VERSION)---"
	docker build -t $(SERVICE_IMAGE):$(SERVICE_VERSION) -f deploy/Dockerfile .

# d-build-debug:
# 	@echo "---Build $(SERVICE_NAME) $(SERVICE_IMAGE):$(SERVICE_VERSION)---"
# 	docker build --progress plain -t $(SERVICE_IMAGE):$(SERVICE_VERSION) -f deploy/Dockerfile .

# prod-up:
# 	@echo "---Start Prod $(SERVICE_NAME)---"
# 	docker-compose -p task-api-prod -f ./deploy/docker-compose.yml -f ./deploy/docker-compose.prod.yml up -d

# prod-down:
# 	@echo "---Stop Prod $(SERVICE_NAME)---"
# 	docker-compose -p task-api-prod -f ./deploy/docker-compose.yml -f ./deploy/docker-compose.prod.yml down