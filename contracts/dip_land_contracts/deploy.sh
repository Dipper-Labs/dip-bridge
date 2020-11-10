set -x

if [ $# -lt 2 ];
then
echo "wrong params."
echo "bash deploy.sh [dipper network addr] [dipper network addr]"
echo "e.g. bash deploy.sh dip1alqdlcgn5pff7x77gcw3hqgvktdks565vknkck dip1alqdlcgn5pff7x77gcw3hqgvktdks565vknkck"
fi

dipcli vm create --code_file=./dip_manager.bin --abi_file=./dip_manager.abi --from $1 --args "$2"  --amount=0pdip --gas=3000000 -y
