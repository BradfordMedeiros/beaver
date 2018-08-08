#/usr/bin/env sh

GITHUB_URL="https://github.com/BradfordMedeiros/beaver.git"
echo "building: $GITHUB_URL"
git clone $GITHUB_URL

rm -rf beaver
