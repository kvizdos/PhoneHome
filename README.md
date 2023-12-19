# Phone Home

*Making HTTP Requests Feel Like They Just Called Mom!*

Phone Home is an interactive application designed to track and monitor HTTP requests, providing real-time insights and detailed request analysis. It's built in Go and uses the `tview` package for a terminal-based user interface, making it ideal for use in environments where a GUI is not available or preferred.

## Features
- Real-time Tracking: Monitor HTTP requests as they happen, and get notified using system notifications. Mom (you) never misses a call!
- Detailed Request View: View detailed information about each request, including method, user agent, and payload.
- Frequency Analysis: Like any vigilant mom, keep a close eye on how often your HTTP kids come knocking - because nobody tracks comings and goings quite like you do!
- Interactive UI: Navigate through request logs and details interactively in the terminal.
- Payload Data: Just like a meticulous mom sorting through a backpack, the '?data' parameter in GET requests can carry all sorts of items - base64 encoded or plain text. And if it's base64, consider it auto-magically decoded, because mom's always got the knack for decoding the trickiest notes!

# Installation
To install Phone Home, you need to have Go installed on your system. You can then clone this repository and build the application.

```
git clone https://github.com/kvizdos/phone-home.git
cd phone-home
go build
```

# Usage

To run Phone Home, simply execute the binary created by the build process:

```
./phone-home
```

Navigate through the request list using the arrow keys, and press 'Enter' to delve into the details of each list item. Need a quick exit from the input field? Just tap 'Escape.' Ready to jump back into command mode? Simply hit ':' and you're there!". Type "q" to go back a page.

## Creating a new Listener

Simply type `new <listener name>` and hit enter. Once you do, you'll see a unique ID (a string of 8 characters) appear next to the "- Waiting..." status. This is your listener's personal ID. To test it out, navigate to `http://localhost:4000/<id>` in your web browser. If everything's set up correctly, your listener will 'phone home' successfully!

## Mark an Event as Read

Hover over the listener you want to mark as read, and click `r` (make sure you're focused on the list by hitting escape)
