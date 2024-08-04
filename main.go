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

type Order struct {
	Id    int64       `json:"Id"`
	Items []OrderItem `json:"menuItems"`
}

type OrderItem struct {
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}

func main() {
	router := gin.Default()
	router.GET("/orders", getOrders)

	//router.POST("/createOrder", createOrder)
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

//total: 6.2

var order2 = Order{
	Id:    2,
	Items: []OrderItem{{"Cappuccino", 2}, {"Croissant", 1}},
}

// total: 7.2

var orders = []Order{order1, order2}

// getOrders responds with the list of all orders as JSON.
func getOrders(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, orders)
}

// get total for an order by Id
func getOrderTotal(c *gin.Context) {
	id := c.Param("id")
	var total float64 = 0.0
	// Loop over the list of orders, looking for
	// an order whose ID value matches the parameter.
	for _, order := range orders {
		var ID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
		}
		if order.Id == ID {
			for _, menuItem := range order.Items {
				total += getPrice(menuItem.Title) * float64(menuItem.Quantity)
			}

		}
	}

	c.IndentedJSON(http.StatusOK, total)
	return
}

func getPrice(name string) (price float64) {
	for _, item := range menuItems {
		if item.Title == name {
			price = item.Price
			return
		}
	}

	return
}

/* postAlbums adds an album from JSON received in the request body.
func createOrder(c *gin.Context) {
	return
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
*/
