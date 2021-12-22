##################################################
# Container Commands - Run All
##################################################
.PHONY: setup build install start stop down remove logs ps

setup: build install proto swagger migrate
	if [ ! -f $(PWD)/.env ]; then \
		cp $(PWD)/.env.temp $(PWD)/.env; \
	fi

build:
	docker-compose build --parallel

install:
	docker-compose run --rm teacher_web yarn
	docker-compose run --rm swagger_generator yarn

start:
	docker-compose up --remove-orphans

stop:
	docker-compose stop

down:
	docker-compose down

remove:
	docker-compose down --rmi all --volumes --remove-orphans

logs:
	docker-compose logs

ps:
	docker-compose ps

##################################################
# Container Commands - Run Container Group
##################################################
.PHONY: start-api start-swagger start-test

start-web:
	docker-compose up teacher_web

start-api:
	docker-compose up teacher_gateway user_api classroom_api mysql

start-swagger:
	docker-compose up swagger_generator swagger_teacher

start-test:
	docker-compose up mysql_test firebase_test

##################################################
# Container Commands - Single
##################################################
.PHONY: proto swagger migrate

proto:
	docker-compose run --rm proto bash -c "cd ./api; make install; make protoc"

swagger:
	docker-compose run --rm swagger_generator yarn generate

migrate:
	$(PWD)/bin/database-migrate.sh
