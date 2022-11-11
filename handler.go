package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	ref *Users
}

func (h *Handler) Donate(c *gin.Context) {
	var objUsr User
	if err := c.ShouldBindJSON(&objUsr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.ref.Donate(objUsr.Money, objUsr.Id)
	c.JSON(http.StatusOK, gin.H{
		"Data": h.ref.Users[objUsr.Id],
	})

}
func (h *Handler) Trade(c *gin.Context) {
	var objTrade moneyTrade
	if err := c.ShouldBindJSON(&objTrade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if objTrade.Idsrc == objTrade.Iddst {
		c.JSON(http.StatusBadRequest, gin.H{"error": "You cant trade money with yourself!"})
		return
	}
	var flag bool
	for k, v := range h.ref.Users {
		if k == objTrade.Idsrc && v.Money >= objTrade.Amount {
			v.Money -= objTrade.Amount
			h.ref.Trade(objTrade.Amount, objTrade.Iddst)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Money have been sent!"})
			flag = true
		}
	}
	if !flag {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not enough money"})
	}

}

func (h *Handler) Info(c *gin.Context) {
	var objUsr idUsr
	if err := c.ShouldBindJSON(&objUsr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	amount := h.ref.getInfo(objUsr.Id)
	c.JSON(http.StatusOK, gin.H{
		"amount of money": amount,
	})

	h.ref.getInfo(objUsr.Id)
}

func (h *Handler) Buy(c *gin.Context) {
	var objMarket Marketplace
	if err := c.ShouldBindJSON(&objMarket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for k, v := range h.ref.Users {
		if k == objMarket.UserId {
			if v.Money >= objMarket.Cost {
				v.InProcess = true
				v.Money -= objMarket.Cost
				v.Market = objMarket
				c.JSON(http.StatusOK, gin.H{"Succesfully created order!": objMarket})
			} else {
				c.JSON(http.StatusOK, gin.H{"error": "Not enough money!"})
				return
			}
		}
	}

}
func (h *Handler) Accept(c *gin.Context) {
	var objMarket Marketplace
	if err := c.ShouldBindJSON(&objMarket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for k, v := range h.ref.Users {
		if k == objMarket.UserId && v.InProcess && v.Market == objMarket {
			if v.Market == objMarket {
				c.JSON(http.StatusOK, gin.H{"Succesfully approved order!": objMarket})
				v.InProcess = false
				WriteCSV(v.Market)
				v.Market = Marketplace{}
			}

		}
	}

}

func (h *Handler) All(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Data": h.ref})
}
