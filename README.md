# 🤖 Golang Slack Bot – Happy to Help Teammates!

A simple yet powerful Slack bot built using **Go (Golang)** that listens for the `ping` command and responds with `pong`, current date, time ⏰, and a cheerful message 😄.

---

## ✨ Features

- 🛠 Built in **Golang** using the [slacker](https://github.com/shomali11/slacker) library
- 🔁 Listens for `/ping` command in Slack
- 📅 Replies with:
  - `pong`
  - Current date and time
  - Friendly weather message
- 📊 Logs command usage in real-time

---

## 🚀 Getting Started

### 1. Clone this repo

```bash
git clone https://github.com/your-username/golang-slack-bot.git
cd golang-slack-bot
2. Set your Slack tokens
Create a .env file or export environment variables:

env
Copy
Edit
SLACK_BOT_TOKEN=your-slack-bot-token
SLACK_APP_TOKEN=your-slack-app-token
3. Run the bot
bash
Copy
Edit
go run main.go
🧪 Example Usage
In Slack:

bash
Copy
Edit
/ping
Bot replies:

pgsql
Copy
Edit
pong
🕒 Current Time: 10 May 2025 15:21:33 IST
😊 I hope the weather is nice!
Happy to help, Slack teammates!
🧠 Skills Showcased
Golang concurrency with goroutines (go PrintCommandEvents(...))

Slack Bot development

Real-time logging

Clean and modular code

🤝 Contributing
Pull requests and stars are welcome! ⭐

📜 License
MIT License

yaml
Copy
Edit

---

### ✅ Now Steps to Push to GitHub:

```bash
echo "# golang-slack-bot" >> README.md
git init
git add .
git commit -m "Initial commit: Slack bot in Golang"
gh repo create golang-slack-bot --public --source=. --remote=origin
git push -u origin master
