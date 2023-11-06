.PHONY: docker_build, docker_status, docker_stop, docker_delete, swag

docker_build:
	docker-compose up -d --build
	docker-compose ps

docker_status:
	docker exec -it go-db bash -c "mysql -u root -p -h localhost -P 3306 -p123456 mydb -e 'show tables;'"

docker_stop:
	docker-compose down

docker_delete:
	docker rm -f $(docker ps -aq)
	docker rmi go-ddd
	docker network rm go-ddd_custom-local-net

swag:
	swag init
