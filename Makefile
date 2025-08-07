DB_URL ?= postgres://owner@localhost:5432/rideshare_development

migrate-create:
	migrate create -ext sql -dir migrations $$(name)

migrate-up:
	migrate -path migrations -database '$(DB_URL)' up

migrate-down:
	migrate -path migrations -database '$(DB_URL)' down $(n)

