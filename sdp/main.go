package main

import (
	"fmt"
)

type Beverage interface {
	Cost() int
	Description() string
}

type Espresso struct{}

func (e Espresso) Cost() int {
	return 150
}

func (e Espresso) Description() string {
	return "Espresso"
}

type Latte struct{}

func (l Latte) Cost() int {
	return 200
}

func (l Latte) Description() string {
	return "Latte"
}

type MilkDecorator struct {
	beverage Beverage
}

func (m MilkDecorator) Cost() int {
	return m.beverage.Cost() + 50
}

func (m MilkDecorator) Description() string {
	return m.beverage.Description() + " with Milk"
}

type BeverageFactory interface {
	Create() Beverage
}

type EspressoFactory struct{}

func (e EspressoFactory) Create() Beverage {
	return Espresso{}
}

type LatteFactory struct{}

func (l LatteFactory) Create() Beverage {
	return Latte{}
}

type Order struct {
	items []Beverage
}

func NewOrder() *Order {
	return &Order{items: make([]Beverage, 0)}
}

func (o *Order) AddItem(item Beverage) {
	o.items = append(o.items, item)
}

func (o *Order) RemoveLastItem() {
	if len(o.items) > 0 {
		o.items = o.items[:len(o.items)-1]
	}
}

func (o *Order) ClearOrder() {
	o.items = make([]Beverage, 0)
}

func (o *Order) GetTotalCost() int {
	total := 0
	for _, item := range o.items {
		total += item.Cost()
	}
	return total
}

func main() {
	order := NewOrder()
	var factory BeverageFactory

	for {
		fmt.Println("Welcome to the Coffee Shop!")
		fmt.Println("Choose an option:")
		fmt.Println("1. Serve Espresso")
		fmt.Println("2. Serve Latte")
		fmt.Println("3. Add Milk")
		fmt.Println("4. Show Order")
		fmt.Println("5. Cancel Last Item")
		fmt.Println("6. Checkout and Exit")
		fmt.Println("7. Clear Order")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			factory = EspressoFactory{}
			beverage := factory.Create()
			order.AddItem(beverage)
			fmt.Printf("Added %s for %d KZT to the order.\n", beverage.Description(), beverage.Cost())
		case 2:
			factory = LatteFactory{}
			beverage := factory.Create()
			order.AddItem(beverage)
			fmt.Printf("Added %s for %d KZT to the order.\n", beverage.Description(), beverage.Cost())
		case 3:
			if len(order.items) > 0 {
				lastItem := order.items[len(order.items)-1]
				milkDecorator := MilkDecorator{beverage: lastItem}
				order.RemoveLastItem()
				order.AddItem(milkDecorator)
				fmt.Printf("Added Milk to %s for %d KZT.\n", lastItem.Description(), lastItem.Cost()+50)
			} else {
				fmt.Println("No items to add milk to.")
			}
		case 4:
			displayOrder(order)
		case 5:
			if len(order.items) > 0 {
				lastItem := order.items[len(order.items)-1]
				fmt.Printf("Removing last item from the order: %s\n", lastItem.Description())
				order.RemoveLastItem()
			} else {
				fmt.Println("No items to remove.")
			}
		case 6:
			totalCost := order.GetTotalCost()
			fmt.Printf("Total: %d KZT. Thank you for your order! Exiting...\n", totalCost)
			return
		case 7:
			order.ClearOrder()
			fmt.Println("Order cleared.")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func displayOrder(order *Order) {
	if len(order.items) == 0 {
		fmt.Println("No items in the order.")
		return
	}

	fmt.Println("Current Order:")
	for _, item := range order.items {
		fmt.Printf("- %s: %d KZT\n", item.Description(), item.Cost())
	}
	fmt.Printf("Total: %d KZT\n", order.GetTotalCost())
}
