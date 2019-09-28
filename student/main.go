package main

import (
	"github.com/MohitArora1/devops/student/controller"
	"github.com/MohitArora1/devops/student/utils"
)

func main() {
	utils.InitConfig()
	controller.RunController(":8080")
}
