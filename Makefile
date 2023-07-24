stop:
	docker-compose -f src/docker-compose.yaml stop

start:
	docker-compose -f src/docker-compose.yaml up  -d

test:
	docker-compose -f src/docker-compose.yaml up --build --exit-code-from test

































