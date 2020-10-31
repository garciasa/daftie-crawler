import json
from scrapy.http.request import Request
from scrapy.spiders import Spider
from housecrawler.items import HousecrawlerItem
from scrapy.loader import ItemLoader


class KenhoeSpider (Spider):
    name = "kenhoe"
    start_urls = [
        'https://www.kehoeproperty.com/property-search/?property_location=7']

    def parse(self, response):
        for house in response.css("article.property_cat-residential"):
            if house.xpath("//div[contains(@class,'wp-property-status')]/span/text()") not in["Let", "Sold", "Reserved", "Sale Agreed"]:
                link = house.css("div.wp-property-thumb a::attr(href)").get()
                yield Request(link, callback=self.parse_link, priority=1)

        next_page = response.css('ul.page-numbers a.next::attr(href)').get()
        if next_page is not None:
            yield Request(next_page, self.parse)

    def parse_link(self, response):
        loader = ItemLoader(item=HousecrawlerItem(), selector=response)
        loader.add_value("provider", "kehoeproperty")
        loader.add_value("url", response.url)
        loader.add_value("eircode", response.xpath(
            "normalize-space(//div[contains(@class,'site-content')])").re("([Yy]35\s?[A-Za-z\d]{4})"))
        loader.add_value("title", response.css(
            "h1.wp-property-title::text").get())
        loader.add_value("price", response.css(
            "span.wp-property-price::text").get())
        details = response.css(
            "div.wp-property-overview li.list-group-item::text").getall()
        for item in details:
            if "bedrooms" in item:
                loader.add_value("beds", item.replace("bedrooms", ""))
            if "File No." in item:
                loader.add_value("propertyId", item.replace("File No.", ""))
        loader.add_value("photo", response.xpath(
            "//ul[@id='wp-property-gallery']/li/img/@src").extract()[0])

        yield loader.load_item()
