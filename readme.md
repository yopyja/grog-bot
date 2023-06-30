# Grog Bot 🍻

## Overview 🤠

Grog Bot is a Discord bot written in Golang. It scrapes Utah's DABS website for liquor stock numbers and sends notifications to subscribed members when their favorite liquors are in stock or arrive at the warehouse. Grog Bot temporarily stores data in JSON files for quick retrieval and updates.

## Features 📋
1. Web scraping: Grog Bot continually monitors Utah's DABS website for updates on liquor stock numbers.
2. Notifications: The bot informs members when their preferred liquor products are available.
3. Subscription: Members can subscribe to notifications for specific liquor types or brands.

## Getting Started 🚀

### Prerequisites 📋
To run this bot, you need:
- Go installed on your machine.
- A Discord bot token.
- A cookie from the Utah's DABS site (obtainable via Postman).

### Extracting Cookie using Postman 🍪
Follow these steps to extract the cookie from Utah's DABS site:
1. Download and install Postman.
2. Open Postman and create a new request.
3. Set the request type to `GET` and enter the URL of the Utah DABS site.
4. Send the request and then check the `Cookies` tab in the response section. 
5. Save the cookie information and add it to the bot's config file.

### Installation and Setup ⚒️
1. Clone the repository: `git clone https://github.com/yopyja/grog-bot.git`
2. Change into the directory: `cd grog-bot`
3. Install dependencies: `go mod tidy`
4. Create a config file: `cp config.json.example config.json`
5. Edit the config file and add your bot token and cookie.
5. Start the bot: `go run main.go`

Config File
```json
    {
    "token": "paste your token here",
    "prefix": "!",

    "testSKU": "016850",

    "general": "",
    "ownerID": "",
    "roleID": "",

    "itemURL": "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore/ProductDetail/Index?sku=",

    "url": "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore/Products/LoadProductTable",
    "host": "webapps2.abc.utah.gov",
    "origin": "https://webapps2.abc.utah.gov",
    "referer": "https://webapps2.abc.utah.gov/ProdApps/ProductLocatorCore",
    "payload": [
    ],

    "counter": 0
}
```

## Roadmap 🗺️

### Website 🕸️
- Implement a frontend using a web framework like React.
- Connect the frontend to the bot to display real-time stock updates.

### Database 💾
- Replace JSON files with a database system (e.g., PostgreSQL).
- Implement a proper data model to organize the data efficiently.

### Distribution Tracker 📈
- Keep track of the distribution of different liquors to different parts of Utah.
- Use this information to forecast future stock availabilities.

### Text Notification 📱
- Integrate with an SMTP server to send text notifications.
- Allow members to opt-in for text notifications.

### Heatmap 🗺️
- Create a heatmap to visualize where liquor is purchased the fastest.
- Use this data to understand trends and improve stock predictions.
