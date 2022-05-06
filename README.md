# Reset BOOTSEL via Serial

Open and close at 1200 bps to the serial port of the Raspberry Pi Pico to reset in BOOTSEL mode.
The Arduino CLI has an implementation that does this reset, so this tool calls it.

## install

```
go install github.com/74th/reset-bootsel-via-serial-tool/cmd/reset-bootsel@latest
```

### how to use

```
# Windows
reset-bootsel COM8

# Linux
reset-bootsel /dev/ttyACM1

# MacOS
reset-bootsel /dev/tty.usbmodem1234
```
