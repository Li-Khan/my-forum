run:
	docker-compose up -d

stop:
	docker-compose stop

delete:
	docker-compose down
	docker volume rm my-forum_cache  my-forum_pg-data