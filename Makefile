run:
	docker-compose up -d
	docker image prune

stop:
	docker-compose stop

delete:
	docker-compose down
	docker volume rm my-forum_cache my-forum_pg-data