package models

import (
	"context"
	"drones/utils"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

var d utils.Details = utils.GetDetails()

func createRandomFarm(t *testing.T) Farm{
	args := CreateFarmParams{
		FarmCode:        d.FarmCode,
		FarmCoordinates: d.FarmCoordinates,
		FarmAirspace:    d.FarmAirspace,
		FarmLocation:    d.FarmLocation,
		FarmGeolocation: d.FarmGeolocation,
		FarmContact:     d.FarmContact,
	}
	farm, err := testingQueries.CreateFarm(context.Background(), args)

	require.NotEmpty(t, farm)
	require.NoError(t, err)
	return farm
}
func TestCreateFarm(t *testing.T) {
	createRandomFarm(t)
}

func createRandomOperator(t *testing.T) Operator{
	args := CreateOperatorParams{
		OperatorID:                   d.OperatorID,
		OperatorHeadquater:           d.OperatorHeadquater,
		OperatorName:                 d.OperatorName,
		OperatorNumberPilotsDeployed: d.OperatorNumberPilotsDeployed,
		OpertorContact:               d.OpertorContact,
	}

	operator, err := testingQueries.CreateOperator(context.Background(), args)

	require.NotEmpty(t, operator)
	require.NoError(t, err)

	return operator

}

func TestCreateOperator(t *testing.T) {
	createRandomOperator(t)

}

func createRandomPilot(t *testing.T) Pilot {
	pilotid := d.OperatorInitials + "-" + d.PilotInitials + "-" + strconv.Itoa(int(d.PilotNumber))
	args := CreatePilotParams{
		PilotID:               pilotid,
		PilotOperatorID:       d.OperatorInitials,
		PilotInitials:         d.PilotInitials,
		PilotNumber:           d.PilotNumber,
		PilotFullName:         d.PilotFullName,
		PilotLicenseNumber:    d.PilotLicenseNumber,
		PilotFarmLocationCode: d.FarmCode,
		PilotFarmLocation:     d.FarmLocation,
		PilotStatus:           PilotStatusesActive,
		PilotClassification:   PilotClassificationsA,
		PilotFlightHours:      "0.0",
		PilotCoveredAcreage:   "0.0",
	}
	pilot, err := testingQueries.CreatePilot(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, pilot)

	require.Equal(t, args.PilotID, pilot.PilotID)
	require.Equal(t, args.PilotFullName, pilot.PilotFullName)

	require.NotZero(t, pilot.ID)
	require.NotZero(t, pilot.PilotEntryDate)

	return pilot
}
func TestCreatePilot(t *testing.T) {
	createRandomPilot(t)

}

func createRandomFlight(t *testing.T) Flight {
	pilotid := d.OperatorInitials + "-" + d.PilotInitials + "-" + strconv.Itoa(int(d.PilotNumber))
	args := CreateFlightParams{
		FlightFarmLocation:    d.FarmLocation,
		FlightFarmGeolocation: d.FarmGeolocation,
		FlightFarmID:          d.FarmCode,
		FlightPilot:           pilotid,
		FlightDuration:        fmt.Sprintf("%f", d.FlightDuration),
		FlightAcreage:         fmt.Sprintf("%f", d.FlightAcreage),
	}
	flight, err := testingQueries.CreateFlight(context.Background(), args)

	require.NotEmpty(t, flight)
	require.NoError(t, err)

	return flight

}

func TestCreateFlight(t *testing.T) {
	createRandomFlight(t)
	

}

func TestCreateVariety(t *testing.T) {
	args := CreateVarietyParams{
		VarietyInternalIdentity: d.VarietyInternalIdentity,
		VarietyBotanicalName:    d.VarietyBotanicalName,
		VarietyFarmID:           d.FarmCode,
		VarietyAcreage:          int64(d.VarietyAcreage),
		VarietyType:             VarietyTypesContinuous,
		VarietyItervalCode:      d.VarietyItervalCode,
	}

	variety, err := testingQueries.CreateVariety(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, variety)

}

