import json
from scrapy.http.request import Request
from scrapy.spiders import Spider
from housecrawler.items import HousecrawlerItem
from scrapy.loader import ItemLoader


class MyHomeSpider(Spider):
    name = "myhome"
    start_urls = [
        "https://www.myhome.ie/recent/wexford/wexford2",
    ]
    base_url = "https://www.myhome.ie"
    nextPages = []
    first = True

    def parse(self, response):

        if self.first:
            # Try to generate pages to visit because they dont adding reference to pagination
            count = response.css("li.small-screen::text").get()
            pages = int(count.rsplit('/', 1)[1])

            for i in range(2, pages+1):
                url = (
                    "https://www.myhome.ie/recent/wexford/wexford2?page=%s" % (i,))
                self.nextPages.append(url)

            self.first = False

        for house in response.css("div.RecentlyAdded__Property"):
            link = house.css(
                "a.RecentlyAddedPropertyCard__Address::attr(href)").get()
            yield Request(self.base_url + link, callback=self.parse_link, priority=1)

        for page in self.nextPages:
            yield Request(page, self.parse)

    def parse_link(self, response):

        loader = ItemLoader(item=HousecrawlerItem(), selector=response)
        loader.add_value("url", response.url)
        loader.add_value("provider", "myhome.ie")
        loader.add_value("price",  response.css(
            "div.PropertyBrochure__Price::text").get())

        details = response.css("span.PropertyInfoStrip__Detail::text").getall()
        for i in details:
            if "beds" in i:
                loader.add_value("beds", i.replace("beds", ""))
            if "bath" in i:
                loader.add_value("baths", i.replace(
                    "baths", "").replace("bath", ""))
            if "Refreshed" in i:
                loader.add_value(
                    "date_renewed", i.replace("Refreshed on ", ""))

        loader.add_value("propertyId", response.url.rsplit('/', 1)[1])
        eircode = response.css("div.PropertyBrochure__Eircode::text").get()
        if eircode is not None:
            loader.add_value("eircode", eircode.replace("Eircode: ", ""))

        loader.add_value("title", response.css(
            "h1.PropertyBrochure__Address::text").get())
        return loader.load_item()
