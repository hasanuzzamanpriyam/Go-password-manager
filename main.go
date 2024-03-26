package main

import (
	"Password-manager/password"
	"Password-manager/storage"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const filename = "passwords.json"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	fmt.Println("Welcome to Password Manager!")
	var entries []storage.PasswordEntry

	entries, err := storage.LoadPasswordsFromFile(filename)
	if err != nil {
		fmt.Println("Error loading passwords:", err)
	}

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Generate password")
		fmt.Println("2. View saved passwords")
		fmt.Println("3. Delete password")
		fmt.Println("4. Exit")

		fmt.Print(">>> ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			entry := getPasswordEntryFromUser()
			entries = append(entries, entry)

			if err := storage.SavePasswordsToFile(entries, filename); err != nil {
				fmt.Println("Error saving passwords:", err)
			}
		case 2:
			displayPasswords(entries)
		case 3:
			deletePasswords(&entries)
		case 4:
			fmt.Println("Thank you for using Password Manager. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func getPasswordEntryFromUser() storage.PasswordEntry {
	var service, username string
	var length int
	var useUppercase, useDigits, useSpecialChars string

	fmt.Print("Enter service name: ")
	fmt.Scanln(&service)
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password length: ")
	fmt.Scanln(&length)
	fmt.Print("Use uppercase letters? (y/n) ")
	fmt.Scanln(&useUppercase)
	fmt.Print("Use digits? (y/n) ")
	fmt.Scanln(&useDigits)
	fmt.Print("Use special characters? (y/n) ")
	fmt.Scanln(&useSpecialChars)

	password := password.GeneratePassword(length, useUppercase == "y", useDigits == "y", useSpecialChars == "y")
	return storage.PasswordEntry{
		Service:  service,
		Username: username,
		Password: password,
	}
}

func displayPasswords(entries []storage.PasswordEntry) {
	if len(entries) == 0 {
		fmt.Println("No passwords found.")
		return
	}
	fmt.Println("\nSaved passwords:")
	for _, entry := range entries {
		fmt.Printf("Service: %s\nUsername: %s\nPassword: %s\n\n", entry.Service, entry.Username, entry.Password)
	}
}

func deletePasswords(entries *[]storage.PasswordEntry) {
	if len(*entries) == 0 {
		fmt.Println("No passwords to delete")
		return
	}

	var index int
	fmt.Print("Enter the index of the password to delete: ")
	fmt.Scanln(&index)

	if index < 1 || index > len(*entries) {
		fmt.Println("Invalid index")
		return
	}

	*entries = append((*entries)[:index-1], (*entries)[index:]...)

	if err := storage.SavePasswordsToFile(*entries, filename); err != nil {
		fmt.Println("Error saving passwords:", err)
	} else {
		fmt.Println("Password deleted successfully")
	}
}
