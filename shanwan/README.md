# Shanwan PS/4 Gamepad
I bought a PS/4 "compatible" Gamepad and while the system recognized the device `evtest` did not respond to buttons. An authentic PS/4 gamepad works both wired and wirelessly via Bluetooth. The quirks solution below restarts the gamepad in "Xbox360 For Windows" mode and it responds to `evtest`. Still not a PS/4 pad but a working improvement. May try modprobe next.

         shanwan
               |    ps/4 type
               V    V
    echo -n "2563:0119:i" | sudo tee /sys/module/usbcore/parameters/quirks
    echo -n "2563:0119:ig" | sudo tee /sys/module/usbcore/parameters/quirks

    <2563:0119:ig>
           
    i = USB_QUIRK_DEVICE_QUALIFIER
	    (device cannot handle device_qualifier descriptor requests);
    g = USB_QUIRK_DELAY_INIT
	    (device needs a pause during initialization,
		    after we read the device descriptor);

Both work, g is probably not required and does not work by itself


`dmesg | tail`

[80024.818700] hid-generic 0003:2563:0119.0025: input,hiddev1,hidraw5: USB HID v1.11 Gamepad [shanwan Wired Controller] on usb-0000:00:1d.0-1.5/input3

[80025.571774] usb 2-1.5: USB disconnect, device number 30

[80026.024984] usb 2-1.5: new full-speed USB device number 31 using ehci-pci

[80026.134650] usb 2-1.5: New USB device found, idVendor=045e, idProduct=028e bcdDevice= 1.10

[80026.134654] usb 2-1.5: New USB device strings: Mfr=1, Product=2, SerialNumber=3

[80026.134656] usb 2-1.5: Product: Xbox360 For Windows

[80026.134657] usb 2-1.5: Manufacturer: shanwan

[80026.134657] usb 2-1.5: SerialNumber: Shanwan202107142050

[80026.175608] input: Microsoft X-Box 360 pad as /devices/pci0000:00/0000:00:1d.0/usb2/2-1/2-1.5/2-1.5:1.0/input/input63

[80026.175748] usbcore: registered new interface driver xpad


