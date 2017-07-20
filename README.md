## golight

A [go](http://golang.org) package for turning a [Delcom](http://www.delcomproducts.com/products_usblmp.asp) led light red, green, or yellow.

### Installation
This project depends on [go.hid](https://github.com/karalabe/hid) which is a wrapper for [signal11/hidapi](https://github.com/signal11/hidapi).

To install hidapi:

```shell
git clone git@github.com:signal11/hidapi.git
cd hidapi
./bootstrap
./configure
```

On OSX I found that bootstrap required automake and other packages.

```shell
sudo port install automake autoconf libtool
```
For linux + hidraw: `cd linux`. (requires libudev. Package libudev-dev on debian/ubuntu.)

For linux + libusb: `cd libusb`. (requires libusb. Package libusb-1.0-0-dev on debian/ubuntu.)

For mac: `cd mac`.

For windows: `cd windows`.

Make and install.
For linux/mac:
```
make
sudo make install
```

Lastly, for linux only:
Create a symlink pointing libhidapi.so to the version you chose:

For linux + hidraw: `cd /usr/local/lib; sudo ln -s libhidapi-hidraw.so libhidapi.so`

For linux + libusb: `cd /usr/local/lib; sudo ln -s libhidapi-libusb.so libhidapi.so`

For more instructions on libhidapi, please visit [signal11/hidapi](https://github.com/signal11/hidapi).

When you have installed hidapi lib, install go.hid package with `go get github.com/GeertJohan/go.hid`.

Then install this package with `go get github.com/erichhansen/golight`

### Usage
There are 3 main methods - On, Off, and Flash.
- **On** Turn light on
- **Off** Turn light off
- **Flash** Make light flash

For more information communicating with the delcom light refer the the [USB IO HID Datasheet](http://www.delcomproducts.com/downloads/USBIOHID.pdf).
