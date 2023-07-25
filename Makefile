stop:
	docker-compose -f src/docker-compose.yaml stop

start:
	docker-compose -f src/docker-compose.yaml up  -d --no-deps --build 

test:
	docker-compose -f src/docker-compose.yaml up --build --exit-code-from test

build:
	docker build -t hamedsalim1999/go-samole:$(version) src 
	docker push hamedsalim1999/go-samole:$(version) 