
PROJECT := example_project
BUILD_DIR ?= build
BUILD_TYPE ?= Release
VERSION ?= 1.0.0

.PHONY: all clean

all: ${PROJECT}

${PROJECT}: main.o ${BUILD_DIR}
	@echo "Linking ${PROJECT}..."
	g++ ${BUILD_DIR}/main.o -o ${PROJECT} -L${BUILD_DIR} -lstdc++

main.o: main.cpp ${BUILD_DIR}
	@echo "Compiling main.cpp..."
	g++ -c main.cpp -o ${BUILD_DIR}/main.o

${BUILD_DIR}:
	mkdir -p ${BUILD_DIR}

clean:
	rm -rf $(BUILD_DIR)
