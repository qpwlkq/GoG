package event_model

// 添加或者删除一个emoji会触发该事件🤣🤣🤣
type EmojiEvent struct {
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
		CiConfigPath      any    `json:"ci_config_path"`
		Homepage          string `json:"homepage"`
		URL               string `json:"url"`
		SSHURL            string `json:"ssh_url"`
		HTTPURL           string `json:"http_url"`
	} `json:"project"`
	ObjectAttributes struct {
		UserID        int    `json:"user_id"`
		CreatedAt     string `json:"created_at"`
		ID            int    `json:"id"`
		Name          string `json:"name"`
		AwardableType string `json:"awardable_type"`
		AwardableID   int    `json:"awardable_id"`
		UpdatedAt     string `json:"updated_at"`
		Action        string `json:"action"`
		AwardedOnURL  string `json:"awarded_on_url"`
	} `json:"object_attributes"`
	Note struct {
		Attachment       any    `json:"attachment"`
		AuthorID         int    `json:"author_id"`
		ChangePosition   any    `json:"change_position"`
		CommitID         any    `json:"commit_id"`
		CreatedAt        string `json:"created_at"`
		DiscussionID     string `json:"discussion_id"`
		ID               int    `json:"id"`
		LineCode         any    `json:"line_code"`
		Note             string `json:"note"`
		NoteableID       int    `json:"noteable_id"`
		NoteableType     string `json:"noteable_type"`
		OriginalPosition any    `json:"original_position"`
		Position         any    `json:"position"`
		ProjectID        int    `json:"project_id"`
		ResolvedAt       any    `json:"resolved_at"`
		ResolvedByID     any    `json:"resolved_by_id"`
		ResolvedByPush   any    `json:"resolved_by_push"`
		StDiff           any    `json:"st_diff"`
		System           bool   `json:"system"`
		Type             any    `json:"type"`
		UpdatedAt        string `json:"updated_at"`
		UpdatedByID      any    `json:"updated_by_id"`
		Description      string `json:"description"`
		URL              string `json:"url"`
	} `json:"note"`
	Issue struct {
		AuthorID            int    `json:"author_id"`
		ClosedAt            any    `json:"closed_at"`
		Confidential        bool   `json:"confidential"`
		CreatedAt           string `json:"created_at"`
		Description         string `json:"description"`
		DiscussionLocked    any    `json:"discussion_locked"`
		DueDate             any    `json:"due_date"`
		ID                  int    `json:"id"`
		Iid                 int    `json:"iid"`
		LastEditedAt        any    `json:"last_edited_at"`
		LastEditedByID      any    `json:"last_edited_by_id"`
		MilestoneID         any    `json:"milestone_id"`
		MovedToID           any    `json:"moved_to_id"`
		DuplicatedToID      any    `json:"duplicated_to_id"`
		ProjectID           int    `json:"project_id"`
		RelativePosition    int    `json:"relative_position"`
		StateID             int    `json:"state_id"`
		TimeEstimate        int    `json:"time_estimate"`
		Title               string `json:"title"`
		UpdatedAt           string `json:"updated_at"`
		UpdatedByID         any    `json:"updated_by_id"`
		Weight              any    `json:"weight"`
		HealthStatus        any    `json:"health_status"`
		URL                 string `json:"url"`
		TotalTimeSpent      int    `json:"total_time_spent"`
		TimeChange          int    `json:"time_change"`
		HumanTotalTimeSpent any    `json:"human_total_time_spent"`
		HumanTimeChange     any    `json:"human_time_change"`
		HumanTimeEstimate   any    `json:"human_time_estimate"`
		AssigneeIds         []int  `json:"assignee_ids"`
		AssigneeID          int    `json:"assignee_id"`
		Labels              []any  `json:"labels"`
		State               string `json:"state"`
		Severity            string `json:"severity"`
	} `json:"issue"`
}