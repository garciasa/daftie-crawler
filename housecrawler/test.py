import psycopg2

hostname = 'localhost'
username = 'postgres'  # the username when you create the database
password = ''  # change to your password
database = 'postgres'


def queryQuotes(conn):
    cur = conn.cursor()
    cur.execute("insert into quotes_content(content) values('hola')")
    conn.commit()
    cur.execute("""select * from quotes_content""")
    rows = cur.fetchall()

    for row in rows:
        print(row[1])


conn = psycopg2.connect(host=hostname, user=username,
                        password=password, dbname=database)
queryQuotes(conn)
conn.close()
