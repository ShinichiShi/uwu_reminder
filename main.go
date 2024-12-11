package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

const (
	markName  = "cli_reminder"
	markValue = "1"
)

var recurConfig = make(map[int]string)     
var messageConfig = make(map[int]string)  
var idCounter int = 1                     
func main() {
	helpFlag := flag.Bool("help", false, "Displays usage information")
	recurFlag := flag.String("recur", "", "Set recurrence pattern (daily, weekly, monthly)")
	removeRecurFlag := flag.Bool("remove-recur", false, "Remove all recurring reminders")
	listRecurFlag := flag.Bool("list-recur", false, "List all recurring reminders")
	removeByIDFlag := flag.Int("remove-id", 0, "Remove recurring reminder by ID")
	flag.Parse()

	// Handle flags first
	if *helpFlag {
		printUsage()
		os.Exit(0)
	}

	if *listRecurFlag {
		listRecurringReminders()
		os.Exit(0)
	}

	if *removeByIDFlag > 0 {
		removeRecurringReminderByID(*removeByIDFlag)
		os.Exit(0)
	}

	if *removeRecurFlag {
		removeAllRecurringReminders()
		os.Exit(0)
	}

	// After flag parsing, get the remaining positional arguments
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Error: Insufficient arguments provided.")
		printUsage()
		os.Exit(1)
	}

	welcomeMessage()

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	// Parse the input date/time
	timeInput := args[0]
	message := strings.Join(args[1:], " ")
	t, err := w.Parse(timeInput, now)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		os.Exit(1)
	}

	if t == nil {
		fmt.Println("Please specify a valid date/time.")
		os.Exit(2)
	}

	if now.After(t.Time) {
		fmt.Println("Cannot set a reminder in the past!")
		os.Exit(3)
	}

	diff := time.Until(t.Time)  // Updated to use time.Until

	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err = beeep.Alert("Reminder", message, "assets/information.png")
		if err != nil {
			fmt.Println("Error displaying notification:", err)
			os.Exit(4)
		}

		handleRecurrence(t.Time, *recurFlag, message)
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))

		if err = cmd.Start(); err != nil {
			fmt.Println("Error starting new process:", err)
			os.Exit(5)
		}

		fmt.Println("Reminder will trigger in:", diff.Round(time.Second))
		os.Exit(0)
	}
}

func handleRecurrence(initialTime time.Time, recurrence, message string) {
	if recurrence == "" {
		return
	}

	var nextTime time.Time

	switch strings.ToLower(recurrence) {
	case "daily":
		nextTime = initialTime.Add(24 * time.Hour)
	case "weekly":
		nextTime = initialTime.Add(7 * 24 * time.Hour)
	case "monthly":
		nextTime = initialTime.AddDate(0, 1, 0)
	default:
		return
	}

	recurConfig[idCounter] = recurrence // Save recurrence info
	messageConfig[idCounter] = message
	currentID := idCounter
	idCounter++

	diff := time.Until(nextTime)
	fmt.Printf("Recurring reminder set (ID: %d) for: %s\n", currentID, nextTime.Format("2006-01-02 15:04:05"))
	time.Sleep(diff)
	err := beeep.Alert("Recurring Reminder", message, "assets/information.png")
	if err != nil {
		fmt.Println("Error displaying recurring notification:", err)
		os.Exit(6)
	}

	handleRecurrence(nextTime, recurrence, message)
}

func removeAllRecurringReminders() {
	if len(recurConfig) == 0 {
		fmt.Println("No recurring reminders to remove.")
		return
	}

	fmt.Println("Removing all recurring reminders...")
	recurConfig = make(map[int]string) // Clear all entries
	messageConfig = make(map[int]string)
	fmt.Println("All recurring reminders removed successfully.")
}

func removeRecurringReminderByID(id int) {
	if _, exists := recurConfig[id]; !exists {
		fmt.Printf("No recurring reminder found with ID: %d\n", id)
		return
	}

	delete(recurConfig, id)
	delete(messageConfig, id)
	fmt.Printf("Recurring reminder with ID: %d removed successfully.\n", id)
}

func listRecurringReminders() {
	if len(recurConfig) == 0 {
		fmt.Println("No recurring reminders set.")
		return
	}

	fmt.Println("Listing all recurring reminders:")
	for id, recurrence := range recurConfig {
		fmt.Printf("ID: %d, Recurrence: %s, Message: %s\n", id, recurrence, messageConfig[id])
	}
}

func printUsage() {
	fmt.Println("CLI Reminder Tool")
	fmt.Println("Usage:")
	fmt.Println("  remind <date/time> <message> [flags]")
	fmt.Println("Flags:")
	fmt.Println("  --help          Display usage information")
	fmt.Println("  --recur         Set recurrence pattern (daily, weekly, monthly)")
	fmt.Println("  --remove-recur  Remove all recurring reminders")
	fmt.Println("  --list-recur    List all recurring reminders")
	fmt.Println("  --remove-id     Remove recurring reminder by ID")
	fmt.Println("Examples:")
	fmt.Println("  remind \"2024-12-12 14:30\" \"Meeting with team\"")
	fmt.Println("  remind \"2pm tomorrow\" \"Doctor's appointment\" --recur weekly")
	fmt.Println("  remind --list-recur")
	fmt.Println("  remind --remove-id 2")
	fmt.Println("  remind --remove-recur")
}

func welcomeMessage() {
	fmt.Println("Welcome to the CLI Reminder Tool!")
	fmt.Println("Type 'remind --help' for usage instructions.")
}
