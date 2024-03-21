package pokemon

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreatePokemonInput struct {
	Num           string           `json:"num"`
	Name          string           `json:"name" binding:"required"`
	Img           string           `json:"img" binding:"required"`
	Type          []string         `json:"type" binding:"required"`
	Height        string           `json:"height" binding:"required"`
	Weight        string           `json:"weight" binding:"required"`
	Candy         string           `json:"candy"`
	Egg           string           `json:"egg"`
	Multipliers   []float64        `json:"multipliers"`
	Weaknesses    []string         `json:"weaknesses"`
	CandyCount    float64          `json:"candy_count"`
	SpawnChance   float64          `json:"spawn_chance" binding:"required"`
	AvgSpawns     float64          `json:"avg_spawns"`
	SpawnTime     string           `json:"spawn_time"`
	NextEvolution []*NextEvolution `json:"next_evolution"`
}

type UpdatePokemonInput struct {
	ID            string           `json:"id" binding:"required"`
	Num           string           `json:"num"`
	Name          string           `json:"name"`
	Img           string           `json:"img"`
	Type          []string         `json:"type"`
	Height        string           `json:"height"`
	Weight        string           `json:"weight"`
	Candy         string           `json:"candy"`
	Egg           string           `json:"egg"`
	Multipliers   []float64        `json:"multipliers"`
	Weaknesses    []string         `json:"weaknesses"`
	CandyCount    float64          `json:"candy_count"`
	SpawnChance   float64          `json:"spawn_chance"`
	AvgSpawns     float64          `json:"avg_spawns"`
	SpawnTime     string           `json:"spawn_time"`
	NextEvolution []*NextEvolution `json:"next_evolution"`
}

type InputNextEvolution struct {
	Num  string `json:"num"`
	Name string `json:"name"`
}

type Pokemon struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	PokemonID     int                `bson:"pokemon_id,omitempty"`
	Num           string             `bson:"num,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Img           string             `bson:"img,omitempty"`
	Type          []string           `bson:"type,omitempty"`
	Height        string             `bson:"height,omitempty"`
	Weight        string             `bson:"weight,omitempty"`
	Candy         string             `bson:"candy,omitempty"`
	Egg           string             `bson:"egg,omitempty"`
	Multipliers   []float64          `bson:"multipliers,omitempty"`
	Weaknesses    []string           `bson:"weaknesses,omitempty"`
	CandyCount    float64            `bson:"candy_count,omitempty"`
	SpawnChance   float64            `bson:"spawn_chance,omitempty"`
	AvgSpawns     float64            `bson:"avg_spawns,omitempty"`
	SpawnTime     string             `bson:"spawn_time,omitempty"`
	NextEvolution []*NextEvolution   `bson:"next_evolution,omitempty"`
}

type NextEvolution struct {
	Num  string `bson:"num,omitempty"`
	Name string `bson:"name,omitempty"`
}
