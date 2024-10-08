.PHONY: keypair migrate-create migrate-up migrate-down migrate-force

PWD = $(shell pwd)
ACCTPATH = $(PWD)/account
MPATH = $(ACCTPATH)/migrations
PORT = 5432

# Default number of migrations to execute up or down
N = 1

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(ACCTPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(ACCTPATH)/rsa_private_$(ENV).pem -pubout -out $(ACCTPATH)/rsa_public_$(ENV).pem

#
## Detect OS and adjust path separator and shell
#ifeq ($(OS),Windows_NT)
#    PATH_SEP := \\
#    SHELL := cmd
#    RM := del /Q
#    NULL_DEVICE := NUL
#else
#    PATH_SEP := /
#    SHELL := /bin/sh
#    RM := rm -f
#    NULL_DEVICE := /dev/null
#endif
#
## Paths and configurations
## Paths and configurations
#MPATH := account/migrations
#PORT := 5432
#NAME ?= add_migration
#N ?= 1
#VERSION ?= 1
#
## Targets
#migrate-create:
#	@echo "---Creating migration files---"
#	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)
#
#migrate-up:
#	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable up $(N)
#
#migrate-down:
#	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable down $(N)
#
#migrate-force:
#	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)

migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)

