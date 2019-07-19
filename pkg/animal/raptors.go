package animal

import (
	"github.com/ironarachne/world/pkg/resource"
	"github.com/ironarachne/world/pkg/size"
)

func getRaptors() []Animal {
	birds := []Animal{
		Animal{
			Name:           "eagle",
			PluralName:     "eagles",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "eagle feathers",
					Origin:      "eagle",
					Type:        "feather",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "falcon",
			PluralName:     "falcons",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "falcon feathers",
					Origin:      "falcon",
					Type:        "feather",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
		Animal{
			Name:           "hawk",
			PluralName:     "hawks",
			EatsAnimals:    true,
			EatsPlants:     false,
			IsAquatic:      false,
			IsMount:        false,
			IsPackAnimal:   false,
			IsScavenger:    true,
			MinHumidity:    0,
			MaxHumidity:    10,
			MinTemperature: 0,
			MaxTemperature: 10,
			Resources: []resource.Resource{
				resource.Resource{
					Name:        "hawk feathers",
					Origin:      "hawk",
					Type:        "feather",
					Commonality: 4,
				},
			},
			Size: size.GetCategoryByName("tiny"),
		},
	}

	return birds
}
