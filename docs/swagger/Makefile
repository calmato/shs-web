.PHONY: build generate lint validate

TEACHER_INPUT_DIR := 'teacher'
TEACHER_OUTPUT_DIR := '../../tmp/data/swagger/teacher'

build: build-teacher

build-teacher:
	$(MAKE) generate TYPE=html2 INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DIR}
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DIR}

build-html:
	$(MAKE) generate TYPE=html2 INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DIR}

build-yaml:
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DIR}

generate:
	openapi-generator-cli generate \
		-g ${TYPE} \
		-i ${INPUT_DIR}/openapi.yaml \
		-o ${OUTPUT_DIR}

lint:
	$(MAKE) validate INPUT_DIR=${TEACHER_INPUT_DIR}

validate:
	openapi-generator-cli validate \
		-i ${INPUT_DIR}/openapi.yaml \
		--recommend
