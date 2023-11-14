package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// | Asisten | Praktikum | Kelas | Ruangan | Hari | Waktu |
// | 0       | 1         | 2     | 3       | 4    | 5     |

func Evaluate(x *Chromosome) {
	for i := 0; i < len(x.genes); i++ {
		var check bool = true
		for y := i + 1; y < len(x.genes); y++ {
			check = true
			check = check && x.genes[i][3] == x.genes[y][3]
			check = check && x.genes[i][4] == x.genes[y][4]
			check = check && x.genes[i][5] == x.genes[y][5]
			if check {
				x.fitness += 1
			}
			check = true
			check = check && x.genes[i][0] == x.genes[y][0]
			check = check && x.genes[i][4] == x.genes[y][4]
			check = check && x.genes[i][5] == x.genes[y][5]
			if check {
				x.fitness += 1
			}
		}
		for kelas, jadwal := range Constraint {
			if Kelas[x.genes[i][2]] == kelas {
				for hari, waktu := range jadwal {
					if Hari[x.genes[i][4]] == hari {
						for i := range waktu {
							src := Durasi[x.genes[i][5]]
							dst := waktu[i]
							if src.collision(dst) {
								x.fitness += 1
							}
						}
					}
				}
			}
		}
	}
	formula := 1 / (x.fitness + 1)
	x.fitness = formula
}

func Crossover(random *rand.Rand, x, y *Chromosome) (*Chromosome, *Chromosome) {
	point := random.Intn(len(x.genes))
	a := &Chromosome{}
	b := &Chromosome{}
	copy(a.genes[:point], x.genes[:point])
	copy(a.genes[point:], y.genes[:point])
	copy(b.genes[:point], x.genes[point:])
	copy(b.genes[point:], y.genes[point:])
	for i := 0; i < len(x.genes); i++ {
		point := random.Intn(len(x.genes[0]))
		copy(a.genes[i][point:], y.genes[i][point:])
		copy(b.genes[i][point:], x.genes[i][point:])
	}
	return a, b
}

func Mutate(random *rand.Rand, x *Chromosome, rate float64) {
	for i := range x.genes {
		if random.Float64() <= rate {
			x.genes[i][0] = 2 - x.genes[i][0]
			x.genes[i][1] = 3 - x.genes[i][1]
			x.genes[i][2] = 5 - x.genes[i][2]
			x.genes[i][3] = 2 - x.genes[i][3]
			x.genes[i][4] = 5 - x.genes[i][4]
			x.genes[i][5] = 5 - x.genes[i][5]
		}
	}
}

func Selection(population Population, n int) Population {
	for i := range population {
		Evaluate(population[i])
	}
	sort.Slice(population, func(x, y int) bool {
		return population[x].fitness > population[y].fitness
	})
	return population[:n]
}

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	population := NewPopulation(100, random)
	generasi := 1
	for {
		parents := Selection(population, 50)
		if population[0].fitness == 1 {
			break
		}
		new_population := NewPopulation(100, nil)
		copy(new_population[50:], parents)
		for i := 0; i < 50; i += 2 {
			p1, p2 := parents[i], parents[i+1]
			c1, c2 := Crossover(random, p1, p2)
			Mutate(random, c1, 0.01)
			Mutate(random, c2, 0.01)
			new_population[i] = c1
			new_population[i+1] = c2
		}
		population = new_population
		generasi++
	}
	best_chromosome := population[0]
	best_fitness := population[0].fitness
	fmt.Println(best_chromosome)
	fmt.Printf("\ngen(%d): %.2f%%\n", generasi, best_fitness*100)
}
