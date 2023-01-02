package model

func migration() {
	err := DB.AutoMigrate(&Issue{})

	if err != nil {
		panic("failed to run data migration")
	}

	// seed issues
	var seedIssues = []Issue{
		{
			Title:       "issue 1",
			Description: "This is issue 1",
		},
		{
			Title:       "issue 2",
			Description: "This is issue 2",
		},
	}
	DB.Create(&seedIssues)
}
