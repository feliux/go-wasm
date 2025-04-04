
cut -d, -f2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,31 creditcard.csv > creditcard_class.csv
sed -e 's/\"//g' creditcard_class.csv > clean_creditcard_class.csv

cut -d, -f2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29 creditcard.csv > creditcard_noclass.csv
sed -e 's/\"//g' creditcard_noclass.csv > clean_creditcard_noclass.csv

rm creditcard_class.csv
rm creditcard_noclass.csv
