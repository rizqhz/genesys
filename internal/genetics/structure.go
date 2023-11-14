package main

import "math/rand"

type Genotype [24][6]int

type Chromosome struct {
	genes   Genotype
	fitness float64
}

func NewChromosome(random *rand.Rand) *Chromosome {
	chromosome := new(Chromosome)
	for i := range chromosome.genes {
		chromosome.genes[i][0] = random.Intn(3)
		chromosome.genes[i][1] = random.Intn(4)
		chromosome.genes[i][2] = random.Intn(6)
		chromosome.genes[i][3] = random.Intn(3)
		chromosome.genes[i][4] = random.Intn(6)
		chromosome.genes[i][5] = random.Intn(6)
	}
	return chromosome
}

type Population []*Chromosome

func NewPopulation(size int, random *rand.Rand) Population {
	population := make(Population, size)
	if random != nil {
		for i := range population {
			population[i] = NewChromosome(random)
		}
	}
	return population
}
