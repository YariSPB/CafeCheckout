package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type menuItem struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

type newOrderItem struct {
	OrderId int64  `json:"orderid"`
	Title   string `json:"title"`
	Count   int64  `json:"count"`
}

type Order struct {
	Id    int64       `json:"Id"`
	Items []OrderItem `json:"menuItems"`
	Total float64     `json:"total"`
}

type OrderItem struct {
	Title    string `json:"title"`
	Quantity int64  `json:"quantity"`
}

func main() {
	router := gin.Default()
	router.GET("/orders", getOrders)
	router.POST("/newOrder", newOrder)
	router.POST("/addItem", addItemInOrder)
	router.GET("/orders/:id", getOrderTotal)

	router.Run("localhost:8080")
}

var menuItems = []menuItem{
	{Title: "Americano", Price: 1.2},
	{Title: "Cappuccino", Price: 2.1},
	{Title: "Juice", Price: 2.5},
	{Title: "Baguette", Price: 1.5},
	{Title: "IceCream", Price: 3},
	{Title: "Croissant", Price: 3},
}

var order1 = Order{
	Id:    1,
	Items: []OrderItem{{"Juice", 2}, {"Americano", 1}},
}

var order2 = Order{
	Id:    2,
	Items: []OrderItem{{"Cappuccino", 2}, {"Croissant", 1}},
}

var orders = []Order{order1, order2}

// Get all orders as JSON.
func getOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)
}

// get total for an order by Id
func getOrderTotal(c *gin.Context) {
	id := c.Param("id")
	for _, order := range orders {
		var ID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		if order.Id == ID {
			for _, menuItem := range order.Items {
				order.Total += getPrice(menuItem.Title) * float64(menuItem.Quantity)
			}
			c.IndentedJSON(http.StatusOK, order)
			return
		}
	}
}

// helper function
func getPrice(name string) (price float64) {
	for _, item := range menuItems {
		if item.Title == name {
			price = item.Price
			return
		}
	}
	return
}

// helper function
func getOrderById(id int64) (orderOut *Order) {
	for i := 0; i < len(orders); i++ {
		if orders[i].Id == id {
			orderOut = &orders[i]
			return
		}
	}
	return
}

// create a new order and return its Id.
func newOrder(c *gin.Context) {
	var ID int64 = int64(len(orders)) + 1
	var newOrder Order = Order{
		Id:    ID,
		Items: []OrderItem{},
	}

	orders = append(orders, newOrder)
	c.IndentedJSON(http.StatusCreated, ID)
}

func addItemInOrder(c *gin.Context) {
	var orderItem newOrderItem
	if err := c.BindJSON(&orderItem); err != nil {
		return
	}

	currentOrder := getOrderById(orderItem.OrderId)
	for _, item := range menuItems {
		if item.Title == orderItem.Title {
			orderItem := OrderItem{
				Title:    orderItem.Title,
				Quantity: orderItem.Count,
			}

			(*currentOrder).Items = append(currentOrder.Items, orderItem)
		}
	}

	c.IndentedJSON(http.StatusCreated, currentOrder)
}
