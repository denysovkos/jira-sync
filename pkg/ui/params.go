package ui

import (
	"jira-sync/pkg/db"

	"github.com/rivo/tview"
)

func validateNotEmpty(textToCheck string) bool {
	return len(textToCheck) == 0
}

func RenedSetParams(db *db.DB, state *AppState) {
	params, err := db.GetParams()
	if err != nil {
		panic(err)
	}

	if params.Domain != "" && params.Project != "" && params.Token != "" && params.User != "" {
		state.Pages.SwitchToPage("tablePage")
		return
	}

	var actionButtonLabel = "Save"
	if params.Project != "" {
		actionButtonLabel = "Edit"
	}

	state.ParamsFormView.
		AddInputField("Domain:", params.Domain, 100, nil, nil).
		AddInputField("Project:", params.Project, 100, nil, nil).
		AddInputField("Username:", params.User, 100, nil, nil).
		AddInputField("Token:", params.Token, 100, nil, nil).
		AddButton(actionButtonLabel, func() {
			domainItem := state.ParamsFormView.GetFormItem(0)
			projectItem := state.ParamsFormView.GetFormItem(1)
			usernameItem := state.ParamsFormView.GetFormItem(2)
			tokenItem := state.ParamsFormView.GetFormItem(3)

			if usernameItem != nil && tokenItem != nil && domainItem != nil && projectItem != nil {
				domainCell := domainItem.(*tview.InputField)
				projectCell := projectItem.(*tview.InputField)
				usernameCell := usernameItem.(*tview.InputField)
				tokenCell := tokenItem.(*tview.InputField)

				domain := domainCell.GetText()
				project := projectCell.GetText()
				username := usernameCell.GetText()
				token := tokenCell.GetText()

				if validateNotEmpty(domain) {
					Alert(state, "Domain can not be empty")
					return
				}

				if validateNotEmpty(project) {
					Alert(state, "Project can not be empty")
					return
				}

				if validateNotEmpty(username) {
					Alert(state, "User name can not be empty")
					return
				}

				if validateNotEmpty(token) {
					Alert(state, "Token can not be empty")
					return
				}

				db.SetParams(domain, project, username, token)
				state.Pages.SwitchToPage("tablePage")
			} else {
				state.App.Stop()
			}
		}).
		AddButton("Quit", func() {
			state.App.Stop()
		})

	state.ParamsFormView.SetBorder(true).SetTitle("Enter your Jira configuration").SetTitleAlign(tview.AlignLeft)

}
