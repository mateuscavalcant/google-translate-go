package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type TranslationResponse struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func Translate(c *gin.Context) {
	content := strings.TrimSpace(c.PostForm("content"))
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	payload := strings.NewReader(fmt.Sprintf("q=%s&target=es&source=en", content))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		fmt.Println("Error creando la solicitud:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la solicitud"})
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "API_KEY")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}
	defer resp.Body.Close()

	// Leer el cuerpo de la respuesta
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	var translationResp TranslationResponse
	if err := json.Unmarshal(body, &translationResp); err != nil {
		fmt.Println("Error al analizar la respuesta:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al analizar la respuesta"})
		return
	}

	if translationResp.Error.Code != 0 {
		fmt.Println("", translationResp.Error.Message)
		c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		return
	}

	// Retornar a tradução diretamente em JSON
	if len(translationResp.Data.Translations) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "successfully", "translated": translationResp.Data.Translations[0].TranslatedText})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "Nenhuma tradução encontrada."})
	}
}
