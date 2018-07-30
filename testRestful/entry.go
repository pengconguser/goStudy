package main

import (
	"github.com/jinzhu/gorm"
)

type (
	//entity类
	article struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	//response entiry

	transformed struct {
		ID        unit   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
