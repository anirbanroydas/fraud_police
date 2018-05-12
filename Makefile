APP_NAME = fraud_police
PROJECT_ROOT_DIR = $(PWD)

# Change thiese values accroding to the make build-tag-push requirement
BRANCH = master
COMMIT = 

DEPLOY_ENVIRONMENT = production
DOCKER = docker 
DOCKER_COMPOSE = docker-compose


.PHONY: build-tag-push deploy

build-tag-push:
	bash -c "build_tag_push.sh $(APP_NAME) $(BRANCH) $(COMMIT)"
