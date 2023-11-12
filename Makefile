
#/-------------------------------------------------Переменные окружения----------------------------------------------------------/
GO = go


PROJECT_NAME = github.com/Markuysa/courceWorkBackendDev

ROOT_DIRS = cmd config internal migrations pkg utils

INTERNAL_DIRS = app system client admin


STRUCT_DIRS = admin client
#/-------------------------------------------------Переменные окружения----------------------------------------------------------/



#/-------------------------------------------------Основные команды----------------------------------------------------------/
lint:
	golangci-lint run -v --config golangci.yml


# для генерилки
build: create_dirs create_structure create_files
	@$(GO) mod init $(PROJECT_NAME)
	@$(GO) mod tidy

clean:
	@rm -rf $(ROOT_DIRS)
	@rm -rf internal
	@rm -f go.mod go.sum

# Прогон тестов
test:
	go test -v ./...

# Посмотреть покрытие в красивеньком хтмл файле в браузере
# Если у вас вообще нет тестовых файлов, то вы увидите черный экранчик :)
coverage:
	go test -coverprofile="coverage.out" ./...
	go tool cover -html="coverage.out" -o index.html
	open index.html
	rm coverage.out

lint:
	golangci-lint run -v --config .golangci.yml

mock:
	mockgen --source=usecase/interface.go -destination=../../mocks/InnerUCMock.go

.PHONY: test coverage build clean link
#/-------------------------------------------------Основные команды----------------------------------------------------------/




#/-------------------------------------------------Вспомогательные команды----------------------------------------------------------/
create_dirs:
	@$(foreach dir,$(ROOT_DIRS),mkdir -p $(dir);)
	@$(foreach dir,$(INTERNAL_DIRS),mkdir -p internal/$(dir);)

create_structure: $(STRUCT_DIRS)

create_files:
	@touch config/config-dev.json
	@touch config/config.go
	@touch config/config.json
	@touch cmd/main.go
	@echo 'package app' > internal/app/app.go
	@echo 'package main' > cmd/main.go

define create_internal_structure
	@mkdir -p $1/delivery
	@mkdir -p $1/repository
	@mkdir -p $1/usecase
	@echo 'package repository' > $1/repository/repository.go
	@echo 'package repository' > $1/repository/interface.go
	@echo 'package repository' > $1/repository/sql_queries.go
	@echo 'package usecase' > $1/usecase/usecase.go
	@echo 'package usecase' > $1/usecase/interface.go
endef

$(STRUCT_DIRS): %:
	$(call create_internal_structure,internal/$@)

#/-------------------------------------------------Вспомогательные команды----------------------------------------------------------/