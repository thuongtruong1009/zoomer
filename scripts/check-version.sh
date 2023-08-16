CURRENT_VERSION="$(go version | { read -r _ _ v _; echo "${v#go}"; })"
REQUIRED_VERSION="1.20"

if [ "$(printf '%s\n' "$REQUIRED_VERSION" "$CURRENT_VERSION" | sort -V | head -n1)" = "$REQUIRED_VERSION" ]; then
      echo "1"
else
      echo "0"
fi
