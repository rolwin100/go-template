
APP_NAME = apiserver
MIGRATIONS_FOLDER = $(PWD)/app/migrations
DATABASE_URL = postgres://root:root@localhost/adopler?sslmode=disable

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

echo:
	echo $(MIGRATIONS_FOLDER)