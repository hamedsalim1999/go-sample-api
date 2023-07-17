stop:
	docker-compose stop

start:
	docker-compose up -d

test:
	docker-compose up --build --exit-code-from test

































