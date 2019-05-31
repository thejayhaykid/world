package race

import (
	"math/rand"

	"github.com/ironarachne/world/pkg/random"
)

// Race is a race. It can have variations within.
type Race struct {
	Name       string
	Adjective  string
	AgeCategories []AgeCategory
	Appearance Appearance
	Commonality int
}

// GenerateSubrace generates a random subrace based on a parent race
func GenerateSubrace(parent Race) Race {
	race := parent

	race.Appearance = parent.generateRandomSubraceAppearance()

	return race
}

// Get returns a specific race
func Get(name string) Race {
	race := Race{}
	races := getAllRaces()

	for _, r := range races {
		if r.Name == name {
			return r
		}
	}

	return race
}

// GetRandom returns a random race from the list
func GetRandom() Race {
	races := getAllRaces()

	return races[rand.Intn(len(races))]
}

// GetRandomWeighted returns a random race, taking commonality into account
func GetRandomWeighted() Race {
	races := getAllRaces()

	weights := map[string]int{}

	for _, c := range races {
		weights[c.Name] = c.Commonality
	}

	name := random.StringFromThresholdMap(weights)

	for _, c := range races {
		if c.Name == name {
			return c
		}
	}

	return Race{}
}