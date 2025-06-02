// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vernsdk_test

import (
	"context"
	"os"
	"testing"

	"github.com/vern-so/sdk-go"
	"github.com/vern-so/sdk-go/internal/testutil"
	"github.com/vern-so/sdk-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := vernsdk.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	run, err := client.Runs.New(context.TODO(), vernsdk.RunNewParams{
		TaskID: "task_123456",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", run.ID)
}
