#!/bin/sh

set -eu

INSTALLROOT=${INSTALLROOT:-"${HOME}/.doker"}
DOKER_VERSION=${DOKER_VERSION:-"v0.1.0"}

happyexit() {
  echo ""
  echo "Add the doker CLI to your path with:"
  echo ""
  echo "  export PATH=\$PATH:${INSTALLROOT}/bin"
  exit 0
}

OS=$(uname -s)
arch=$(uname -m)
case $OS in
  CYGWIN* | MINGW64*)
    OS=windows.exe
    ;;
  Darwin)
    ;;
  Linux)
    case $arch in
      x86_64)
        ;;
      *)
        echo "There is no doker $OS support for $arch. Please open an issue with your platform details."
        exit 1
        ;;
    esac
    ;;
  *)
    echo "There is no doker support for $OS/$arch. Please open an issue with your platform details."
    exit 1
    ;;
esac
OS=$(echo $OS | tr '[:upper:]' '[:lower:]')
arch=""
case $(uname -m) in
    i386)   arch="386" ;;
    i686)   arch="386" ;;
    x86_64) arch="amd64" ;;
esac

tmpdir=$(mktemp -d /tmp/DOKER.XXXXXX)
srcfile="doker_${DOKER_VERSION}_${OS}_${arch}"
dstfile="${INSTALLROOT}/bin/doker_${DOKER_VERSION}_${OS}_${arch}"
url="https://github.com/pothulapati/doker/releases/${DOKER_VERSION}/download/${srcfile}"

(
  cd "$tmpdir"

  echo "Downloading ${srcfile}..."
  curl -fLO "${url}"
  echo "Download complete!"

  echo ""
)

(
  mkdir -p "${INSTALLROOT}/bin"
  mv "${tmpdir}/${srcfile}" "${dstfile}"
  chmod +x "${dstfile}"
  rm -f "${INSTALLROOT}/bin/doker"
  ln -s "${dstfile}" "${INSTALLROOT}/bin/doker"
)


rm -r "$tmpdir"

echo "doker ${DOKER_VERSION} was successfully installed 🎉"
echo ""
happyexit
