package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type FuelType string

const (
	fuelTypeGas      FuelType = "gas"
	fuelTypeDiesel            = "diesel"
	fuelTypeLpg               = "lpg"
	fuelTypeElectric          = "ele"
)

var FuelTypes = []FuelType{fuelTypeGas, fuelTypeDiesel, fuelTypeLpg, fuelTypeElectric}

type car struct {
	id          int
	fuelType    FuelType
	createdTime time.Time
}

type stationConfig struct {
	num              int
	minWait, maxWait time.Duration
}

var fuelStations = map[FuelType]*stationConfig{
	fuelTypeGas:      {num: 4, minWait: time.Second, maxWait: 5 * time.Second},
	fuelTypeDiesel:   {num: 4, minWait: time.Second, maxWait: 5 * time.Second},
	fuelTypeLpg:      {num: 1, minWait: 500 * time.Millisecond, maxWait: 5 * time.Second},
	fuelTypeElectric: {num: 8, minWait: 1500 * time.Millisecond, maxWait: 10 * time.Second},
}

var cashRegisters = &stationConfig{num: 2, minWait: 500 * time.Millisecond, maxWait: 2 * time.Second}

const (
	numCars        = 100
	newCarMaxDelay = 250 * time.Millisecond
	newCarMaxWait  = 250 * time.Millisecond

	stationQueueSize = 10
)

func main() {
	fuelQueues := map[FuelType]chan *car{}
	cashRegisterQueue := make(chan *car)
	outQueue := make(chan *car)

	fuelStationWg := &sync.WaitGroup{}
	fuelStationWg.Add(len(fuelStations))

	// run fuel stations
	for ft, c := range fuelStations {
		ch := make(chan *car, stationQueueSize)
		fuelQueues[ft] = ch
		go func(c *stationConfig) {
			runStation(ch, cashRegisterQueue, c)
			fuelStationWg.Done()
		}(c)
	}

	// close register queue when all fuel stations done
	go func() {
		fuelStationWg.Wait()
		close(cashRegisterQueue)
	}()

	// run cash registers
	go func() {
		runStation(cashRegisterQueue, outQueue, cashRegisters)
		close(outQueue)
	}()

	// simulate arriving cars
	go func() {
		wg := sync.WaitGroup{}
		wg.Add(numCars)

		for i := 0; i < numCars; i++ {
			go func() {
				defer wg.Done()

				c := &car{
					id:          i,
					fuelType:    FuelTypes[rand.Intn(len(FuelTypes))],
					createdTime: time.Now(),
				}

				select {
				case fuelQueues[c.fuelType] <- c:
				case <-time.After(newCarMaxWait):
					// drop car
				}
			}()

			time.Sleep(time.Duration(rand.Int63n(int64(newCarMaxDelay))))
		}

		wg.Wait()
		for _, c := range fuelQueues {
			close(c)
		}
	}()

	numServed := 0

	for car := range outQueue {
		numServed++
		fmt.Println("car", car.id, "fuel", car.fuelType, "delay", time.Now().Sub(car.createdTime))
	}

	fmt.Println(fmt.Sprintf("served %d out of %d", numServed, numCars))
}

func runStation(in chan *car, out chan *car, cfg *stationConfig) {
	wg := &sync.WaitGroup{}
	wg.Add(cfg.num)

	for i := 0; i < cfg.num; i++ {
		go func() {
			defer wg.Done()
			for car := range in {
				time.Sleep(cfg.minWait + time.Duration(rand.Int63n(int64(cfg.maxWait-cfg.minWait))))
				out <- car
			}
		}()
	}

	wg.Wait()
}
