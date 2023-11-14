package event_model

import (
	"time"
)

type CommentOnACommitEvent struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
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
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id"`
		Note         string `json:"note"`
		NoteableType string `json:"noteable_type"`
		AuthorID     int    `json:"author_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		ProjectID    int    `json:"project_id"`
		Attachment   any    `json:"attachment"`
		LineCode     string `json:"line_code"`
		CommitID     string `json:"commit_id"`
		NoteableID   any    `json:"noteable_id"`
		System       bool   `json:"system"`
		StDiff       struct {
			Diff        string `json:"diff"`
			NewPath     string `json:"new_path"`
			OldPath     string `json:"old_path"`
			AMode       string `json:"a_mode"`
			BMode       string `json:"b_mode"`
			NewFile     bool   `json:"new_file"`
			RenamedFile bool   `json:"renamed_file"`
			DeletedFile bool   `json:"deleted_file"`
		} `json:"st_diff"`
		URL string `json:"url"`
	} `json:"object_attributes"`
	Commit struct {
		ID        string    `json:"id"`
		Message   string    `json:"message"`
		Timestamp time.Time `json:"timestamp"`
		URL       string    `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commit"`
}

type CommentOnAMergeRequestEvent struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
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
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id"`
		Note         string `json:"note"`
		NoteableType string `json:"noteable_type"`
		AuthorID     int    `json:"author_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		ProjectID    int    `json:"project_id"`
		Attachment   any    `json:"attachment"`
		LineCode     any    `json:"line_code"`
		CommitID     string `json:"commit_id"`
		NoteableID   int    `json:"noteable_id"`
		System       bool   `json:"system"`
		StDiff       any    `json:"st_diff"`
		URL          string `json:"url"`
	} `json:"object_attributes"`
	MergeRequest struct {
		ID              int    `json:"id"`
		TargetBranch    string `json:"target_branch"`
		SourceBranch    string `json:"source_branch"`
		SourceProjectID int    `json:"source_project_id"`
		AuthorID        int    `json:"author_id"`
		AssigneeID      int    `json:"assignee_id"`
		Title           string `json:"title"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		MilestoneID     int    `json:"milestone_id"`
		State           string `json:"state"`
		MergeStatus     string `json:"merge_status"`
		TargetProjectID int    `json:"target_project_id"`
		Iid             int    `json:"iid"`
		Description     string `json:"description"`
		Position        int    `json:"position"`
		Labels          []struct {
			ID          int       `json:"id"`
			Title       string    `json:"title"`
			Color       string    `json:"color"`
			ProjectID   any       `json:"project_id"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
			Template    bool      `json:"template"`
			Description any       `json:"description"`
			Type        string    `json:"type"`
			GroupID     int       `json:"group_id"`
		} `json:"labels"`
		Source struct {
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
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"source"`
		Target struct {
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
			Homepage          string `json:"homepage"`
			URL               string `json:"url"`
			SSHURL            string `json:"ssh_url"`
			HTTPURL           string `json:"http_url"`
		} `json:"target"`
		LastCommit struct {
			ID        string `json:"id"`
			Message   string `json:"message"`
			Timestamp string `json:"timestamp"`
			URL       string `json:"url"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
		} `json:"last_commit"`
		WorkInProgress bool `json:"work_in_progress"`
		Draft          bool `json:"draft"`
		Assignee       struct {
			Name      string `json:"name"`
			Username  string `json:"username"`
			AvatarURL string `json:"avatar_url"`
		} `json:"assignee"`
		DetailedMergeStatus string `json:"detailed_merge_status"`
	} `json:"merge_request"`
}

type CommentOnAnIssueEvent struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
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
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id"`
		Note         string `json:"note"`
		NoteableType string `json:"noteable_type"`
		AuthorID     int    `json:"author_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		ProjectID    int    `json:"project_id"`
		Attachment   any    `json:"attachment"`
		LineCode     any    `json:"line_code"`
		CommitID     string `json:"commit_id"`
		NoteableID   int    `json:"noteable_id"`
		System       bool   `json:"system"`
		StDiff       any    `json:"st_diff"`
		URL          string `json:"url"`
	} `json:"object_attributes"`
	Issue struct {
		ID          int    `json:"id"`
		Title       string `json:"title"`
		AssigneeIds []any  `json:"assignee_ids"`
		AssigneeID  any    `json:"assignee_id"`
		AuthorID    int    `json:"author_id"`
		ProjectID   int    `json:"project_id"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
		Position    int    `json:"position"`
		BranchName  any    `json:"branch_name"`
		Description string `json:"description"`
		MilestoneID any    `json:"milestone_id"`
		State       string `json:"state"`
		Iid         int    `json:"iid"`
		Labels      []struct {
			ID          int       `json:"id"`
			Title       string    `json:"title"`
			Color       string    `json:"color"`
			ProjectID   any       `json:"project_id"`
			CreatedAt   time.Time `json:"created_at"`
			UpdatedAt   time.Time `json:"updated_at"`
			Template    bool      `json:"template"`
			Description any       `json:"description"`
			Type        string    `json:"type"`
			GroupID     int       `json:"group_id"`
		} `json:"labels"`
	} `json:"issue"`
}

type CommentOnACodeSnippet struct {
	ObjectKind string `json:"object_kind"`
	EventType  string `json:"event_type"`
	User       struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Username  string `json:"username"`
		AvatarURL string `json:"avatar_url"`
		Email     string `json:"email"`
	} `json:"user"`
	ProjectID int `json:"project_id"`
	Project   struct {
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
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
	} `json:"repository"`
	ObjectAttributes struct {
		ID           int    `json:"id"`
		Note         string `json:"note"`
		NoteableType string `json:"noteable_type"`
		AuthorID     int    `json:"author_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		ProjectID    int    `json:"project_id"`
		Attachment   any    `json:"attachment"`
		LineCode     any    `json:"line_code"`
		CommitID     string `json:"commit_id"`
		NoteableID   int    `json:"noteable_id"`
		System       bool   `json:"system"`
		StDiff       any    `json:"st_diff"`
		URL          string `json:"url"`
	} `json:"object_attributes"`
	Snippet struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		Content         string `json:"content"`
		AuthorID        int    `json:"author_id"`
		ProjectID       int    `json:"project_id"`
		CreatedAt       string `json:"created_at"`
		UpdatedAt       string `json:"updated_at"`
		FileName        string `json:"file_name"`
		ExpiresAt       any    `json:"expires_at"`
		Type            string `json:"type"`
		VisibilityLevel int    `json:"visibility_level"`
		URL             string `json:"url"`
	} `json:"snippet"`
}