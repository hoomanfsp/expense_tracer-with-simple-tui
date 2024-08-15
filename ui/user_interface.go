package ui

import (
	"et_sui/proceed"

	"github.com/jinzhu/gorm"
	"github.com/rivo/tview"
)

func Start(db *gorm.DB) {
	app := tview.NewApplication()
	if err := app.SetRoot(mainPage(app, db), true).Run(); err != nil {
		panic(err)
	}
}

func mainPage(app *tview.Application, db *gorm.DB) *tview.Flex {
	// Create the welcome text view
	text := tview.NewTextView().
		SetText("Welcome to the Expense Manager").
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	// Create the Start button
	startButton := tview.NewButton("Start").
		SetSelectedFunc(func() {
			app.SetRoot(commitPage(app, db), true)
		})

	// Create a flex container to center the button
	buttonFlex := tview.NewFlex().
		AddItem(nil, 0, 1, false).         // Empty space before the button
		AddItem(startButton, 10, 1, true). // Button occupies 10 columns (width)
		AddItem(nil, 0, 1, false)          // Empty space after the button

	// Create the main Flex layout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(text, 0, 1, false).     // Text occupies the main space
		AddItem(buttonFlex, 3, 1, true) // Button row with controlled height

	return flex
}

func commitPage(app *tview.Application, db *gorm.DB) *tview.List {
	// Create a new List
	list := tview.NewList().
		AddItem("Add Expense", "Add a new expense", 'a', func() {
			app.SetRoot(addExpensePage(app, db), true)
		}).
		AddItem("Delete Expense", "Delete an existing expense", 'd', func() {
			app.SetRoot(deleteExpensePage(app, db), true)
		}).
		AddItem("View Expenses", "View all expenses", 'v', func() {
			app.SetRoot(viewExpensesPage(app, db), true)
		}).
		AddItem("Quit", "Exit the application", 'q', func() {
			app.Stop()
		})

	// Return the list as the root element of the commitPage
	return list
}

func addExpensePage(app *tview.Application, db *gorm.DB) *tview.Flex {
	form := tview.NewForm().
		AddInputField("Amount", "", 20, nil, nil).      // Input field for amount
		AddInputField("Description", "", 40, nil, nil). // Input field for description
		AddButton("Submit", nil)                        // Add the Submit button

	// Capture the form reference in a closure for the Submit button
	form.GetButton(0).SetSelectedFunc(func() {
		amount := form.GetFormItemByLabel("Amount").(*tview.InputField).GetText()
		description := form.GetFormItemByLabel("Description").(*tview.InputField).GetText()

		// Save the expense using the GORM DB connection
		proceed.Add(amount, description, db)

		// Return to the main commit page after successful addition
		app.SetRoot(commitPage(app, db), true)
	})

	// Add TextView elements for static text (ID and Note)
	idTextView := tview.NewTextView().
		SetText("ID will be: Automatically generated").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	noteTextView := tview.NewTextView().
		SetText("Note: Remember the ID for deleting later or just select 'View' in the menu").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	// Create a flex layout to arrange the TextViews and the form
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(idTextView, 2, 1, false).  // ID text at the top
		AddItem(form, 0, 1, true).         // The form takes the main space
		AddItem(noteTextView, 2, 1, false) // Note text at the bottom

	flex.SetBorder(true).SetTitle("Add Expense").SetTitleAlign(tview.AlignLeft)
	return flex
}

func deleteExpensePage(app *tview.Application, db *gorm.DB) *tview.Form {
	// Create the form
	form := tview.NewForm()

	// Add the input field and "Delete Expense" button to the form
	form.AddInputField("ID", "", 20, nil, nil).
		AddButton("Delete Expense", func() {
			// Handle deletion of the expense by ID
			id := form.GetFormItemByLabel("ID").(*tview.InputField).GetText()
			proceed.Delete(id, db) // Call your deletion logic here

			// Return to the commit page after deletion
			app.SetRoot(commitPage(app, db), true)
		})

	// Set the form's border and title
	form.SetBorder(true).SetTitle("Delete Expense").SetTitleAlign(tview.AlignLeft)

	return form
}

func viewExpensesPage(app *tview.Application, db *gorm.DB) *tview.Flex {
	// Create a TextView to display the list of expenses
	report := proceed.List(db)
	textView := tview.NewTextView().
		SetDynamicColors(true). // Allows color tags if needed
		SetText(report)         // Replace with real data

	// Create a "Return" button
	returnButton := tview.NewButton("Return").SetSelectedFunc(func() {
		// Navigate back to the commit page when "Return" is selected
		app.SetRoot(commitPage(app, db), true)
	})

	// Create a Flex layout to arrange the TextView and Button vertically
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(textView, 0, 1, false).   // The TextView takes most of the space
		AddItem(returnButton, 1, 1, true) // The "Return" button gets initial focus

	return flex
}
