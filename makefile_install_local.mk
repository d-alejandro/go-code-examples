.PHONY: go-install
go-install:
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest
	go get -u golang.org/x/lint/golint && \
	go install golang.org/x/lint/golint && \
	go mod tidy

.PHONY: gomock
gomock:
	go install go.uber.org/mock/mockgen@latest && \
	mockgen -version

.PHONY: pre-commit
pre-commit:
	sudo apt update && \
	sudo apt install -y pre-commit && \
	pre-commit -V

.PHONY: allure-inst
allure-inst:
	sudo apt update && \
	sudo apt install -y default-jre
	curl -LO 'https://github.com/allure-framework/allure2/releases/download/2.32.0/allure_2.32.0-1_all.deb' && \
	sudo dpkg -i 'allure_2.32.0-1_all.deb' && \
	rm 'allure_2.32.0-1_all.deb' && \
	sudo apt install -f && \
	allure --version

.PHONY: list
list: NAME = pre-commit
list:
	sudo dpkg --list | grep --color=never -E '(\|\|\/|$(NAME))'

.PHONY: remove
remove: NAME = pre-commit
remove:
	sudo apt remove -y $(NAME)
