hello:
	echo "Hello"

build:
	GOOS=windows GOARCH=amd64 go build -o bin/dnetwork.exe cmd/v1/main.go

build-docker:
	docker build -t dnetwork-backend:1.0.0 . --build-arg VERSION=1.0.0

run:
	go run cmd/v1/main.go

sonar:
	docker run --rm -e SONAR_HOST_URL="http://192.168.120.12:9000" -e SONAR_LOGIN="sqp_8bcc54bcca1728610ea0e2e15f3c211840261dfa" -v "/Users/opas/Projects/github.com/opas-keat/ect/dnetwork-phase2/backend:/usr/src" sonarsource/sonar-scanner-cli -Dsonar.sonar.projectName=dnetwork-backend

# git tag v1.0.1 
# git push --tags