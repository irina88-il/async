package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) issueState(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Началась обработка авиарейса №", input.FlightId)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(3 * time.Second)
		sendStateRequest(input)
	}()
}

func sendStateRequest(request models.Request) {
	var state = 0
	if rand.Intn(10)%10 > 2 {
		state = rand.Intn(3)

		fmt.Printf("Авиарейс №%d обработан\n", request.FlightId)
		fmt.Println("Статус авиарейса: ", state)
	} else {
		fmt.Printf("Не удалось обработать авиарейс №%d\n", request.FlightId)
	}

	answer := models.StateRequest{
		AccessToken: 123,
		State:       state,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/flights/%d/update_state/", request.FlightId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
