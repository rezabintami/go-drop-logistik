[![codecov](https://codecov.io/gh/rezabintami/go-drop-logistik/branch/dev/graph/badge.svg?token=WIIYL0IBQN)](https://codecov.io/gh/rezabintami/go-drop-logistik)

# Go Drop Logistik
Drop Logistic Indonesia is a freight forwarding company that handles retail goods packages and other logistics services. The entire territory of Indonesia has become our reach, with the support of 5 Branch Offices and more than hundreds of customers spread across major cities and regencies. We have been serving the delivery of goods packages starting in the Java Island area, then expanding to all over Indonesia, it is proven that we have earned the trust of customers. We are increasingly confident to continue to improve service to customers by optimally making improvements to facilities and infrastructure.

## Architecture
- Language [GoLang](https://golang.org/)
- Framework [Framework Echo](https://echo.labstack.com/)
- Object Relational Mapping [Gorm io](https://gorm.io/docs/index.html)
- Postgres
- JWT [JWT](https://github.com/dgrijalva/jwt-go)
- Log [Logrus](https://github.com/sirupsen/logrus)

## Usage
-> Create database on your system, name is free

1. Clone app from repo
2. Navigate to project folder
3. Execute go mod init
4. Execute go mod vendor
5. Adjust your config on `config.yaml` in `app/config/config.yaml`
6. Migrate your table using `go run cmd/migrations/migrations.go go-drop-logistik:migrate --up`

## How to Add Migration

1. Install [CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) in your computer.
2. Create file migration using `migrate create -ext sql -dir your_project_dir -seq your_file_name` 
3. Example cli command `migrate create -ext sql -dir drivers/postgres/files/migrations -seq add_foreign_key_tracks`

## How to Add Mock

1. Go to usecase folder.
2. run cli command `mockery --all`