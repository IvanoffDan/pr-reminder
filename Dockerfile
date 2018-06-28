# Base Image
FROM golang:latest AS base
ARG APP_PATH

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
 

COPY . ${APP_PATH}/
WORKDIR ${APP_PATH}
RUN dep ensure
 
 
CMD if [ ${ENV} = production ]; \
  then \
  go build -o ${GOPATH}/bin/${APP_NAME} ${APP_PATH}/main.go && \
  ${GOPATH}/bin/${APP_NAME}; \
  else \
  go get github.com/IvanoffDan/fresh && \
  fresh -c ${APP_PATH}/runner.conf; \
  fi