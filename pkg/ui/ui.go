package ui

import (
	"jira-sync/pkg/db"

	"github.com/rivo/tview"
)

type AppState struct {
	App            *tview.Application
	Pages          *tview.Pages
	ParamsFormView *tview.Form
	TableView      *tview.Flex
	ModalForm      *tview.Modal
}

func Alert(state *AppState, message string) {
	state.ModalForm.SetText(message)
	state.Pages.ShowPage("modalAlert")
}

func Run(db *db.DB) {
	app := tview.NewApplication()
	state := &AppState{
		App:            app,
		Pages:          tview.NewPages(),
		ParamsFormView: tview.NewForm(),
		ModalForm:      tview.NewModal().AddButtons([]string{"Ok!"}),
		TableView:      tview.NewFlex(),
	}

	state.ModalForm.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		state.Pages.HidePage("modalAlert")
	})

	state.Pages.AddPage("paramsFormPage", state.ParamsFormView, true, true)
	state.Pages.AddPage("tablePage", state.TableView, true, false)
	state.Pages.AddPage("modalAlert", state.ModalForm, false, false)

	RenedSetParams(db, state)
	RenderTable(db, state)

	if err := app.SetRoot(state.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
