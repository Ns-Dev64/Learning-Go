package main

import (
	"fmt"
	"math/rand"
	// "strings"
)

type Car struct {
	companyName   string
	licenseNumber string
	mileage       int
	rentFee       int
	carId         int
}
type Customer struct {
	customerName string
	phoneNumber  int
	emailId      string
}
type RentalAgreement struct {
	Customer
	Car
	bookingTime string
	returnTime  string
}

func main() {

	fmt.Println("Welcome to Car Rental System")
	for {
		var choice int
		fmt.Println("\nEnter 1 for Renting a Car and 2 to register your own Car")
		fmt.Scan(&choice)
		if choice == 1 {
			for {
			label1:
			   Customers := []Customer{}

			}

		}
		if choice == 2 {
			for {
			label:
				Cars := []Car{}
				companyName, licenseNumber, rentFee, mileage := inputCars()
				var choices string
				fmt.Scan(&choices)
				if choices == "no" || choices == "No" {
					goto label
				}
				if choices == "yes" || choices == "Yes" {
					Cars := alterCars(Cars, companyName, licenseNumber, rentFee, mileage)
					fmt.Print(Cars)
					break
				}
			}
		}
	}
}

func inputCars() (string, string, int, int) {
	var companyName, licenseNumber string
	var rentFee, mileage int
	fmt.Println("Enter the Car details")
	fmt.Println("Enter the Company Name : ")
	fmt.Scan(&companyName)
	fmt.Println("Enter the License Number : ")
	fmt.Scan(&licenseNumber)
	fmt.Println("Enter the rent fee per hour (20% is taken as commission) : ")
	fmt.Scan(&rentFee)
	fmt.Println("Enter the Car's Mileage (in km/l) : ")
	fmt.Scan(&mileage)
	fmt.Printf("Here's your Car details:\n Company: %v\n Rent: %v\n License Number :%v\n Mileage :%v\n Do you wanna proceed or edit your car details 'yes' to proceed 'no' to edit the details", companyName, rentFee, licenseNumber, mileage)
	return companyName, licenseNumber, rentFee, mileage
}

func alterCars(Cars []Car, companyName string, licenseNumber string, rentFee int, mileage int) []Car {
	newCar := Car{
		companyName:   companyName,
		licenseNumber: licenseNumber,
		rentFee:       rentFee,
		mileage:       mileage,
		carId:         rand.Int(),
	}
	Cars = append(Cars, newCar)
	return Cars
}

func getCustDetails() (string, string, int) {
	var customerName,emailId string
	var phoneNumber  int

}
