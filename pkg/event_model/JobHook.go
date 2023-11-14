package event_model

import (
	"time"
)

type JobEvent struct {
	ObjectKind          string    `json:"object_kind"`
	Ref                 string    `json:"ref"`
	Tag                 bool      `json:"tag"`
	BeforeSha           string    `json:"before_sha"`
	Sha                 string    `json:"sha"`
	BuildID             int       `json:"build_id"`
	BuildName           string    `json:"build_name"`
	BuildStage          string    `json:"build_stage"`
	BuildStatus         string    `json:"build_status"`
	BuildCreatedAt      time.Time `json:"build_created_at"`
	BuildStartedAt      any       `json:"build_started_at"`
	BuildFinishedAt     any       `json:"build_finished_at"`
	BuildDuration       any       `json:"build_duration"`
	BuildQueuedDuration float64   `json:"build_queued_duration"`
	BuildAllowFailure   bool      `json:"build_allow_failure"`
	BuildFailureReason  string    `json:"build_failure_reason"`
	RetriesCount        int       `json:"retries_count"`
	PipelineID          int       `json:"pipeline_id"`
	ProjectID           int       `json:"project_id"`
	ProjectName         string    `json:"project_name"`
	User                struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		AvatarURL string `json:"avatar_url"`
	} `json:"user"`
	Commit struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Sha         string `json:"sha"`
		Message     string `json:"message"`
		AuthorName  string `json:"author_name"`
		AuthorEmail string `json:"author_email"`
		Status      string `json:"status"`
		Duration    any    `json:"duration"`
		StartedAt   any    `json:"started_at"`
		FinishedAt  any    `json:"finished_at"`
	} `json:"commit"`
	Repository struct {
		Name            string `json:"name"`
		Description     string `json:"description"`
		Homepage        string `json:"homepage"`
		GitSSHURL       string `json:"git_ssh_url"`
		GitHTTPURL      string `json:"git_http_url"`
		VisibilityLevel int    `json:"visibility_level"`
	} `json:"repository"`
	Project struct {
		ID                int    `json:"id"`
		Name              string `json:"name"`
		Description       string `json:"description"`
		WebURL            string `json:"web_url"`
		AvatarURL         any    `json:"avatar_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		GitHTTPURL        string `json:"git_http_url"`
		Namespace         string `json:"namespace"`
		VisibilityLevel   int    `json:"visibility_level"`
		PathWithNamespace string `json:"path_with_namespace"`
		DefaultBranch     string `json:"default_branch"`
	} `json:"project"`
	Runner struct {
		Active      bool     `json:"active"`
		RunnerType  string   `json:"runner_type"`
		IsShared    bool     `json:"is_shared"`
		ID          int      `json:"id"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	} `json:"runner"`
	Environment    any `json:"environment"`
	SourcePipeline struct {
		Project struct {
			ID                int    `json:"id"`
			WebURL            string `json:"web_url"`
			PathWithNamespace string `json:"path_with_namespace"`
		} `json:"project"`
		PipelineID int `json:"pipeline_id"`
		JobID      int `json:"job_id"`
	} `json:"source_pipeline"`
}