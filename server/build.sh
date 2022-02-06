rm -rf output
mkdir output
cp hardware/cat_lock_api.py output
go build -o output/server
