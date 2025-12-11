package main

import "fmt"

type Contact struct {
    Name  string
    Phone string
    Email string
}

func main() {
    
    contacts := []Contact{
        {Name: "Alice", Phone: "07700123", Email: "alice@mail.com"},
        {Name: "Bob", Phone: "07900456", Email: "bob@mail.com"},
        {Name: "Charlie", Phone: "07900789", Email: "charlie@mail.com"},
    }
    
   
    fmt.Println("All Contacts:")
    for _, contact := range contacts {
        fmt.Printf("Name: %s, Phone: %s, Email: %s\n", 
            contact.Name, contact.Phone, contact.Email)
    }
    
   
    searchName := "Bob"
    fmt.Printf("\nSearching for: %s\n", searchName)
    
    
    for _, contact := range contacts {
        if contact.Name == searchName {
            fmt.Println("Found!")
            fmt.Printf("Name: %s\n", contact.Name)
            fmt.Printf("Phone: %s\n", contact.Phone)
            fmt.Printf("Email: %s\n", contact.Email)
            break  // Stop searching once we find it
        }
    }
}
