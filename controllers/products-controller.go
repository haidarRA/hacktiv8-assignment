package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
    "strconv"
    "backend/models"
)

func GetProducts(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "success get all products",
        "data":    models.Products,
        "error":   nil,
    })
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	for _, product := range models.Products {
		if id == product.ID {
            c.JSON(200, gin.H{
                "message": "success get product by ID",
                "data":    product,
                "error":   nil,
            })
			return
		}
	}
    c.JSON(404, gin.H{
        "message": "Error - Product not found",
        "data":    nil,
        "error":   "Product not found",
    })
}

func AddProducts(c *gin.Context) {
	var newProduct models.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
        c.JSON(400, gin.H{
            "message": "error - Invalid input",
            "data":    nil,
            "error":   err.Error(),
        })
		return
	}

	if newProduct.Name == "" || newProduct.Price <= 0 || newProduct.Stock < 0 {
        c.JSON(400, gin.H{
            "message": "error - Invalid product input",
            "data":    nil,
            "error":   "Invalid product input",
        })
		return
	}

    newProduct.ID = strconv.Itoa(len(models.Products) + 1)
	models.Products = append(models.Products, newProduct)
	c.JSON(201, gin.H{
		"message": "Success - Product added",
		"data": newProduct,
        "error":   nil,
	})
}

func UpdateProduct(c *gin.Context) {
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

    if name, ok := updateData["name"]; ok {
        if strName, ok := name.(string); !ok || strName == "" {
            c.JSON(400, gin.H{
                "message": "error - Invalid product input",
                "data":    nil,
                "error":   "Product name cannot be empty",
            })
            return
        }
    }

    if price, ok := updateData["price"]; ok {
        var priceVal float64
        switch v := price.(type) {
        case float64:
            priceVal = v
        case int:
            priceVal = float64(v)
        }
        if priceVal <= 0 {
            c.JSON(400, gin.H{
                "message": "error - Invalid product input",
                "data":    nil,
                "error":   "Product price must be greater than 0",
            })
            return
        }
    }
	
    if stock, ok := updateData["stock"]; ok {
        var stockVal int
        switch v := stock.(type) {
        case float64:
            stockVal = int(v)
        case int:
            stockVal = v
        }
        if stockVal < 0 {
            c.JSON(400, gin.H{
                "message": "error - Invalid product input",
                "data":    nil,
                "error":   "Product stock cannot be negative",
            })
            return
        }
    }

    for i, product := range models.Products {
        if id == product.ID {
            if name, ok := updateData["name"].(string); ok {
                product.Name = name
            }
            if desc, ok := updateData["description"].(string); ok {
                product.Description = desc
            }
            if price, ok := updateData["price"].(float64); ok {
                product.Price = price
            }
            if stock, ok := updateData["stock"]; ok {
                switch v := stock.(type) {
                case float64:
                    product.Stock = int(v)
                case int:
                    product.Stock = v
                }
            }
            if sourceID, ok := updateData["source_id"].(string); ok {
                product.SourceID = sourceID
            }
            models.Products[i] = product
            c.JSON(200, gin.H{
                "message": "Success - Product updated",
                "data":    product,
                "error":   nil,
            })
            return
        }
    }

    c.JSON(404, gin.H{
        "message": "Error - Product not found",
        "data":    nil,
        "error":   "Product not found",
    })
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, product := range models.Products {
		if id == fmt.Sprintf("%v", product.ID) {
			models.Products = append(models.Products[:i], models.Products[i+1:]...)
			c.JSON(200, gin.H{
				"message": "Success - Product deleted",
			})
			return
		}
	}
    c.JSON(404, gin.H{
        "message": "Error - Product not found",
        "data":    nil,
        "error":   "Product not found",
    })
}