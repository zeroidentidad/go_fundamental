#!/bin/bash


# Build ArchLinux
arch_build() {
  cp go-asisten-gui.desktop arch;
  cd arch;
  makepkg -cs;
}

# Build Ubuntu
ubuntu_build() {
  cp go-asisten-gui.desktop ubuntu;
  #cp any/anybuild.sh ubuntu;
  cd ubuntu;
  ./build.sh;
}

# Build any
any_build() {
  cp go-asisten-gui.desktop any;
  cd any
  ./anybuild.sh;
}

case $1 in
arch)
    arch_build;
    ;;
ubuntu)
    ubuntu_build
    ;;
any)
    any_build
    ;;
*)
    echo "Set Distribution example './build.sh arch'"
    ;;
esac