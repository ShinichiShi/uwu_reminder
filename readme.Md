# UwU Reminder

UwU Reminder is a fun and quirky tool to help you keep track of your tasks and deadlines in your linux System. It turns boring reminders into something delightful with an adorable UwU twist, ensuring that staying on top of your schedule never feels like a chore!

---

## Why UwU Reminder?
Terminals are a developer's favorite workspace, so why not use them for reminders? UwU Reminder delivers cute, playful notifications like "UwU, time to hydrate! 💧" directly in your Linux environment. It's lightweight, blends seamlessly into your workflow, and keeps you on track with a touch of fun! ✨

---

## Key Features

- **Custom Notifications:** Add a fun UwU-style flair to your reminders.
- **Task Management:** Simple and intuitive interface to schedule and track tasks.
- **Multi-Platform Support:** Works across devices to keep you reminded wherever you go.

---

## Tech Stack : 
 Primarily made in Golang. 

---

## How to Run Locally

### Prerequisites:
- Go installed on your system (minimum version: 1.16 or higher).
- A terminal with access to git commands.

### Clone the repository:
```bash
git clone https://github.com/ShinichiShi/uwu_reminder.git
cd uwu_reminder
```
Install dependencies: Use go mod tidy to download and sync the required modules:

``` 
go mod tidy 
```
Run the application: Start the CLI tool with:

```
go run main.go
```
Usage Example:
```
go run main.go "2024-12-25 10:00" "UwU~ I'm always With You!"

```
## How to Set it Up Forever
```
go build -o remind
sudo mv remind /usr/local/bin/
remind <time> <message>
```
Replace <code>\<time\></code> into 24 hr/ 12 hr format with am/pm :)

---

# Contributing

We welcome contributions to UwU Reminder! Here’s how you can help:

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m 'Add feature-name'
   ```
4. Push to the branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request.

---

# License

This project is licensed under the [MIT License](LICENSE).
