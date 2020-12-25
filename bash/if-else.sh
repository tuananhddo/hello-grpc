
if [ $1 -ge 18 ]; then
  echo You may go to the party.
elif [ $2 == 'yes' ]; then
  echo You may go to the party but be back before midnight.
else
  echo You may not go to the party.
fi
<<COMMENT
//Number
  num1 -eq num2                  check if 1st  number is equal to 2nd number
  num1 -ge num2                  checks if 1st  number  is greater than or equal to 2nd number
  num1 -gt num2                  checks if 1st  number is greater than 2nd number
  num1 -le num2                   checks if 1st number is less than or equal to 2nd number
  num1 -lt num2                   checks if 1st  number  is less than 2nd number
  num1 -ne num2                  checks if 1st  number  is not equal to 2nd number
  // String
  var1 = var2     checks if var1 is the same as string var2
  var1 != var2    checks if var1 is not the same as var2
  var1 < var2     checks if var1 is less than var2
  var1 > var2     checks if var1 is greater than var2
  -n var1             checks if var1 has a length greater than zero
  -z var1             checks if var1 has a length of zero
  //File
  -d file                        checks if the file exists and is it’s a directory
  -e file                        checks if the file exists on system
  -w file                       checks if the file exists on system and if it is writable
  -r file                        checks if the file exists on system and it is readable
  -s file                        checks if the file exists on system and it is not empty
  -f file                         checks if the file exists on system and it is a file
  -O file                       checks if the file exists on system and if it’s is owned by the current user
  -G file                        checks if the file exists and the default group is the same as the current user
  -x file                         checks if the file exists on system and is executable
  file A -nt file B         checks if file A is newer than file B
  file A -ot file B          checks if file A is older than file B
COMMENT