APP_BINARY=structApp
APP_IMAGE=struct-practise



build_app:
	@echo "Building struct-practise binary"
	env GOOS=linux CGO_ENABLED=0 go build -o ${APP_BINARY} .
	@echo "Done!"

up_build: build_app
	@echo "Building dockerfile"
	docker build -f structs-practice.dockerfile -t ${APP_IMAGE} .
	@echo "Starting docker container"
	docker run ${APP_IMAGE}
