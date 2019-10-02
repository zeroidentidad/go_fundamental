#!/bin/bash
pkgname="go-asisten"
nowfolder="$(pwd)/"
echo "UBUNTU-BUILD: R1"
# BUild go-asisten
./anybuild.sh;
# Create Basi Files
mkdir -p "$pkgname/DEBIAN/"
cp control "$pkgname/DEBIAN/"
mv pkg/* "$pkgname"
# permisos
sudo chown root:root -R "$pkgname"
sudo chmod 0755 "$pkgname/opt/$pkgname/$pkgname-gui/$pkgname-gui"
# Build To ubuntu
dpkg -b "$pkgname"
# Clear
# sudo rm -rf "$pkgname";