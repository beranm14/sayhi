build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o hello main.go
	docker build . -t 515719629808.dkr.ecr.eu-west-1.amazonaws.com/beranm:job
	docker push 515719629808.dkr.ecr.eu-west-1.amazonaws.com/beranm:job