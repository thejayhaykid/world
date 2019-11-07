package climate

import (
	"github.com/ironarachne/world/pkg/species"
	"github.com/ironarachne/world/pkg/tree"
)

func (gen Generator) getTrees(humidity int, temperature int) []species.Species {
	trees := tree.All()
	trees = species.FilterHumidity(humidity, trees)
	trees = species.FilterTemperature(temperature, trees)
	trees = species.ByTagIn(gen.TreeTags, trees)

	return trees
}
