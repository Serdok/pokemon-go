package health

import (
	"encoding/json"
	"github.com/Serdok/pokemon-go/internal/api/mocking"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	// Create context
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Call
	mocking.MockBody(c, "GET", nil)
	GetHealth(c)

	// Check return codes
	var body interface{}
	err := json.Unmarshal(w.Body.Bytes(), &body)
	if err != nil {
		t.Fatal(err)
	}
	if w.Code != 200 {
		t.Fatalf("Health -- get: Got %v", body)
	}

	// Since the body contents may vary if the return code is 200, we do not check the response text
}
