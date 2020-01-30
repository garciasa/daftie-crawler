# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html

import psycopg2


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
        # print("BD:"+item["url"])
        self.cur.execute(
            "insert into quotes_content(content) values(%s)", (item["url"],))
        self.connection.commit()
        return item
