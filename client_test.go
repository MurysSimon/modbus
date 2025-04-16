package modbus

import (
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"
)

func TestNewReadRegsResponse(t *testing.T) {
	var client *ModbusClient
	var err error
	unitID := 1

	client, err = NewClient(&ClientConfiguration{
		URL:     "rtuovertcp://127.0.0.100:502",
		Speed:   9600,
		Timeout: 1 * time.Second,
	})

	client.SetUnitId(uint8(unitID))

	err = client.Open()
	if err != nil {
		fmt.Println("modbus connection has failed")
		return
	}

	time.Sleep(time.Second * 2)

	results, txRaw, rxRaw, err := client.ReadRegister(uint16(0), HOLDING_REGISTER)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data type of register: %v\n", reflect.TypeOf(results))
	fmt.Printf("Value of register: %d\n", results)
	fmt.Printf("Raw request: %x\n", txRaw)
	fmt.Printf("Raw response: %x\n", rxRaw)

	client.Close()
}

func TestTCPNewReadCoilsResponse(t *testing.T) {
	var client *ModbusClient
	var err error
	unitID := 2

	client, err = NewClient(&ClientConfiguration{
		URL:     "tcp://127.0.0.100:502",
		Timeout: 1 * time.Second,
	})

	client.SetUnitId(uint8(unitID))

	err = client.Open()
	if err != nil {
		fmt.Println("modbus connection has failed")
		return
	}

	time.Sleep(time.Second * 2)

	results, txRaw, rxRaw, err := client.ReadCoils(uint16(0), uint16(20))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data type of register: %v\n", reflect.TypeOf(results))
	fmt.Printf("Value of register: %t\n", results)
	fmt.Printf("Raw request: %x\n", txRaw)
	fmt.Printf("Raw response: %x\n", rxRaw)

	client.Close()
}

func TestRTUNewReadCoilsResponse(t *testing.T) {
	var client *ModbusClient
	var err error
	unitID := 2

	client, err = NewClient(&ClientConfiguration{
		URL:     "rtuovertcp://127.0.0.100:502",
		Speed:   9600,
		Timeout: 1 * time.Second,
	})

	client.SetUnitId(uint8(unitID))

	err = client.Open()
	if err != nil {
		fmt.Println("modbus connection has failed")
		return
	}

	time.Sleep(time.Second * 2)

	results, txRaw, rxRaw, err := client.ReadCoils(uint16(0), uint16(20))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data type of register: %v\n", reflect.TypeOf(results))
	fmt.Printf("Value of register: %t\n", results)
	fmt.Printf("Raw request: %x\n", txRaw)
	fmt.Printf("Raw response: %x\n", rxRaw)

	client.Close()
}
