#/usr/bin/env sh

mkdir build
(
	cd ../common/hooker
	./build.sh
)

cp -r ./src/* ./build/
cp ../common/hooker/build/hooker ./build/plugin
