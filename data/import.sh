mongoimport --uri mongodb+srv://dbUser:87rt45th67@cluster0.swvad.mongodb.net/bookstore --collection books --type json --jsonArray --file ./books.json

mongoimport --uri mongodb+srv://dbUser:87rt45th67@cluster0.swvad.mongodb.net/bookstore --collection authors --type json --jsonArray --file ./authors.json

mongoimport --uri mongodb+srv://dbUser:87rt45th67@cluster0.swvad.mongodb.net/bookstore --collection publishers --type json --jsonArray --file ./publishers.json

#mongoexport --uri mongodb+srv://dbUser:<PASSWORD>@cluster0.swvad.mongodb.net/<DATABASE> --collection <COLLECTION> --type <FILETYPE> --out <FILENAME>