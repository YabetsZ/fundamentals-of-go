package main

import (
	"library-management/controllers"
	"library-management/services"
)

func main() {
	library := services.NewLibrary()
	controllers.NewLibraryController(library).Run()
}
