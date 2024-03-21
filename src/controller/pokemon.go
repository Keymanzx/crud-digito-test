package controller

import (
	"api-gin/src/models/pokemon"
	repoPoke "api-gin/src/repository/pokemon"

	servicePkm "api-gin/src/services/pokemon"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllPokemon(c *gin.Context) {

	pageIndex := c.Query("pageIndex")
	pageSize := c.Query("pageSize")

	users, err := repoPoke.GetAllPokemon(pageIndex, pageSize)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get Pokemon Error", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get All Pokemon", "data": &users})
	return
}

func GetByPokemonID(c *gin.Context) {
	PokemonID := c.Param("id")

	pokenObj, err := repoPoke.GetByPokemonID(PokemonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Get Pokemon Error", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Get Pokemon Data", "data": pokenObj})
	return
}

func CreatePokemon(c *gin.Context) {
	var importInput pokemon.CreatePokemonInput

	// Bind request body to CreateUserRequest struct and perform validation
	if err := c.ShouldBindJSON(&importInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert Body
	log.Println("importInput:", importInput)
	bodyUser := servicePkm.MapBodyCreatePokemon(importInput)

	// Save User In DB
	err := repoPoke.CreatePokemon(bodyUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Create Pokemon Error", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Create Pokemon Success", "data": nil})
	return
}

func UpdatePokemon(c *gin.Context) {
	var importInput pokemon.UpdatePokemonInput

	// Bind request body to CreateUserRequest struct and perform validation
	if err := c.ShouldBindJSON(&importInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert Body
	log.Println("importInput:", importInput)
	bodyUser := servicePkm.MapBodyUpdatePokemon(importInput)

	// Save User In DB
	err := repoPoke.UpdateByID(importInput.ID, bodyUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Update Pokemon Error", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Update Pokemon Success", "data": nil})
	return
}

func DeleteByPokemonID(c *gin.Context) {
	pkmID := c.Param("id")

	// Delete User In DB
	err := repoPoke.DeleteByID(pkmID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Delete Pokemon Error", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete Pokemon Success", "id": pkmID})
	return
}
