package controllers

import (
	"github.com/gin-gonic/gin"
    "strconv"
    "backend/models"
)

func AddTransaction(c *gin.Context) {
    var req models.Transaction
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "message": "error - Invalid input",
            "data":    nil,
            "error":   err.Error(),
        })
        return
    }

    var foundProduct *models.Product
    for i := range models.Products {
        if models.Products[i].ID == req.ProductID {
            foundProduct = &models.Products[i]
            break
        }
    }
    if foundProduct == nil {
        c.JSON(404, gin.H{
            "message": "Error - Product not found",
            "data":    nil,
            "error":   "Product not found",
        })
        return
    }

    var foundSource *models.Source
    for i := range models.Sources {
        if models.Sources[i].ID == foundProduct.SourceID {
            foundSource = &models.Sources[i]
            break
        }
    }
    if foundSource == nil {
        c.JSON(404, gin.H{
            "message": "Error - Source for this product not found",
            "data":    nil,
            "error":   "Source for this product not found",
        })
        return
    }

    if req.Quantity <= 0 {
        c.JSON(400, gin.H{
            "message": "Error - Quantity must be greater than 0",
            "data":    nil,
            "error":   "Quantity must be greater than 0",
        })
        return
    }
    if foundProduct.Stock < req.Quantity {
        c.JSON(400, gin.H{
            "message": "Error - Not enough stock",
            "data":    nil,
            "error":   "Not enough stock",
        })
        return
    }

    foundProduct.Stock -= req.Quantity
    req.ID = strconv.Itoa(len(models.Transactions) + 1)
    req.Total = float64(req.Quantity) * foundProduct.Price
    models.Transactions = append(models.Transactions, req)

    c.JSON(201, gin.H{
        "message": "Success - Transaction created",
        "data":    req,
        "error":   nil,
    })
}

func GetTransactions(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "success get all transactions",
        "data":    models.Transactions,
        "error":   nil,
    })
}

func GetTransactionByID(c *gin.Context) {
    id := c.Param("id")
    for _, trx := range models.Transactions {
        if id == trx.ID {
            c.JSON(200, gin.H{
                "message": "success get transaction by ID",
                "data":    trx,
                "error":   nil,
            })
            return
        }
    }
    c.JSON(404, gin.H{
        "message": "Error - Transaction not found",
        "data":    nil,
        "error":   "Transaction not found",
    })
}