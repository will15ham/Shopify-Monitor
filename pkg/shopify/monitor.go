package shopify

import "fmt"

func CompareStockStatus(previous, current map[int64]bool) {
    for id, currentStatus := range current {
        if previousStatus, exists := previous[id]; exists {
            if previousStatus != currentStatus {
                if currentStatus {
                    fmt.Printf("Variant %d is back in stock.\n", id)
                } else {
                    fmt.Printf("Variant %d went out of stock.\n", id)
                }
            }
        } else {
            if currentStatus {
                fmt.Printf("Variant %d is newly in stock.\n", id)
            } else {
                fmt.Printf("Variant %d is newly out of stock.\n", id)
            }
        }
    }
}