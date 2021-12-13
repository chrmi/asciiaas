COLOR_A=\033[90m
COLOR_B=\033[91m
COLOR_C=\033[92m
COLOR_D=\033[93m
COLOR_E=\033[94m
COLOR_F=\033[95m
COLOR_G=\033[96m

STOP_COLOR=\033[0m

RAINBOW=${COLOR_A}=${COLOR_B}-${COLOR_C}=${COLOR_D}-${COLOR_E}=${COLOR_F}-${COLOR_G}=

DOUBLE_RAINBOW=${COLOR_A}=${COLOR_B}-${COLOR_C}=${COLOR_D}-${COLOR_E}=${COLOR_F}-${COLOR_G}=${COLOR_A}=${COLOR_B}-${COLOR_C}=${COLOR_D}-${COLOR_E}=${COLOR_F}-${COLOR_G}=

GOPKGS = $(shell go list ./...)

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

SRC_ROOT="."
BUILD_PATH="./build"
DATA_PATH="./data"
MOD_NAME="asciiaas"

.PHONY: init test clean build run package demo

default: test clean build run

init:
	@echo "${RAINBOW} ${COLOR_C}Init: ${COLOR_D}${MOD_NAME}${STOP_COLOR}"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Removing vendor folder${STOP_COLOR}"
	@rm -rf vendor
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Removing and recreating ${COLOR_E}${BUILD_PATH}${COLOR_F} folder${STOP_COLOR}"
	@rm -rf $(BUILD_PATH) && mkdir $(BUILD_PATH)
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Getting dependencies for: ${COLOR_D}${MOD_NAME} ${STOP_COLOR}"
	@go get -t $(MOD_NAME)
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Removing unused dependencies${STOP_COLOR}"
	@go mod tidy
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Saving dependencies to vendor folder${STOP_COLOR}"
	@go mod vendor
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}Showing how dependencies are structured${STOP_COLOR}"
	@go mod graph

clean:
	@echo "${RAINBOW} ${COLOR_C}Clean: ${COLOR_D}${MOD_NAME}${STOP_COLOR}"
	@rm -rf $(BUILD_PATH)

test:
	@echo "${RAINBOW} ${COLOR_C}Test: ${COLOR_D}${MOD_NAME}${STOP_COLOR}"
	@go test ./... --cover -v $(GOPKGS)
	@cd asciimodel && go get -t asciimodel && go test
	@cd asciigenerator && go get -t asciigenerator && go test
	@cd asciibuilder && go get -t asciibuilder && go test

build:
	@echo "${RAINBOW} ${COLOR_C}Build: ${COLOR_D}${MOD_NAME}${COLOR_C} for ${COLOR_G}${GOOS}${COLOR_C} on ${COLOR_E}${GOARCH}${STOP_COLOR}"
	@go get -t $(MOD_NAME)
	@GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	go build -o $(BUILD_PATH)/$(MOD_NAME) $(SRC_ROOT) 
	@cp -r $(DATA_PATH) $(BUILD_PATH)/

run:
	@echo "${RAINBOW} ${COLOR_C}Run: ${COLOR_D}${MOD_NAME}${STOP_COLOR}"
	@cd ./$(BUILD_PATH) && ./$(MOD_NAME)

package:
	@docker build -t asciiaas/v1 .

demo:
	@echo "\n"
	@echo "${RAINBOW} ${COLOR_C}Demo: ${COLOR_D}${MOD_NAME}${COLOR_C} for ${COLOR_G}${GOOS}${COLOR_C}${STOP_COLOR}"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}GET from demo.json${STOP_COLOR}"
	@curl "0.0.0.0:8099/text?key=wave-top"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}GET to show empty value${STOP_COLOR}"
	@curl "0.0.0.0:8099/text?key=demo"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}PUT new value${STOP_COLOR}"
	@curl -X PUT -d "{\"demo\": \"New Demo Value.\"}" -H 'Content-Type: application/json' "http://0.0.0.0:8099/text"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}GET to show new value${STOP_COLOR}"
	@curl "0.0.0.0:8099/text?key=demo"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}POST change to existing value${STOP_COLOR}"
	@curl -X POST -d "{\"demo\": \"Replaced Demo Value\"}" -H 'Content-Type: application/json' "http://0.0.0.0:8099/text"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}GET to show modified value${STOP_COLOR}"
	@curl "0.0.0.0:8099/text?key=demo"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}DELETE value${STOP_COLOR}"
	@curl -X DELETE "http://0.0.0.0:8099/text?key=demo"
	@echo "\n"
	@echo "${DOUBLE_RAINBOW} ${COLOR_F}GET to show removed value${STOP_COLOR}"
	@curl "0.0.0.0:8099/text?key=demo"
	@echo "\n\n"


