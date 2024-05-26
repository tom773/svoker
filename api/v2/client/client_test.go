package client

import (
	"encoding/json"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/tom773/casgo/models"
)

func TestClientActualInput(t *testing.T) {
	client := resty.New()
	var apiResponse models.ApiResponse
	resp, err := client.R().Get("http://localhost:8090/api/collections/users/records")
	if resp != nil {
		t.Fatal("Response is nil")
	}
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(resp.Body(), &apiResponse)
	if err != nil {
		t.Fatal(err)
	}
}
func TestMain(m *testing.M) {
	log.Println("Tests are about to start")
	code := m.Run()
	log.Println("Tests are done")
	os.Exit(code)
}

// Actual Example from DB
var user = models.User{
	Avatar:          "cartman_0TTcu3tpAu.jpg",
	Balance:         0,
	CollectionID:    "_pb_users_auth_",
	CollectionName:  "users",
	Created:         "2024-05-08 10:53:29.106Z",
	EmailVisibility: false,
	ID:              "lu5500p6h7qn1z5",
	Name:            "Tom Matthews",
	Updated:         "2024-05-08 10:53:29.106Z",
	Username:        "tom773",
	Verified:        false,
}

func TestClientManualInput(t *testing.T) {
	userResponse := models.ApiResponse{
		Page:       1,
		PerPage:    30,
		TotalItems: 3,
		TotalPages: 1,
		Items:      []models.User{user},
	}

	userResponseJSON, err := json.Marshal(userResponse)
	if err != nil {
		t.Fatalf("Failed to marshal user response: %v", err)
	}

	var unmarshalledUserResponse models.ApiResponse
	if err := json.Unmarshal(userResponseJSON, &unmarshalledUserResponse); err != nil {
		t.Fatalf("Failed to unmarshal user response: %v", err)
	}

	users, ok := unmarshalledUserResponse.Items.([]interface{})
	if !ok {
		t.Fatalf("Expected []interface{}, got %T", unmarshalledUserResponse.Items)
	}

	var convertedUsers []models.User
	for _, item := range users {
		userMap, ok := item.(map[string]interface{})
		if !ok {
			t.Fatalf("Expected map[string]interface{}, got %T", item)
		}
		var user models.User
		userBytes, _ := json.Marshal(userMap)
		json.Unmarshal(userBytes, &user)
		convertedUsers = append(convertedUsers, user)
	}

	if !reflect.DeepEqual(convertedUsers, userResponse.Items) {
		t.Errorf("Expected %v, got %v", userResponse.Items, convertedUsers)
	}

	if reflect.TypeOf(convertedUsers) != reflect.TypeOf([]models.User{}) {
		t.Fatalf("Expected slice, got %v", reflect.TypeOf(unmarshalledUserResponse.Items).Kind())
	}

}
