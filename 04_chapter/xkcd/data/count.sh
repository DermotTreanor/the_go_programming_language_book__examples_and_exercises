#!/usr/bin/bash
#
FILE='./comics.json'
lines=$(wc -l $FILE | grep --only-matching -P '^\d+')
finds=$(grep --only-matching -n -P 'num\":\s\d+' $FILE | wc -l) 
if [[ lines -eq finds ]];then
	echo "The LINES ($lines) are the same as the FINDS ($finds). TRUE"
else 
	echo "LINES ($lines) =/= FINDS ($finds). FALSE"
fi
