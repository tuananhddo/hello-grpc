testfunction(){

   echo "My first function"
   echo $1
}

testfunction "Hmm"

get_directory() {
    dirname $(echo $1 | sed "s/^$PROTO_DIR\///g")
}

echo $?