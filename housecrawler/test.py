import psycopg2
import uuid
from datetime import datetime


hostname = 'localhost'
username = 'postgres'  # the username when you create the database
password = ''  # change to your password
database = 'postgres'


def queryQuotes(conn):
    item = {}
    item["date_renewed"] = "01.12.2020"
    item["first_listed"] = "Jan 1, 2020"
    item["url"] = "https://www.daft.ie/wexford/houses-for-sale/wexford-town/3-temperance-row-selskar-wexford-town-wexford-2369172/"
    item["price"] = "€95,000"
    item["title"] = "3 Temperance Row, Selskar, Wexford Town, Co. Wexford"
    item["beds"] = "3"
    item["baths"] = "4"
    item["provider"] = "asdf"
    item["eircode"] = "asdf"
    item["propertyId"] = "12344556555"

    cur = conn.cursor()
    id = str(uuid.uuid4())
    datefl = None
    if (item["date_renewed"]):
        dater = datetime.strptime(
            item["date_renewed"], "%d.%m.%Y")
    if ("first_listed" in item):
        datefl = datetime.strptime(item["first_listed"], "%b %d, %Y")
    print(item)
    cur.execute(
        "insert into house(id, url, price, title, beds, baths, provider, eircode, date_renewed, first_listed, propertyid) values(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
        (id, item["url"], item["price"], item["title"], int(item["beds"]), int(item["baths"]), item["provider"], item["eircode"], dater, datefl, item["propertyId"],))
    conn.commit()
    cur.execute("""select * from house""")
    rows = cur.fetchall()

    for row in rows:
        print(row[1])


def myquery(conn):
    item = {}
    item["propertyId"] = "12343350"
    cur = conn.cursor()
    cur.execute(
        """ select * from house where propertyid = %s """, (item["propertyId"],))
    rows = cur.fetchall()

    for row in rows:
        print(row[1])


conn = psycopg2.connect(host=hostname, user=username,
                        password=password, dbname=database)
# queryQuotes(conn)
myquery(conn)
conn.close()
