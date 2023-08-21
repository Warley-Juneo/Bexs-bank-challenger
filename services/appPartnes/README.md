mockgen -source=repository/partnerrepository/partnerrepository.go -destination=repository/partnerrepository/partnerrepository_test.go -package=mocks

#####################################################
migrate create -ext sql -dir database/migrations -seq create_product_table

migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force 2
#####################################################
