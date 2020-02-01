#!/bin/sh

set -eu

INSTALLROOT=${INSTALLROOT:-"${HOME}/.doker"}

happyexit() {
  echo ""
  echo "Add the doker CLI to your path with:"
  echo ""
  echo "  export PATH=\$PATH:${INSTALLROOT}/bin"
  exit 0
}

validate_checksum() {
  filename=$1
  SHA=$(curl -sfL "${url}.sha256")
  echo ""
  echo "Validating checksum..."

  case $checksumbin in
    *openssl)
      checksum=$($checksumbin dgst -sha256 "${filename}" | sed -e 's/^.* //')
      ;;
    *shasum)
      checksum=$($checksumbin -a256 "${filename}" | sed -e 's/^.* //')
      ;;
  esac

  if [ "$checksum" != "$SHA" ]; then
    echo "Checksum validation failed." >&2
    return 1
  fi
  echo "Checksum valid."
  return 0
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

checksumbin=$(command -v openssl) || checksumbin=$(command -v shasum) || {
  echo "Failed to find checksum binary. Please install openssl or shasum."
  exit 1
}


tmpdir=$(mktemp -d /tmp/DOKER.XXXXXX)
srcfile="doker_${OS}_${arch}"
dstfile="${INSTALLROOT}/bin/doker_${OS}_${arch}"
url="https://github.com/pothulapati/doker/releases/latest/download/${srcfile}"

if [ -e "${dstfile}" ]; then
  if validate_checksum "${dstfile}"; then
    echo ""
    echo "doker was already downloaded; making it the default 🎉"
    echo ""
    echo "To force re-downloading, delete '${dstfile}' then run me again."
    (
      rm -f "${INSTALLROOT}/bin/doker"
      ln -s "${dstfile}" "${INSTALLROOT}/bin/doker"
    )
    happyexit
  fi
fi

(
  cd "$tmpdir"

  echo "Downloading ${srcfile}..."
  curl -fLO "${url}"
  echo "Download complete!"

  if ! validate_checksum "${srcfile}"; then
    exit 1
  fi
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
