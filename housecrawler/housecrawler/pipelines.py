# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html

import psycopg2
from datetime import datetime
import uuid
import re


class HousecrawlerPipeline(object):
    def open_spider(self, spider):
        hostname = 'localhost'
        username = 'postgres'
        password = ''  # your password
        database = 'postgres'
        self.connection = psycopg2.connect(
            host=hostname, user=username, password=password, dbname=database)
        self.cur = self.connection.cursor()

    def close_spider(self, spider):
        self.cur.close()
        self.connection.close()

    def process_item(self, item, spider):
        # Check if it already exists base on eircode
        if ("eircode" in item):
            eircode = item["eircode"]
            pattern = re.compile("([Yy]35\s?[A-Za-z\d]{4})")
            if pattern.match(eircode):
                # valid eircode
                result = self.check_eircode(eircode)
                if not result:
                    self.create(item)
            else:
                result = self.check_code(item["propertyId"])
                if not result:
                    self.create(item)
        else:
            result = self.check_code(item["propertyId"])
            if not result:
                self.create(item)

        return item

    def check_eircode(self, eircode):
        # checking eircode
        result = False
        self.cur.execute(
            """ select * from house where eircode = %s """, (eircode,))
        rows = self.cur.fetchall()

        result = len(rows) > 0

        return result

    def check_code(self, propertyId):
        # checking_agencycode
        result = False
        self.cur.execute(
            """ select * from house where propertyid = %s """, (propertyId,))
        rows = self.cur.fetchall()

        result = len(rows) > 0

        return result

    def create(self, item):
        id = str(uuid.uuid4())
        datefl = None
        dater = None
        if ("date_renewed" in item):
            try:
                dater = datetime.strptime(
                    item["date_renewed"], "%d.%m.%Y")
            except:
                # for myhome date comes in this format
                dater = datetime.strptime(item["date_renewed"], "%b %d, %Y")

        if ("first_listed" in item):
            datefl = datetime.strptime(item["first_listed"], "%b %d, %Y")

        beds = 0
        if ("beds" in item):
            beds = item["beds"]

        baths = 0
        if ("baths" in item):
            baths = item["baths"]
        eircode = ""

        if ("eircode" in item):
            eircode = item["eircode"]
            pattern = re.compile("([Yy]35\s?[A-Za-z\d]{4})")
            if not pattern.match(eircode):
                eircode = ""

        price = ""
        if ("price" in item):
            price = item["price"]

        self.cur.execute(
            "insert into house(id, url, price, title, beds, baths, provider, eircode, date_renewed, first_listed, propertyid) values(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)",
            (id, item["url"], price, item["title"], beds, baths, item["provider"], eircode, dater, datefl, item["propertyId"],))

        self.connection.commit()
