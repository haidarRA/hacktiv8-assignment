package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "strconv"
    "backend/models"
)

func GetSources(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "success get all sources",
        "data":    models.Sources,
        "error":   nil,
    })
}

func GetSourceByID(c *gin.Context) {
	id := c.Param("id")
	for _, source := range models.Sources {
		if id == source.ID {
			c.JSON(200, gin.H{
				"message": "success get source by ID",
				"data": source,
                "error": nil,
			})
			return
		}
	}
    c.JSON(404, gin.H{
        "message": "Error - Source not found",
        "data":    nil,
        "error":   "Source not found",
    })
}

func AddSources(c *gin.Context) {
	var newSource models.Source
	if err := c.ShouldBindJSON(&newSource); err != nil {
        c.JSON(400, gin.H{
            "message": "error - Invalid input",
            "data":    nil,
            "error":   err.Error(),
        })
		return
	}

    newSource.ID = strconv.Itoa(len(models.Sources) + 1)
	models.Sources = append(models.Sources, newSource)
    c.JSON(201, gin.H{
        "message": "Success - Source added",
        "data":    newSource,
        "error":   nil,
    })
}

func UpdateSource(c *gin.Context) {
    id := c.Param("id")

    var updateData map[string]interface{}
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(400, gin.H{
            "message": "error - Invalid input",
            "data":    nil,
            "error":   err.Error(),
        })
        return
    }

    for i, source := range models.Sources {
        if id == source.ID {
            if name, ok := updateData["name"].(string); ok {
                source.Name = name
			}
			
            models.Sources[i] = source
            c.JSON(200, gin.H{
                "message": "Success - Source updated",
                "data":    source,
                "error":   nil,
            })
            return
        }
    }

    c.JSON(404, gin.H{
        "message": "Error - Source not found",
        "data":    nil,
        "error":   "Source not found",
    })
}

func DeleteSource(c *gin.Context) {
	id := c.Param("id")
	for i, source := range models.Sources {
		if id == fmt.Sprintf("%v", source.ID) {
			models.Sources = append(models.Sources[:i], models.Sources[i+1:]...)
            c.JSON(200, gin.H{
                "message": "Success - Source deleted",
                "data":    nil,
                "error":   nil,
            })
			return
		}
	}
    c.JSON(404, gin.H{
        "message": "Error - Source not found",
        "data":    nil,
        "error":   "Source not found",
    })
}