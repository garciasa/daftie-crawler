# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

from scrapy.item import Item, Field
from scrapy.loader.processors import MapCompose, TakeFirst


def clean_chars(text):
    remove = ".: "
    table = str.maketrans("", "", remove)
    text = text.translate(table)
    text = text.replace("&nbsp;", "")
    return text


class HousecrawlerItem(Item):
    url = Field(output_processor=TakeFirst())
    title = Field(output_processor=TakeFirst())
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
        clean_chars), output_processor=TakeFirst())
