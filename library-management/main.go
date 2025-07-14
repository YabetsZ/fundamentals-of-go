package main

import (
	"fundamentals-of-go/library-management/controllers"
	"fundamentals-of-go/library-management/services"
)

func main() {
	library := services.NewLibrary()
	controllers.NewLibraryController(library).Run()
}