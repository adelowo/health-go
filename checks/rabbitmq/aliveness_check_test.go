package rabbitmq

import (
	"os"
	"testing"

	"github.com/hellofresh/health-go/checks/http"
)

const httpURLEnv = "HEALTH_GO_MQ_URL"

func TestAliveness(t *testing.T) {
	if os.Getenv(httpURLEnv) == "" {
		t.SkipNow()
	}

	check := http.New(http.Config{
		URL: os.Getenv(httpURLEnv),
	})

	if err := check(); err != nil {
		t.Fatalf("HTTP check failed: %s", err.Error())
	}
}
