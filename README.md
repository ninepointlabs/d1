# d1 - Day One Journal CLI (Linux)

A lightweight Linux CLI and TUI for the Day One journaling platform, written in Go. 

Since Day One does not provide a native Linux client or a public API for personal use, this tool bridges the gap by leveraging Day One's **"Email to Journal"** feature. It sends your entries securely via your own Gmail account using SMTP.

## Features

- **Quick Post**: One-line command to post immediately (`d1 "Had a great coffee today"`)
- **Compose Mode**: specific full-screen TUI (Terminal User Interface) for distraction-free writing.
- **Secure Auth**: Credentials are stored securely in your system's keyring (GNOME Keyring, KWallet, etc.), not in plain text.
- **Offline-ish**: If you are online, it sends immediately. (True offline queuing is planned for future versions).

## Installation

Download the latest release for your distribution from the [Releases Page](https://github.com/timappledotcom/d1/releases).

### Debian / Ubuntu
```bash
sudo dpkg -i d1_0.1.0_linux_amd64.deb
```

### Fedora / RHEL
```bash
sudo rpm -i d1_0.1.0_linux_amd64.rpm
```

### Arch Linux
```bash
sudo pacman -U d1_0.1.0_linux_amd64.pkg.tar.zst
```

## Setup Guide

Before using `d1`, you need to gather two pieces of information.

### 1. Get a Gmail App Password
Standard Google passwords do not work with third-party apps like this. You need a dedicated App Password.
1. Go to your [Google Account via Security](https://myaccount.google.com/security).
2. Under "How you sign in to Google", select **2-Step Verification**.
3. Scroll to the bottom and select **App passwords**.
4. Create a new one named "DayOneCLI".
5. **Copy the 16-character code** (e.g., `abcd efgh ijkl mnop`).

### 2. Get Your Day One Email
1. Open the Day One app on your phone or Mac.
2. Go to **Settings > Journal**.
3. Select the journal you want to post to.
4. Look for **"Email to Journal"**.
5. Copy the address (e.g., `journal-abc123xyz@dayone.me`).

### 3. Configure d1
Run the interactive configuration wizard:

```bash
d1 config
```

Follow the prompts to enter your Gmail address, the App Password you just generated, and your Day One email address. These will be securely saved to your system keychain.

## Usage

### Quick Post
Great for short logs or fleeting thoughts.
```bash
d1 "Just arrived at the station. #travel"
```

### Compose Mode
Run without arguments to open the TUI editor.
```bash
d1
```
- Write your entry.
- Press **Ctrl+S** to send.
- Press **Esc** to discard and quit.

## Troubleshooting

- **"Credentials missing"**: Run `d1 config` again.
- **"Application password required"**: Ensure you are using an App Password, not your login password.
- **Entry not appearing**: Check your "Sent" folder in Gmail. If it's there, check Day One. It can take a minute for the email to be processed by Day One's servers.

## License
MIT
