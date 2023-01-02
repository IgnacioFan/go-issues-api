package model

func migration() {
	err := DB.AutoMigrate(&Issue{})

	if err != nil {
		panic("failed to run data migration")
	}
}
