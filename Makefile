all: clean

clean:
	@rm ./bin/app_linux || true

build:
	GOOS=linux go build -o ./bin/app_linux ./cmd/main.go

dockerize:
	GOOS=linux go build -o ./bin/app_linux ./cmd/main.go
	docker stop minicube_tutor && docker rm minicube_tutor || true
	docker build -t local/minicube_tutor .
	docker run -d --name minicube_tutor -p 8888:8888 local/minicube_tutor
