# Telmo

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

### Advertisement

If you are looking for a removable eUICC solution, such as eSIM.me or 5ber.eSIM, I recommend buying [eSTK.me](https://www.estk.me?aid=esim). It is more reliable and offers more features. You can purchase it [here](https://www.estk.me?aid=esim), and you can use the coupon code `eSIMCyou` to get a `5%` discount.

### Introduction

This program enables you to manage your SMS via a Telegram bot, including tasks such as receiving and sending SMS. To use this program, you will need a Telegram account and a Telegram bot. You can create a Telegram bot by interacting with [BotFather](https://t.me/botfather) and following the provided instructions.

I have thoroughly tested this program and found it to work well. However, its compatibility with your system may vary. Should you encounter any issues, please do not hesitate to inform me.

### Tested Devices

* Qualcomm 410 WiFi Dongle
* Quectel EC20
* Quectel EM12-G

### Requirements

* *ModemManager*: Essential for managing modems.
* *libqmi*: Version 1.35.5 or higher is required. (for QMI modems)
* *libmbim*: Version 1.26.0 or higher is required. (for MBIM modems)

### Installation & Usage

You can obtain the latest release from the [releases page](https://github.com/damonto/telmo/releases).

If you want to build it yourself, you can run the following commands:

```bash
git clone git@github.com:damonto/telmo.git
cd telmo
go build -trimpath -ldflags="-w -s" -o telmo main.go
```

Sometimes, you might need to set executable permissions for the binary file using the following command:

```bash
chmod +x telmo
```

Once done, you can run the program with root privileges:

```bash
sudo ./telmo --bot-token=YourTelegramToken --admin-id=YourTelegramChatID
```

#### QMI

`libqmi` requires you to have `libqmi` version **1.35.5** or later installed. If your package manager doesn't provide the latest version, you can build it from source. The following instructions are for building `libqmi` using `meson` and `ninja`. Make sure you have the necessary dependencies installed on your system.

For detailed build instructions, refer to [libqmi's official documentation](https://modemmanager.org/docs/libqmi/building/building-meson/).

```bash
# sudo pacman -S --needed meson ninja pkg-config bash-completion libgirepository help2man glib2 libgudev libmbim libqrtr-glib (Arch Linux)
# sudo apt-get install -y meson ninja-build pkg-config bash-completion libgirepository1.0-dev help2man libglib2.0-dev libgudev-1.0-dev libmbim-glib-dev libqrtr-glib-dev (Ubuntu/Debian)
git clone https://gitlab.freedesktop.org/mobile-broadband/libqmi.git
cd libqmi
meson setup build --prefix=/usr --buildtype=release
sudo ninja -j$(nproc) -C build
sudo ninja -C build install
```

Once you have compiled and installed `libqmi`, you can run the program with the following command:

```bash
sudo ./telmo --bot-token=YourTelegramToken --admin-id=YourTelegramChatID
```

If you wish to run the program in the background, you can utilize the `systemctl` command. Here is an example of how to achieve this:

1. Start by creating a service file in the `/etc/systemd/system` directory. For instance, you can name the file `telmo.service` and include the following content:

```plaintext
[Unit]
Description=Telegram Mobile
After=network.target

[Service]
Type=simple
User=root
Restart=on-failure
ExecStart=/your/binary/path/here/telmo --bot-token=YourTelegramToken --admin-id=YourTelegramChatID
RestartSec=10s
TimeoutStopSec=30s

[Install]
WantedBy=multi-user.target
```

2. Then, use the following command to start the service:

```bash
sudo systemctl start telmo
```

3. If you want the service to start automatically upon system boot, use the following command:

```bash
sudo systemctl enable telmo
```
