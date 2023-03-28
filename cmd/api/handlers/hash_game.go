package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zthiagovalle/jogo_da_velha/internal/entity"
)

func RunHashGame(c *fiber.Ctx) error {
	hashGame := entity.NewHashGame()
	err := json.Unmarshal(c.Body(), hashGame)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	jsonData, err := json.Marshal(hashGame)
	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	hashGameString := string(jsonData)
	err = ChatGpt(hashGameString, hashGame)
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	return c.Status(200).JSON(hashGame)
}

func ChatGpt(message string, hashGame *entity.HashGame) error {
	query := fmt.Sprint("Vamos jogar jogo da velha\nO objetivo do jogo é conseguir três marcas iguais em uma linha horizontal, vertical ou diagonal, ou preencher todo o tabuleiro sem que nenhum jogador consiga três marcas em linha.\nVoce é a string “X”\nA matriz está:\n", message, "\nSua vez de jogar, me responda apenas com o json da resposta")
	req := entity.Request{
		Model: "gpt-4",
		Messages: []entity.Message{
			{
				Role:    "user",
				Content: query,
			},
		},
		MaxTokens: 1000,
	}

	reqJson, err := json.Marshal(req)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqJson))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer APIKEY")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var resp entity.Response
	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), hashGame)
	if err != nil {
		return err
	}

	return nil
}
