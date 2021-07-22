# Usage:
# make                # compile full application
# make generate_sqlc  # generate sqlc db queries

generate_sqlc:
	# must be run in Unix system (for pwd)
	# docker pull kjconroy/sqlc  # first run
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
