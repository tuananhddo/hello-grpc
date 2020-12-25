echo  $1 $2 $3 $4 $5 $6 $7 $8 $9 ${10} ${11}
# Tất cả các đối số cung cấp cho file  Bash script
echo $#
# shellcheck disable=SC2068
# Số lượng các  arguments chúng ta truyền vào cho file the Bash script.
echo $@
# ID của script hiện tại .
echo $$
#Trạng thái của câu lệnh thực hiện gần nhất ( 0 -> true , 1 -> false )
echo $?
