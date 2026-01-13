# d1 - Day One Journal CLI (Linux)

A lightweight Linux CLI and TUI for the Day One journaling platform, written in Go.

## Features

- **Quick Post**: `d1 "My thought"`
- **Compose Mode**: Interactive TUI (Bubble Tea)
- **Secure Auth**: Uses system keyring for Gmail credentials
- **Email Integration**: Posts via Day One's "Email to Journal" feature

## Installation

### Arch Linux
```bash
yay -S d1-bin
```

### Debian/Ubuntu
Download the `.deb` from releases and run:
```bash
sudo dpkg -i d1_*.deb
```

## Configuration

Run the setup wizard:
```bash
d1 config
```
