package ui

import (
	"fmt"
	"jira-sync/pkg/db"
	"jira-sync/pkg/jira"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RenderTable(db *db.DB, state *AppState) {
	table := tview.NewTable().SetSelectable(true, false)

	table.SetBorder(true).
		SetTitle(" Issues ").
		SetTitleAlign(tview.AlignCenter).
		SetBorderPadding(0, 0, 1, 1).
		SetBorderAttributes(tcell.AttrBold)

	params, err := db.GetParams()
	if err != nil {
		panic(err)
	}

	jiraClient := jira.NewJiraClient(params.User, params.Token)
	pendingTaskResponse, err := jiraClient.GetPendingTasks(params)

	if err != nil {
		fmt.Println("Error getting pending tasks:", err)
		return
	}

	table.SetCell(0, 0, tview.NewTableCell("Issue id").SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Estimated time").SetAlign(tview.AlignCenter).SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Tracked time").SetAlign(tview.AlignCenter).SetSelectable(false))

	// Create a form on the right
	jiraIdInput := tview.NewTextView().SetLabel("Id").SetSize(1, 50)
	jiraTypeInput := tview.NewTextView().SetLabel("Issue type").SetSize(1, 50)
	jiraTitleInput := tview.NewTextView().SetLabel("Title").SetSize(1, 50)
	jiraDescriptionInput := tview.NewTextView().SetLabel("Description").SetSize(5, 50)
	jiraEstimatedTimeInput := tview.NewTextView().SetLabel("Estimated time").SetSize(1, 50)
	jiraSpentTimeInput := tview.NewTextView().SetLabel("Logged time").SetSize(1, 50)
	jiraAddSpentTimeImput := tview.NewInputField().SetLabel("Log time")

	if (len(pendingTaskResponse.Issues)) > 0 {
		// FILL THE TABLE
		for i, v := range pendingTaskResponse.Issues {
			keyCell := tview.NewTableCell(v.Key).
				SetAlign(tview.AlignCenter).
				SetSelectable(true)

			estimatedDuration := time.Duration(v.Fields.Timeestimate * 1000 * 1000 * 1000)
			spentDuration := time.Duration(v.Fields.Timespent * 1000 * 1000 * 1000)

			estimateCell := tview.NewTableCell(estimatedDuration.String()).
				SetAlign(tview.AlignCenter).
				SetSelectable(true)

			spentCell := tview.NewTableCell(spentDuration.String()).
				SetAlign(tview.AlignCenter).
				SetSelectable(true)

			table.SetCell(i+1, 0, keyCell)
			table.SetCell(i+1, 1, estimateCell)
			table.SetCell(i+1, 2, spentCell)
		}

		estimatedDuration := time.Duration(pendingTaskResponse.Issues[0].Fields.Timeestimate * 1000 * 1000 * 1000)
		spentDuration := time.Duration(pendingTaskResponse.Issues[0].Fields.Timespent * 1000 * 1000 * 1000)

		// SET FIRST ISSUE VALUE TO PREVIEW
		jiraIdInput.SetText(pendingTaskResponse.Issues[0].Key)
		jiraTypeInput.SetText(pendingTaskResponse.Issues[0].Fields.Issuetype.Name)
		jiraTitleInput.SetText(pendingTaskResponse.Issues[0].Fields.Summary)
		jiraDescriptionInput.SetText(pendingTaskResponse.Issues[0].Fields.Status.Description)
		jiraEstimatedTimeInput.SetText(estimatedDuration.String())
		jiraSpentTimeInput.SetText(spentDuration.String())
	}

	updateForm := func() {
		row, _ := table.GetSelection()
		selectedRow := row - 1

		if row > 0 {
			estimatedDuration := time.Duration(pendingTaskResponse.Issues[selectedRow].Fields.Timeestimate * 1000 * 1000 * 1000)
			spentDuration := time.Duration(pendingTaskResponse.Issues[selectedRow].Fields.Timespent * 1000 * 1000 * 1000)

			jiraIdInput.SetText(pendingTaskResponse.Issues[selectedRow].Key).SetMaxLines(1)
			jiraTypeInput.SetText(pendingTaskResponse.Issues[selectedRow].Fields.Issuetype.Name).SetMaxLines(2)
			jiraEstimatedTimeInput.SetText(estimatedDuration.String()).SetMaxLines(2)
			jiraDescriptionInput.SetText(pendingTaskResponse.Issues[selectedRow].Fields.Status.Description).SetMaxLines(5)
			jiraSpentTimeInput.SetText(spentDuration.String()).SetMaxLines(2)
			jiraTitleInput.SetText(pendingTaskResponse.Issues[selectedRow].Fields.Summary).SetMaxLines(2)
		}
	}

	form := tview.NewForm().
		AddFormItem(jiraIdInput).
		AddFormItem(jiraTypeInput).
		AddFormItem(jiraTitleInput).
		AddFormItem(jiraDescriptionInput).
		AddFormItem(jiraEstimatedTimeInput).
		AddFormItem(jiraSpentTimeInput).
		AddFormItem(jiraAddSpentTimeImput).
		AddButton("Submit", func() {
			row, _ := table.GetSelection()
			selectedRow := row - 1

			issueKey := pendingTaskResponse.Issues[selectedRow].Key
			addValue := jiraAddSpentTimeImput.GetText()

			err := jiraClient.UpdateTime(params, issueKey, addValue)
			if err != nil {
				Alert(state, err.Error())
			} else {
				updatedPendingTaskResponse, err := jiraClient.GetPendingTasks(params)
				if err != nil {
					fmt.Println("Error getting pending tasks:", err)
					return
				}
				pendingTaskResponse = updatedPendingTaskResponse
				updateForm()
			}
		}).
		AddButton("Quit", func() {
			state.App.Stop()
		})

	state.TableView.
		SetDirection(tview.FlexColumn).
		AddItem(table, 0, 1, true).
		AddItem(form, 0, 1, false)

	state.TableView.SetBorder(true).SetTitle(" Current tasks without time tracking ")

	table.SetSelectedFunc(func(row, column int) {
		if (len(pendingTaskResponse.Issues)) > 0 {
			updateForm()
		}
	})
}
