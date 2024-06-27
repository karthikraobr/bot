TOOLS_PATH=$$PWD/build/tools

define run-tool
$(eval SHELL:=/bin/bash)
_() { \
test -f "$(TOOLS_PATH)/$$1" || make install-deps ; \
cmd="$(TOOLS_PATH)/$$1"; \
cmd=".$${cmd#"$$PWD"}"; \
echo $$cmd "$${@:2}"; \
$$cmd "$${@:2}"; \
}; _
endef

.PHONY: install-deps
install-deps:
	@for var in $(shell go list -f '{{join .Imports " "}}' -e deps.go) ; do \
	if [[ $$var == *"sqlc"* ]] ;  then \
		var=$$var@latest ; \
	fi ; \
	if [[ $$var == *"migrate"* ]] ;  then \
		var="-tags 'postgres' "$$var ; \
	fi ; \
	(GOBIN=$(TOOLS_PATH) go install $$var) ; \
	done




.PHONY: migrate-db
migrate-db:
	@$(run-tool) migrate -database "postgresql://user:password@127.0.0.1:5432/customers?sslmode=disable" -path db/migrations up

.PHONY: run-server
run-server:
	docker-compose up --build


.PHONY: run-client
run-client:
	go run cmd/client/client.go


.PHONY: dev-psql
dev-psql:
	@$(run-tool) sqlc generate


