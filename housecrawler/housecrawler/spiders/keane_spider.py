import json
from scrapy.http.request import Request
from scrapy.spiders import Spider
from housecrawler.items import HousecrawlerItem
from scrapy.loader import ItemLoader


class KeaneSpider(Spider):
    name = 'keane'
    start_urls = ['https://keaneauctioneers.com/location/wexford-town/']

    def parse(self, response):
        for link in response.css(".listing_wrapper"):
            if (link.css(".inforoom::text").get() is not None and
                link.css(".ribbon-inside.Sale-Agreed").get() is None and
                    link.css(".ribbon-inside.Sold").get() is None):
                # house
                ref = link.css("h4>a::attr(href)").get()
                yield Request(ref, callback=self.parse_link, priority=1)

        for next_page in response.css('.pagination>li>a'):
            yield response.follow(next_page, self.parse)

    def parse_link(self, response):
        loader = ItemLoader(item=HousecrawlerItem(), selector=response)
        loader.add_value("provider", "keaneauctioneers")
        loader.add_value("url", response.url)
        loader.add_value("photo", response.css(
            "div#property_slider_carousel div.item img::attr(src)").get())
        loader.add_value("title", response.css(
            ".entry-title.entry-prop::text").get())

        eircode = response.xpath(
            "//div[contains(@class, 'wpestate_property_description')]//strong/following-sibling::text()").get()
        if (eircode is not None):
            loader.add_value("eircode", response.xpath(
                "//div[contains(@class, 'wpestate_property_description')]//strong/following-sibling::text()")[-1].extract())

        for items in response.xpath("//*[@id='collapseOne']//div"):
            if (items.xpath(".//strong/text()").extract()[0] == "Price:"):
                loader.add_value("price", items.xpath(
                    ".//text()").extract()[1])
            if (items.xpath(".//strong/text()").extract()[0] == "Bedrooms:"):
                loader.add_value("beds", items.xpath(".//text()").extract()[1])
            if (items.xpath(".//strong/text()").extract()[0] == "Property ID:"):
                loader.add_value(
                    "propertyId", items.xpath(".//text()").extract()[1])
            if (items.xpath(".//strong/text()").extract()[0] == "Bathrooms:"):
                loader.add_value(
                    "baths", items.xpath(".//text()").extract()[1])

        yield loader.load_item()
