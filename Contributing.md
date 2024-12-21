
## Setup:

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
