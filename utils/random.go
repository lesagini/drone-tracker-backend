package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

type Details struct {
	FarmCode                     string
	FarmCoordinates              string
	FarmAirspace                 string
	FarmLocation                 string
	FarmGeolocation              string
	FarmContact                  int64
	OperatorID                   string
	OperatorHeadquater           string
	OperatorName                 string
	OperatorInitials             string
	OperatorNumberPilotsDeployed int32
	OpertorContact               int64
	PilotID                      string
	PilotOperatorID              string
	PilotInitials                string
	PilotNumber                  int32
	PilotFullName                string
	PilotLicenseNumber           int64
	PilotFarmLocationCode        string
	PilotFarmLocation            string
	FlightDuration               float32
	FlightAcreage                float32
	VarietyInternalIdentity      string
	VarietyBotanicalName         string
	VarietyAcreage               float32
	VarietyItervalCode           string
}

func GetDetails() Details {

	gofakeit.Seed(0) // Optional: Set the seed for reproducibility

	// Generate random values for the fields
	farmCode := fmt.Sprintf("%03d-%03d", gofakeit.Number(1, 100), gofakeit.Number(1, 100))
	farmCoordinates := fmt.Sprintf("%f,%f", gofakeit.Latitude(), gofakeit.Longitude())
	farmAirspace := gofakeit.RandomString([]string{"HKP4", "HKP3", "ARP", "HDD1", "HPQ3", "APP#"})
	farmLocation := gofakeit.City() + ", " + gofakeit.Country()
	farmGeolocation := "6CHM+696, Karagita"
	farmContact, _ := strconv.Atoi(gofakeit.Phone())
	operatorHeadquater := gofakeit.City()
	operatorName := gofakeit.Company() + gofakeit.Company()
	operatorInitials := GetNameInitials(operatorName)
	operatorID := operatorInitials
	operatorNumberPilotsDeployed := gofakeit.Number(1, 20)
	opertorContact, _ := strconv.Atoi(gofakeit.Phone())
	pilotID := gofakeit.UUID()
	pilotOperatorID := operatorID
	pilotFullName := gofakeit.Name()
	pilotInitials := GetNameInitials(pilotFullName)
	pilotNumber := gofakeit.Number(1, 100)
	pilotLicenseNumber, _ := strconv.ParseInt(gofakeit.Numerify("##########"), 10, 64)
	pilotFarmLocationCode := farmCode
	//pilotFarmLocationCode := gofakeit.RandomString([]string{"A", "B", "C", "D", "E"}) + "-" + gofakeit.Number(1, 999)
	pilotFarmLocation := farmLocation
	flightDuration := gofakeit.Float32Range(1.0, 8.0)
	flightAcreage := flightDuration * gofakeit.Float32Range(0.8, 1.2) * 2000

	variety_botanical_name := gofakeit.Color() + " " + gofakeit.Fruit()
	variety_internal_identity := GetNameInitials(variety_botanical_name)
	variety_acreage := gofakeit.Float32Range(10.0, 50.0)
	variety_iterval_code := fmt.Sprintf("%07b", rand.Intn(128))

	details := Details{
		farmCode,
		farmCoordinates,
		farmAirspace,
		farmLocation,
		farmGeolocation,
		int64(farmContact),
		operatorID,
		operatorHeadquater,
		operatorName,
		operatorInitials,
		int32(operatorNumberPilotsDeployed),
		int64(opertorContact),
		pilotID,
		pilotOperatorID,
		pilotInitials,
		int32(pilotNumber),
		pilotFullName,
		pilotLicenseNumber,
		pilotFarmLocationCode,
		pilotFarmLocation,
		flightDuration,
		flightAcreage,
		variety_internal_identity,
		variety_botanical_name,
		variety_acreage,
		variety_iterval_code,
	}

	return details
}

func GetNameInitials(name string) string {
	initials := ""

	for _, v := range strings.Split(name, " ") {
		initials += string(v[0])
	}

	return initials
}
