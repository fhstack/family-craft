rm -rf output
mkdir output
cp hardware/cat_lock.py output
go build -o output/server
