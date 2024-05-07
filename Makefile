export DATABASE_URL=postgres://admin:123456@localhost:5432/training?sslmode=disable

run:
	go run cmd/main.go

run-debug:
	DEBUG=1 go run cmd/main.go

# make create-migration name=add_a_column
create-migration:
	migrate create -ext sql -dir internal/pg/migration/sql -seq $(name)

# make migration-up num=1
migration-up:
	migrate -database ${DATABASE_URL} -path internal/pg/migration/sql up ${num}

# make migration-down num=1
migration-down:
	migrate -database ${DATABASE_URL} -path internal/pg/migration/sql down $(num)

sql-gen:
	sqlboiler --wipe psql --output internal/pg/model --pkgname modelpg


