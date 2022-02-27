.PHONY: build generate lint validate

TEACHER_INPUT_DIR       := teacher
TEACHER_OUTPUT_DOCS_DIR := ../../tmp/data/swagger/teacher/docs
TEACHER_OUTPUT_TS_DIR   := ../../tmp/data/swagger/teacher/types
TEACHER_WEB_TYPES_DIR   := ../../web/teacher/src/types/api/__tmp__
STUDENT_INPUT_DIR       := student
STUDENT_OUTPUT_DOCS_DIR := ../../tmp/data/swagger/student/docs
STUDENT_OUTPUT_TS_DIR   := ../../tmp/data/swagger/student/types
STUDENT_WEB_TYPES_DIR   := ../../web/student/src/types/api/__tmp__

build: build-teacher build-student

build-teacher:
	$(MAKE) generate TYPE=html2 INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=typescript-axios INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_TS_DIR}
	mkdir -p ${TEACHER_WEB_TYEPS_DIR}; mv ${TEACHER_OUTPUT_TS_DIR}/*.ts ${TEACHER_WEB_TYEPS_DIR}

build-student:
	$(MAKE) generate TYPE=html2 INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=typescript-axios INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_TS_DIR}
	mkdir -p ${STUDENT_WEB_TYEPS_DIR}; mv ${STUDENT_OUTPUT_TS_DIR}/*.ts ${STUDENT_WEB_TYEPS_DIR}

build-html:
	$(MAKE) generate TYPE=html2 INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=html2 INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_DOCS_DIR}

build-yaml:
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_DOCS_DIR}
	$(MAKE) generate TYPE=openapi-yaml INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_DOCS_DIR}

build-typescript:
	$(MAKE) generate TYPE=typescript-axios INPUT_DIR=${TEACHER_INPUT_DIR} OUTPUT_DIR=${TEACHER_OUTPUT_TS_DIR}
	$(MAKE) generate TYPE=typescript-axios INPUT_DIR=${STUDENT_INPUT_DIR} OUTPUT_DIR=${STUDENT_OUTPUT_TS_DIR}
	mkdir -p ${TEACHER_WEB_TYPES_DIR}; mv ${TEACHER_OUTPUT_TS_DIR}/*.ts ${TEACHER_WEB_TYPES_DIR}
	mkdir -p ${STUDENT_WEB_TYPES_DIR}; mv ${STUDENT_OUTPUT_TS_DIR}/*.ts ${STUDENT_WEB_TYPES_DIR}

generate:
	openapi-generator-cli generate \
		-g ${TYPE} \
		-i ${INPUT_DIR}/openapi.yaml \
		-o ${OUTPUT_DIR} \

lint:
	$(MAKE) validate INPUT_DIR=${TEACHER_INPUT_DIR}
	$(MAKE) validate INPUT_DIR=${STUDENT_INPUT_DIR}

validate:
	openapi-generator-cli validate \
		-i ${INPUT_DIR}/openapi.yaml \
		--recommend
