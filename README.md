[![codecov](https://codecov.io/gh/rezabintami/go-drop-logistik/branch/dev/graph/badge.svg?token=WIIYL0IBQN)](https://codecov.io/gh/rezabintami/go-drop-logistik)

# Clean Architecture Structure Folder

This is clean architecture structure folder in golang. In `app` folder have 3 folder that is `config` for convert environment file, `middleware` use for authenticated user, and `routes` for routing file into url.


## Usage
-> Create database on your system, name is free

1. Clone app from repo
2. Navigate to project folder
3. Execute go mod init
4. Execute go mod vendor
5. Adjust your config on `config.yaml` in `app/config/config.yaml`
6. Migrate your table using `go run cmd/migrations/migrations.go go-drop-logistik:migrate --up`