/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 24-11-2015
 * |
 * | File Name:     desgo.go
 * +===============================================
 */
package main

import (
	"./desgo"
	"fmt"
)

func main() {
	err := desgo.ChangeDesktopBackground("~/Downloads/m.jpg")
	if err != nil {
		fmt.Errorf("Desgo: %v\n", err)
	}
}
