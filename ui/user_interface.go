package ui

import "github.com/rivo/tview"

func Start() {
	app := tview.NewApplication()
	if err := app.SetRoot(mainPage(app), true).Run(); err != nil {
		panic(err)
	}
}

func mainPage(app *tview.Application) *tview.Flex {
	text := tview.NewTextView().SetText("Welcome to the Expense Manager").SetTextAlign(tview.AlignCenter)
	startButton := tview.NewButton("Start").SetSelectedFunc(func() {
		app.SetRoot(commitPage(app), true)
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, false).
		AddItem(startButton, 0, 1, true)

	return flex
}
func commitPage(app *tview.Application) *tview.Flex {
	addButton := tview.NewButton("Add Expense").SetSelectedFunc(func() {
		app.SetRoot(addExpensePage(app), true)
	})
	deleteButton := tview.NewButton("Delete Expense").SetSelectedFunc(func() {
		app.SetRoot(deleteExpensePage(app), true)
	})
	viewButton := tview.NewButton("View Expenses").SetSelectedFunc(func() {
		app.SetRoot(viewExpensesPage(app), true)
	})

	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(addButton, 0, 1, true).
		AddItem(deleteButton, 0, 1, true).
		AddItem(viewButton, 0, 1, true)

	return flex
}
func addExpensePage(app *tview.Application) *tview.Form {
	form := tview.NewForm().
		AddInputField("ID", "123", 20, nil, nil). // Hardcoded ID for now
		AddInputField("Amount", "", 20, nil, nil).
		AddInputField("Description", "", 20, nil, nil).
		AddButton("Add Expense", func() {
			// Handle form submission, add expense to the database
			amount := form.GetFormItemByLabel("Amount").(*tview.InputField).GetText()
			description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()
			// Convert and add to the database
			// ...
			app.SetRoot(commitPage(app), true)
		}).
		AddButton("Cancel", func() {
			app.SetRoot(commitPage(app), true)
		})

	return form
}
func deleteExpensePage(app *tview.Application) *tview.Form {
	form := tview.NewForm().
		AddInputField("ID", "", 20, nil, nil).
		AddButton("Delete Expense", func() {
			id := form.GetFormItemByLabel("ID").(*tview.InputField).GetText()
			// Delete the expense from the database by ID
			// ...
			app.SetRoot(commitPage(app), true)
		}).
		AddButton("Cancel", func() {
			app.SetRoot(commitPage(app), true)
		})

	return form
}
func viewExpensesPage(app *tview.Application) *tview.TextView {
	textView := tview.NewTextView()

	// Retrieve all expenses from the database and display them
	// ...
	textView.SetText("List of expenses:\n...") // Replace with real data

	return textView
}
