##################################################
# Container Commands - Run All
##################################################
.PHONY: setup build install start stop down remove logs

setup: build install proto swagger
	if [ ! -f $(PWD)/.env ]; then \
		cp $(PWD)/.env.temp $(PWD)/.env; \
	fi

build:
	docker-compose build --parallel

install:
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

##################################################
# Container Commands - Run Container Group
##################################################
.PHONY: start-api start-swagger start-test

start-api:
	docker-compose up teacher_gateway user_api mysql

start-swagger:
	docker-compose up swagger_generator swagger_ui_teacher

start-test:
	docker-compose up mysql_test firebase_test

##################################################
# Container Commands - Single
##################################################
.PHONY: proto swagger

proto:
	docker-compose run --rm proto bash -c "cd ./api; make install; make protoc"

swagger:
	docker-compose run --rm swagger_generator yarn generate
