#/usr/bin/env bash

rm -r res/plugins/github
mkdir res/plugins/github

(
	 cd ./pluginTypes/github
	./build.sh
)
cp ./pluginTypes/github/build/plugin/* res/plugins/github
cp ./pluginTypes/github/build/build-tools/* res/plugins/github
