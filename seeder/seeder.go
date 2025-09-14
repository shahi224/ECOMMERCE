package seeder

import "ECOMMERCE/utils/seeder"

func GroupSeeder() {
	seeder.SeedCategories()
	seeder.SeedBrands()
	seeder.SeedProducts()
	seeder.SeedUsers()
}
