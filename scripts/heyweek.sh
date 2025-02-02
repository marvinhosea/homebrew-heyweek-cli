#!/bin/bash

installPath=$1
heyweekPath=""

if [ "$installPath" != "" ]; then
    heyweekPath=$installPath
else
  heyweekPath=/usr/local/bin
fi

UNAME=$(uname)
ARCH=$(uname -m)

removePrevInstallation() {
  if [ -f $heyweekPath/hw ]; then
      sudo rm -rf $heyweekPath/hw*
  fi
}

version=$(curl --silent "https://api.github.com/repos/marvinhosea/heyweek-cli/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
releasesApiUrl=https://github.com/marvinhosea/heyweek-cli/releases/download

installationProcess() {
  echo "Installing Heyweek CLI"
  heyweekName=""
  if [ "$UNAME" = "Linux" ]; then
    if [ $ARCH = "x86_64" ]; then
      heyweekName="Heyweek_${version}_linux_x86_64"
    elif [ $ARCH = "i386" ]; then
      heyweekName="Heyweek_${version}_linux_i386"
    elif [ $ARCH = "arm64" ]; then
      heyweekName="Heyweek_${version}_linux_arm64"
    fi

    heyweekURL=$releasesApiUrl/$version/$heyweekName.tar.gz
    wget $heyweekURL
    sudo chmod 755 $heyweekName.tar.gz
    tar -xzf $heyweekName.tar.gz
    sudo mv $heyweekName/hw $heyweekPath

    rm -rf $heyweekName
  elif [ "$UNAME" == "Darwin" ]; then
    if [ $ARCH = "x86_64" ]; then
      heyweekName="Heyweek_${version}_macOS_amd64"
    elif [ $ARCH = "arm64" ]; then
      heyweekName="Heyweek_${version}_macOS_arm64"
    fi

    heyweekURL=$releasesApiUrl/$version/$heyweekName.zip

    wget $heyweekURL
    sudo chmod 755 $heyweekName.zip
    unzip $heyweekName.zip
    rm $heyweekName.zip
    sudo mv $heyweekName/hw $heyweekPath

    rm -rf $heyweekName
  fi
}

removePrevInstallation
installationProcess

if [ -x "$(command -v hw)" ]; then
    echo "Heyweek CLI installed successfully"
else
  echo "Heyweek CLI failed to install."
  echo "Please try again"
fi
