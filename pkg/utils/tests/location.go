package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	migrationFolder := "internal/database/postgres/migrations"
	abs, err := filepath.Abs(migrationFolder)
	if err != nil {
		fmt.Println("could not find absuolute ", err.Error())
		return
	}
	fmt.Println(abs)
}
