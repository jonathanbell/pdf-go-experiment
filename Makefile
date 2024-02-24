build:
	docker build --tag win-loss-pdf-service .

run:
	docker run -d -p 9999:9999 --name win-loss-pdf-service win-loss-pdf-service

dev:
	docker-compose up -d win-loss-pdf-service

down:
	docker-compose down
