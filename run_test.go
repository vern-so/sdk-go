// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vernsdk_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/vern-so/sdk-go"
	"github.com/vern-so/sdk-go/internal/testutil"
	"github.com/vern-so/sdk-go/option"
)

func TestRunNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Runs.New(context.TODO(), vernsdk.RunNewParams{
		TaskID: "task_123456",
		Inputs: map[string]any{
			"prompt": "bar",
			"text":   "bar",
		},
		ProfileID: vernsdk.String("profileId"),
		URL:       vernsdk.String("https://example.com"),
	})
	if err != nil {
		var apierr *vernsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRunGet(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Runs.Get(context.TODO(), "id")
	if err != nil {
		var apierr *vernsdk.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
