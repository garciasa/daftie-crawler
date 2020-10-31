# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

from scrapy.item import Item, Field
from scrapy.loader.processors import MapCompose, TakeFirst
import datetime


def clean_chars(text):
    remove = ".: "
    table = str.maketrans("", "", remove)
    text = text.translate(table)
    text = text.replace("&nbsp;", "")
    text = text.replace(u'\xa0', u'')
    return text


def clean_date(text):
    text = text.replace("\n", "")
    return text


class HousecrawlerItem(Item):
    url = Field(output_processor=TakeFirst())
    title = Field(input_processor=MapCompose(
        str.strip), output_processor=TakeFirst())
    price = Field(input_processor=MapCompose(
        str.strip, clean_chars), output_processor=TakeFirst())
    beds = Field(input_processor=MapCompose(
        str.strip), output_processor=TakeFirst())
    baths = Field(input_processor=MapCompose(
        str.strip), output_processor=TakeFirst())
    propertyId = Field(input_processor=MapCompose(
        str.strip), output_processor=TakeFirst())
    provider = Field(output_processor=TakeFirst())
    eircode = Field(input_processor=MapCompose(
        str.strip, clean_chars), output_processor=TakeFirst())
    date_renewed = Field(input_processor=MapCompose(
        str.strip, clean_date), output_processor=TakeFirst())
    first_listed = Field(input_processor=MapCompose(
        str.strip, clean_date), output_processor=TakeFirst())
    photo = Field(output_processor=TakeFirst())
