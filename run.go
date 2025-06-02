// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package vernsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/vern-sdk-go/internal/apijson"
	"github.com/stainless-sdks/vern-sdk-go/internal/requestconfig"
	"github.com/stainless-sdks/vern-sdk-go/option"
	"github.com/stainless-sdks/vern-sdk-go/packages/param"
	"github.com/stainless-sdks/vern-sdk-go/packages/respjson"
)

// RunService contains methods and other services that help with interacting with
// the vern API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRunService] method instead.
type RunService struct {
	Options []option.RequestOption
}

// NewRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRunService(opts ...option.RequestOption) (r RunService) {
	r = RunService{}
	r.Options = opts
	return
}

// Executes a task with the provided inputs
func (r *RunService) New(ctx context.Context, body RunNewParams, opts ...option.RequestOption) (res *RunNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves the details of a specific task run
func (r *RunService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *RunGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("runs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RunNewResponse struct {
	// Unique identifier for the run
	ID string `json:"id"`
	// The inputs provided for the task
	Inputs map[string]any `json:"inputs"`
	// Timestamp when the run was queued
	QueuedAt time.Time `json:"queued_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Inputs      respjson.Field
		QueuedAt    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *RunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RunGetResponse struct {
	// Unique identifier for the run
	ID string `json:"id"`
	// Timestamp when the run completed
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// Timestamp when the run was created
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The inputs provided for the task
	Inputs map[string]any `json:"inputs"`
	// The response data from the task execution
	Response map[string]any `json:"response"`
	// Timestamp when the run started executing
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// The current status of the run
	//
	// Any of "queued", "running", "complete", "failed".
	Status RunGetResponseStatus `json:"status"`
	// The name of the task that was executed
	Task string `json:"task"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Inputs      respjson.Field
		Response    respjson.Field
		StartedAt   respjson.Field
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *RunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the run
type RunGetResponseStatus string

const (
	RunGetResponseStatusQueued   RunGetResponseStatus = "queued"
	RunGetResponseStatusRunning  RunGetResponseStatus = "running"
	RunGetResponseStatusComplete RunGetResponseStatus = "complete"
	RunGetResponseStatusFailed   RunGetResponseStatus = "failed"
)

type RunNewParams struct {
	// The ID of the task to execute
	TaskID string `json:"taskId,required"`
	// Optional user-specified UID for a profile linked via magic link
	ProfileID param.Opt[string] `json:"profileId,omitzero"`
	// An optional URL to be processed by the task
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// The inputs required for the task
	Inputs map[string]any `json:"inputs,omitzero"`
	paramObj
}

func (r RunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
