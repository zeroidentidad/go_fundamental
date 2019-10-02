#!/bin/bash
echo "ANY-BUILD: R2"
nowfolder="$(pwd)/"
pkgdir="$(pwd)/pkg";
srcdir="$(pwd)/src";
mkdir -p $srcdir;
mkdir -p $pkgdir;
pkgname="go-asisten"

build() {
  cd "$srcdir"
  git clone git@gitlab.com:AndrusGerman/go-asisten-core.git
  git clone git@gitlab.com:AndrusGerman/go-asisten-gui.git
  # Go AsistenCore Build
  cd "$pkgname-core"
  go build -buildmode=plugin
  # Go AsistenGUI build
  cd "$srcdir/$pkgname-gui"
  go build
  # Go Asisten icon
  cp "$nowfolder/$pkgname-gui.desktop" $srcdir;
}


package () {
  cd "$pkgdir"
  mkdir -p "opt/$pkgname"

  # GoAsistenGUI package
  cp -Rv "$srcdir/$pkgname-gui" "$pkgdir/opt/$pkgname"

  # GoAsistenCORE package
  cp -Rv "$srcdir/$pkgname-core" "$pkgdir/opt/$pkgname"

  # GoAsistenGUI Launcher
  install -Dm644 "$srcdir/$pkgname-gui.desktop"    "$pkgdir/usr/share/applications/$pkgname-gui.desktop"

  # GoAsistenGUI binary Launch
  mkdir -p "usr/bin"
  ln -s "/opt/$pkgname/$pkgname-gui/$pkgname-gui" "$pkgdir/usr/bin/$pkgname-gui"
}
build
package
#Clear
rm -rf "$srcdir"