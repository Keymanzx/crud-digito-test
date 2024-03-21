package pokemon

import (
	"api-gin/src/models/pokemon"
	repoPoke "api-gin/src/repository/pokemon"
)

func MapBodyCreatePokemon(dataInput pokemon.CreatePokemonInput) *pokemon.Pokemon {

	countPkm, _ := repoPoke.GetAllPokemon("", "")

	return &pokemon.Pokemon{

		PokemonID:     len(countPkm) + 1,
		Num:           dataInput.Num,
		Name:          dataInput.Name,
		Img:           dataInput.Img,
		Type:          dataInput.Type,
		Height:        dataInput.Height,
		Weight:        dataInput.Weight,
		Candy:         dataInput.Candy,
		CandyCount:    dataInput.CandyCount,
		Egg:           dataInput.Candy,
		Multipliers:   dataInput.Multipliers,
		Weaknesses:    dataInput.Weaknesses,
		SpawnChance:   dataInput.SpawnChance,
		AvgSpawns:     dataInput.AvgSpawns,
		SpawnTime:     dataInput.SpawnTime,
		NextEvolution: dataInput.NextEvolution,
	}
}

func MapBodyUpdatePokemon(dataInput pokemon.UpdatePokemonInput) *pokemon.Pokemon {
	return &pokemon.Pokemon{
		Num:           dataInput.Num,
		Name:          dataInput.Name,
		Img:           dataInput.Img,
		Type:          dataInput.Type,
		Height:        dataInput.Height,
		Weight:        dataInput.Weight,
		Candy:         dataInput.Candy,
		CandyCount:    dataInput.CandyCount,
		Egg:           dataInput.Candy,
		Multipliers:   dataInput.Multipliers,
		Weaknesses:    dataInput.Weaknesses,
		SpawnChance:   dataInput.SpawnChance,
		AvgSpawns:     dataInput.AvgSpawns,
		SpawnTime:     dataInput.SpawnTime,
		NextEvolution: dataInput.NextEvolution,
	}
}
