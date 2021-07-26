# Usage:
# make                # Compile full application
# make generate_sqlc  # Generate sqlc db queries

generate_sqlc:
	# Must be run in Unix system (for pwd)
	# docker pull kjconroy/sqlc  # First run
	docker run --rm -v $(shell pwd):/src -w /src kjconroy/sqlc generate
