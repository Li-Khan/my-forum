run:
	docker-compose up -d
	docker image prune
	docker rmi alpine:latest

stop:
	docker-compose stop

delete:
	docker-compose down
	docker volume rm my-forum_cache my-forum_pg-data