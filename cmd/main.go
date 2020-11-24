package main

import (
	"fpl-strategy-tester/internal"
	"log"
	"math/rand"
	"time"
)

func main() {

	resolver := internal.NewResolver()
	resolver.ResolveDatabase()
	resolver.ResolveCache()

	// Set the seed used for generating random numbers
	rand.Seed(time.Now().UnixNano())

	// Simulate the teams used to feed into the different FPL strategies
	log.Printf("-> Simulating 10,000 random FPL teams...\t")
	simulatedTeams, errs := resolver.GenerateTeams()
	for err := range errs {
		log.Printf("Error: %v\n", err)
		return
	}

	// Run the cost variation strategy
	log.Printf("-> Running Cost Variation strategy...\t")
	if err := resolver.RunCostVariationStrategy(simulatedTeams); err != nil {
		log.Printf("Error: %v\n", err)
	}

	// Run the distribution strategy
	//if err := resolver.RunDistributionStrategy(); err != nil {
	//	fmt.Printf("Error while running distribution strategy: %v\n", err)
	//}
}
