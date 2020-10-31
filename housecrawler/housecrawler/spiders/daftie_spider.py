import json
from scrapy.http.request import Request
from scrapy.spiders import Spider
from housecrawler.items import HousecrawlerItem
from scrapy.loader import ItemLoader


class DaftieSpider(Spider):
    name = "daftie"
    start_urls = ["https://www.daft.ie/wexford/houses-for-sale/wexford-town,whiterock/?pt_id=1&mnb=4&advanced=1&cc_id=c13&a_id%5B0%5D=731&a_id%5B1%5D=4273&s%5Bmnb%5D=4&s%5Badvanced%5D=1&s%5Bsort_by%5D=date&s%5Bsort_type%5D=d&searchSource=sale"]
    base_url = "https://www.daft.ie"

    def parse(self, response):
        for house in response.css("div.PropertyCardContainer__container"):
            # check is sold let or sale agreed
            link = house.css(
                "a.PropertyInformationCommonStyles__addressCopy--link::attr(href)").get()
            yield Request(self.base_url + link, callback=self.parse_link, priority=1)

        next_page = response.css('li.next_page a::attr(href)').get()
        if next_page is not None:
            yield Request(self.base_url + next_page, self.parse)

    def parse_link(self, response):
        loader = ItemLoader(item=HousecrawlerItem(), selector=response)
        loader.add_value("url", response.url)
        loader.add_value("provider", "daft.ie")
        loader.add_value("price",  response.css(
            "strong.PropertyInformationCommonStyles__costAmountCopy::text").get())
        loader.add_value("beds", response.css(
            "div.QuickPropertyDetails__iconCopy::text").get())
        loader.add_value("baths", response.css(
            "div.QuickPropertyDetails__iconCopy--WithBorder::text").get())
        propertyId = response.css(
            "a.PropertyShortcode__link::text").get()
        loader.add_value("propertyId", propertyId.replace(
            "https://www.daft.ie/", ""))
        eircode = response.xpath(
            "normalize-space(//span[contains(@class, 'PropertyMainInformation__eircodeLabel')]/following-sibling::text())").extract()
        loader.add_value("eircode", eircode)
        loader.add_value("title", response.css(
            "h1.PropertyMainInformation__address::text").get())
        loader.add_value("date_renewed", response.css(
            "div.PropertyStatistics__iconData::text").get())
        loader.add_value("first_listed", response.css(
            "div.PropertyPriceHistory__propertyPriceDate::text").get())
        loader.add_value("photo", response.css(
            "div#pbxl_carousel ul li img::attr(src)").getall()[0])
        return loader.load_item()
