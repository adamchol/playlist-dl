#!/usr/bin/env bash

package_name="playlist-dl"
	
platforms=("windows/amd64" "windows/386" "darwin/amd64" "darwin/arm64" "linux/amd64" "linux/arm64" "linux/arm")

rm -rf build/ && mkdir build
for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	output_name=$package_name
	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi	

	env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name
	if [ $? -ne 0 ]; then
   		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
	system=$GOOS
	if [ $GOOS = "darwin" ]; then
		system="macos"
	fi
	zip_name=$package_name'-'$system'-'$GOARCH
	zip_name+='.zip'
	zip $zip_name $output_name
	rm $output_name
	mv $zip_name build/
done
