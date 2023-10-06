package jira

type PendingTaskResponse struct {
	Expand     string `json:"expand,omitempty"`
	StartAt    int    `json:"startAt,omitempty"`
	MaxResults int    `json:"maxResults,omitempty"`
	Total      int    `json:"total,omitempty"`
	Issues     []struct {
		Expand string `json:"expand,omitempty"`
		ID     string `json:"id,omitempty"`
		Self   string `json:"self,omitempty"`
		Key    string `json:"key,omitempty"`
		Fields struct {
			Statuscategorychangedate string `json:"statuscategorychangedate,omitempty"`
			Parent                   struct {
				ID     string `json:"id,omitempty"`
				Key    string `json:"key,omitempty"`
				Self   string `json:"self,omitempty"`
				Fields struct {
					Summary string `json:"summary,omitempty"`
					Status  struct {
						Self           string `json:"self,omitempty"`
						Description    string `json:"description,omitempty"`
						IconURL        string `json:"iconUrl,omitempty"`
						Name           string `json:"name,omitempty"`
						ID             string `json:"id,omitempty"`
						StatusCategory struct {
							Self      string `json:"self,omitempty"`
							ID        int    `json:"id,omitempty"`
							Key       string `json:"key,omitempty"`
							ColorName string `json:"colorName,omitempty"`
							Name      string `json:"name,omitempty"`
						} `json:"statusCategory,omitempty"`
					} `json:"status,omitempty"`
					Priority struct {
						Self    string `json:"self,omitempty"`
						IconURL string `json:"iconUrl,omitempty"`
						Name    string `json:"name,omitempty"`
						ID      string `json:"id,omitempty"`
					} `json:"priority,omitempty"`
					Issuetype struct {
						Self           string `json:"self,omitempty"`
						ID             string `json:"id,omitempty"`
						Description    string `json:"description,omitempty"`
						IconURL        string `json:"iconUrl,omitempty"`
						Name           string `json:"name,omitempty"`
						Subtask        bool   `json:"subtask,omitempty"`
						HierarchyLevel int    `json:"hierarchyLevel,omitempty"`
					} `json:"issuetype,omitempty"`
				} `json:"fields,omitempty"`
			} `json:"parent,omitempty"`
			FixVersions []any  `json:"fixVersions,omitempty"`
			Resolution  any    `json:"resolution,omitempty"`
			LastViewed  string `json:"lastViewed,omitempty"`
			Priority    struct {
				Self    string `json:"self,omitempty"`
				IconURL string `json:"iconUrl,omitempty"`
				Name    string `json:"name,omitempty"`
				ID      string `json:"id,omitempty"`
			} `json:"priority,omitempty"`
			Labels                        []string `json:"labels,omitempty"`
			Timeestimate                  int      `json:"timeestimate,omitempty"`
			Aggregatetimeoriginalestimate int      `json:"aggregatetimeoriginalestimate,omitempty"`
			Versions                      []any    `json:"versions,omitempty"`
			Issuelinks                    []any    `json:"issuelinks,omitempty"`
			Assignee                      struct {
				Self         string `json:"self,omitempty"`
				AccountID    string `json:"accountId,omitempty"`
				EmailAddress string `json:"emailAddress,omitempty"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48,omitempty"`
					Two4X24   string `json:"24x24,omitempty"`
					One6X16   string `json:"16x16,omitempty"`
					Three2X32 string `json:"32x32,omitempty"`
				} `json:"avatarUrls,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
				Active      bool   `json:"active,omitempty"`
				TimeZone    string `json:"timeZone,omitempty"`
				AccountType string `json:"accountType,omitempty"`
			} `json:"assignee,omitempty"`
			Status struct {
				Self           string `json:"self,omitempty"`
				Description    string `json:"description,omitempty"`
				IconURL        string `json:"iconUrl,omitempty"`
				Name           string `json:"name,omitempty"`
				ID             string `json:"id,omitempty"`
				StatusCategory struct {
					Self      string `json:"self,omitempty"`
					ID        int    `json:"id,omitempty"`
					Key       string `json:"key,omitempty"`
					ColorName string `json:"colorName,omitempty"`
					Name      string `json:"name,omitempty"`
				} `json:"statusCategory,omitempty"`
			} `json:"status,omitempty"`
			Components []struct {
				Self string `json:"self,omitempty"`
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"components,omitempty"`
			Aggregatetimeestimate int `json:"aggregatetimeestimate,omitempty"`
			Creator               struct {
				Self         string `json:"self,omitempty"`
				AccountID    string `json:"accountId,omitempty"`
				EmailAddress string `json:"emailAddress,omitempty"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48,omitempty"`
					Two4X24   string `json:"24x24,omitempty"`
					One6X16   string `json:"16x16,omitempty"`
					Three2X32 string `json:"32x32,omitempty"`
				} `json:"avatarUrls,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
				Active      bool   `json:"active,omitempty"`
				TimeZone    string `json:"timeZone,omitempty"`
				AccountType string `json:"accountType,omitempty"`
			} `json:"creator,omitempty"`
			Subtasks []any `json:"subtasks,omitempty"`
			Reporter struct {
				Self         string `json:"self,omitempty"`
				AccountID    string `json:"accountId,omitempty"`
				EmailAddress string `json:"emailAddress,omitempty"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48,omitempty"`
					Two4X24   string `json:"24x24,omitempty"`
					One6X16   string `json:"16x16,omitempty"`
					Three2X32 string `json:"32x32,omitempty"`
				} `json:"avatarUrls,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
				Active      bool   `json:"active,omitempty"`
				TimeZone    string `json:"timeZone,omitempty"`
				AccountType string `json:"accountType,omitempty"`
			} `json:"reporter,omitempty"`
			Aggregateprogress struct {
				Progress int `json:"progress,omitempty"`
				Total    int `json:"total,omitempty"`
				Percent  int `json:"percent,omitempty"`
			} `json:"aggregateprogress,omitempty"`
			Progress struct {
				Progress int `json:"progress,omitempty"`
				Total    int `json:"total,omitempty"`
				Percent  int `json:"percent,omitempty"`
			} `json:"progress,omitempty"`
			Votes struct {
				Self     string `json:"self,omitempty"`
				Votes    int    `json:"votes,omitempty"`
				HasVoted bool   `json:"hasVoted,omitempty"`
			} `json:"votes,omitempty"`
			Issuetype struct {
				Self           string `json:"self,omitempty"`
				ID             string `json:"id,omitempty"`
				Description    string `json:"description,omitempty"`
				IconURL        string `json:"iconUrl,omitempty"`
				Name           string `json:"name,omitempty"`
				Subtask        bool   `json:"subtask,omitempty"`
				HierarchyLevel int    `json:"hierarchyLevel,omitempty"`
			} `json:"issuetype,omitempty"`
			Timespent int `json:"timespent,omitempty"`
			Project   struct {
				Self           string `json:"self,omitempty"`
				ID             string `json:"id,omitempty"`
				Key            string `json:"key,omitempty"`
				Name           string `json:"name,omitempty"`
				ProjectTypeKey string `json:"projectTypeKey,omitempty"`
				Simplified     bool   `json:"simplified,omitempty"`
				AvatarUrls     struct {
					Four8X48  string `json:"48x48,omitempty"`
					Two4X24   string `json:"24x24,omitempty"`
					One6X16   string `json:"16x16,omitempty"`
					Three2X32 string `json:"32x32,omitempty"`
				} `json:"avatarUrls,omitempty"`
			} `json:"project,omitempty"`
			Aggregatetimespent any `json:"aggregatetimespent,omitempty"`
			Resolutiondate     any `json:"resolutiondate,omitempty"`
			Workratio          int `json:"workratio,omitempty"`
			Watches            struct {
				Self       string `json:"self,omitempty"`
				WatchCount int    `json:"watchCount,omitempty"`
				IsWatching bool   `json:"isWatching,omitempty"`
			} `json:"watches,omitempty"`
			Created     string `json:"created,omitempty"`
			Updated     string `json:"updated,omitempty"`
			Description any    `json:"description,omitempty"`
			Security    any    `json:"security,omitempty"`
			Summary     string `json:"summary,omitempty"`
			Environment any    `json:"environment,omitempty"`
			Duedate     any    `json:"duedate,omitempty"`
		} `json:"fields,omitempty"`
	} `json:"issues,omitempty"`
}
