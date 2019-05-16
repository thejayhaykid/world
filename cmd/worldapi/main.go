package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ironarachne/world/pkg/character"
	"github.com/ironarachne/world/pkg/climate"
	"github.com/ironarachne/world/pkg/country"
	"github.com/ironarachne/world/pkg/culture"
	"github.com/ironarachne/world/pkg/heavens"
	"github.com/ironarachne/world/pkg/heraldry"
	"github.com/ironarachne/world/pkg/language"
	"github.com/ironarachne/world/pkg/organization"
	"github.com/ironarachne/world/pkg/pantheon"
	"github.com/ironarachne/world/pkg/random"
	"github.com/ironarachne/world/pkg/region"
	"github.com/ironarachne/world/pkg/town"
	"github.com/ironarachne/world/pkg/world"
	"github.com/ironarachne/world/pkg/worldmap"
)

func getCharacter(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o character.Character

	random.SeedFromString(id)

	o = character.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCharacterRandom(w http.ResponseWriter, r *http.Request) {
	var o character.Character

	rand.Seed(time.Now().UnixNano())

	o = character.Generate()

	json.NewEncoder(w).Encode(o)
}

func getClimate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o climate.Climate

	random.SeedFromString(id)

	o = climate.Generate()

	json.NewEncoder(w).Encode(o)
}

func getClimateRandom(w http.ResponseWriter, r *http.Request) {
	var o climate.Climate

	rand.Seed(time.Now().UnixNano())

	o = climate.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCountry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o country.Country

	random.SeedFromString(id)

	o = country.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCountryRandom(w http.ResponseWriter, r *http.Request) {
	var o country.Country

	rand.Seed(time.Now().UnixNano())

	o = country.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCulture(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o culture.Culture

	random.SeedFromString(id)

	o = culture.Generate()

	json.NewEncoder(w).Encode(o)
}

func getCultureRandom(w http.ResponseWriter, r *http.Request) {
	var o culture.Culture

	rand.Seed(time.Now().UnixNano())

	o = culture.Generate()

	json.NewEncoder(w).Encode(o)
}

func getHeavens(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o heavens.Heavens

	random.SeedFromString(id)

	o = heavens.Generate()

	json.NewEncoder(w).Encode(o)
}

func getHeavensRandom(w http.ResponseWriter, r *http.Request) {
	var o heavens.Heavens

	rand.Seed(time.Now().UnixNano())

	o = heavens.Generate()

	json.NewEncoder(w).Encode(o)
}

func getHeraldry(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o heraldry.Heraldry

	random.SeedFromString(id)

	o = heraldry.GenerateHeraldry()

	json.NewEncoder(w).Encode(o)
}

func getHeraldryRandom(w http.ResponseWriter, r *http.Request) {
	var o heraldry.Heraldry

	rand.Seed(time.Now().UnixNano())

	o = heraldry.GenerateHeraldry()

	json.NewEncoder(w).Encode(o)
}

func getLanguage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o language.Language

	random.SeedFromString(id)

	o = language.Generate()

	json.NewEncoder(w).Encode(o)
}

func getLanguageRandom(w http.ResponseWriter, r *http.Request) {
	var o language.Language

	rand.Seed(time.Now().UnixNano())

	o = language.Generate()

	json.NewEncoder(w).Encode(o)
}

func getOrganization(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o organization.Organization

	random.SeedFromString(id)

	o = organization.Generate()

	json.NewEncoder(w).Encode(o)
}

func getOrganizationRandom(w http.ResponseWriter, r *http.Request) {
	var o organization.Organization

	rand.Seed(time.Now().UnixNano())

	o = organization.Generate()

	json.NewEncoder(w).Encode(o)
}

func getPantheon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o pantheon.Pantheon
	var l language.Language

	random.SeedFromString(id)

	l = language.Generate()
	o = pantheon.Generate(15, l)

	json.NewEncoder(w).Encode(o)
}

func getPantheonRandom(w http.ResponseWriter, r *http.Request) {
	var o pantheon.Pantheon
	var l language.Language

	rand.Seed(time.Now().UnixNano())

	l = language.Generate()
	o = pantheon.Generate(15, l)

	json.NewEncoder(w).Encode(o)
}

func getRegion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o region.Region

	random.SeedFromString(id)

	o = region.Generate("random")

	json.NewEncoder(w).Encode(o)
}

func getRegionRandom(w http.ResponseWriter, r *http.Request) {
	var o region.Region

	rand.Seed(time.Now().UnixNano())

	o = region.Generate("random")

	json.NewEncoder(w).Encode(o)
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	o := "This is the World Generation API."

	json.NewEncoder(w).Encode(o)
}

func getTown(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o town.Town

	random.SeedFromString(id)

	o = town.Generate("random", "random")

	json.NewEncoder(w).Encode(o)
}

func getTownRandom(w http.ResponseWriter, r *http.Request) {
	var o town.Town

	rand.Seed(time.Now().UnixNano())

	o = town.Generate("random", "random")

	json.NewEncoder(w).Encode(o)
}

func getWorld(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var o world.World

	random.SeedFromString(id)

	o = world.Generate()

	json.NewEncoder(w).Encode(o)
}

func getWorldRandom(w http.ResponseWriter, r *http.Request) {
	var o world.World

	rand.Seed(time.Now().UnixNano())

	o = world.Generate()

	json.NewEncoder(w).Encode(o)
}

func getWorldMap(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var m worldmap.WorldMap

	random.SeedFromString(id)

	m = worldmap.Generate(60, 80)

	json.NewEncoder(w).Encode(m)
}

func getWorldMapSVGImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var m worldmap.WorldMap

	random.SeedFromString(id)

	m = worldmap.Generate(60, 80)
	o := m.RenderAsSVG()

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(o))
}

func getWorldMapTextImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var m worldmap.WorldMap

	random.SeedFromString(id)

	m = worldmap.Generate(60, 80)
	o := m.RenderAsText()

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(o))
}

func getWorldMapRandom(w http.ResponseWriter, r *http.Request) {
	var m worldmap.WorldMap

	rand.Seed(time.Now().UnixNano())

	m = worldmap.Generate(60, 80)

	json.NewEncoder(w).Encode(m)
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/character", getCharacterRandom)
	r.Get("/character/{id}", getCharacter)

	r.Get("/climate", getClimateRandom)
	r.Get("/climate/{id}", getClimate)

	r.Get("/country", getCountryRandom)
	r.Get("/country/{id}", getCountry)

	r.Get("/culture", getCultureRandom)
	r.Get("/culture/{id}", getCulture)

	r.Get("/heavens", getHeavensRandom)
	r.Get("/heavens/{id}", getHeavens)

	r.Get("/heraldry", getHeraldryRandom)
	r.Get("/heraldry/{id}", getHeraldry)

	r.Get("/language", getLanguageRandom)
	r.Get("/language/{id}", getLanguage)

	r.Get("/organization", getOrganizationRandom)
	r.Get("/organization/{id}", getOrganization)

	r.Get("/pantheon", getPantheonRandom)
	r.Get("/pantheon/{id}", getPantheon)

	r.Get("/region", getRegionRandom)
	r.Get("/region/{id}", getRegion)

	r.Get("/town", getTownRandom)
	r.Get("/town/{id}", getTown)

	r.Get("/world", getWorldRandom)
	r.Get("/world/{id}", getWorld)
	r.Get("/world/{id}/map", getWorldMapSVGImage)

	r.Get("/worldmap", getWorldMapRandom)
	r.Get("/worldmap/{id}", getWorldMap)
	r.Get("/worldmap/{id}/image", getWorldMapSVGImage)
	r.Get("/worldmap/{id}/textimage", getWorldMapTextImage)

	r.Get("/", getRoot)

	fmt.Println("World Generator API is online.")
	log.Fatal(http.ListenAndServe(":7531", r))
}
