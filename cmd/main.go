package main

import (
	"fmt"
	"frappucino/cmd/help"
	"net/http"
	"regexp"
	"strings"
)

type RequestHandler struct{}

const (
	host     = "db"
	port     = ":5432"
	user     = "latte"
	password = "latte"
	database = "frappuccino"
)

// Error(w, 400, "Invalid Request.")

var (
	REMenu          = regexp.MustCompile(`^\/menu\/?$`)
	REMenuItem      = regexp.MustCompile(`^\/menu\/\w+\/?$`)
	REInventory     = regexp.MustCompile(`^\/inventory$`)
	REInventoryItem = regexp.MustCompile(`^\/inventory\/\w+\/?$`)
	REOrders        = regexp.MustCompile(`^\/orders\/?$`)
	REOrdersID      = regexp.MustCompile(`^\/orders\/\w+\/?$`)
	REOrdersIDClose = regexp.MustCompile(`^\/orders\/\w+/close\/?$`)
	REPopularItems  = regexp.MustCompile(`^\/reports\/popular-items\/?$`)
	RETotalSales    = regexp.MustCompile(`^\/reports\/total-sales\/?$`)
)

func (req *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	// ORDERS
	case r.Method == http.MethodPost && REOrders.MatchString(r.URL.Path): // Create a new order.
		fmt.Println("Create a new order.")
		return

	case r.Method == http.MethodGet && REOrders.MatchString(r.URL.Path): // Retrieve all orders.
		fmt.Println("Retrieve all orders.")
		return

	case r.Method == http.MethodGet && REOrdersID.MatchString(r.URL.Path): // Retrieve a specific order by ID.
		fmt.Println("Retrieve a specific order by ID.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodPut && REOrdersID.MatchString(r.URL.Path): // Update an existing order.
		fmt.Println("Update an existing order.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodDelete && REOrdersID.MatchString(r.URL.Path): // Delete an order.
		fmt.Println("Delete an order.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodPost && REOrdersIDClose.MatchString(r.URL.Path): // Close an order.
		fmt.Println("Close an order.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

		// Menu
	case r.Method == http.MethodPost && REMenu.MatchString(r.URL.Path): // Add a new menu item.
		fmt.Println("Add a new menu item.")
		return

	case r.Method == http.MethodGet && REMenu.MatchString(r.URL.Path): // Retrieve all menu items.
		fmt.Println("Retrieve all menu items.")
		return

	case r.Method == http.MethodGet && REMenu.MatchString(r.URL.Path): // Retrieve a specific menu item.
		fmt.Println("Retrieve a specific menu item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodPut && REMenuItem.MatchString(r.URL.Path): // Update a menu item.
		fmt.Println("Update a menu item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodDelete && REMenuItem.MatchString(r.URL.Path): // Delete a menu item.
		fmt.Println("Delete a menu item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

		// Inventory
	case r.Method == http.MethodPost && REInventory.MatchString(r.URL.Path): // Add a new inventory item.
		fmt.Println("Add a new inventory item.")
		return

	case r.Method == http.MethodGet && REInventory.MatchString(r.URL.Path): // Retrieve all inventory items.
		fmt.Println("Retrieve all inventory items.")
		return

	case r.Method == http.MethodGet && REInventoryItem.MatchString(r.URL.Path): // Retrieve a specific inventory item.
		fmt.Println("Retrieve a specific inventory item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodPut && REInventoryItem.MatchString(r.URL.Path): // Update an inventory item.
		fmt.Println("Update an inventory item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

	case r.Method == http.MethodDelete && REInventoryItem.MatchString(r.URL.Path): // Delete an inventory item.
		fmt.Println("Delete an inventory item.")
		id := strings.Split(r.URL.Path, "/")[2]
		fmt.Println(id)
		return

		// Aggregations
	case r.Method == http.MethodGet && RETotalSales.MatchString(r.URL.Path): // Get the total sales amount.
		fmt.Println("Get the total sales amount.")
		return

	case r.Method == http.MethodGet && REPopularItems.MatchString(r.URL.Path): // Get a list of popular menu items.
		fmt.Println("Get a list of popular menu items.")
		return

	default:
		help.Error(w, 400, "Invalid Request.")
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &RequestHandler{})

	http.ListenAndServe(port, mux)
}
