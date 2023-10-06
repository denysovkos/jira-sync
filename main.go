package main

import (
	"fmt"
	"jira-sync/pkg/db"
	"jira-sync/pkg/ui"
)

// Functions:
// PARAMETERS
// 1. get user parameters
// 2. set user parameters

// JIRA ISSUES
// 1. get tasks without time tracking
// 2. set time for particular issue

// UI
// PARAMETERS (CONFIG)
// 1. add parameters
// 2. view parameters
// 3. delete parameters
// 4. edit parameters

// JIRA ISSUES
// 1. issue table
// 2. open particular issue
// 3. log time

func main() {
	db, err := db.NewDB()
	if err != nil {
		fmt.Println("Error creating DB instance:", err)
		return
	}

	ui.Run(db)
}
