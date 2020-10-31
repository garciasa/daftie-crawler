import json
from scrapy.http.request import Request
from scrapy.http import FormRequest
from scrapy.spiders import Spider
from housecrawler.items import HousecrawlerItem
from scrapy.loader import ItemLoader


class MyHomeSpider(Spider):
    name = "myhome"
    start_urls = [
        "https://api.myhome.ie/property/recentlyadded",
        # "https://www.myhome.ie/recent/wexford/wexford2",
    ]
    base_url = "https://www.myhome.ie"
    form_data = {'ApiKey': '5f4bc74f-8d9a-41cb-ab85-a1b7cfc86622',
                 'LocalitySlug': "wexford2",
                 'Page': "1"}
    brochure_url = "https://api.myhome.ie/brochure/?ApiKey=5f4bc74f-8d9a-41cb-ab85-a1b7cfc86622&SessionId=null&format=json"
    nextPages = []
    total_pages = 0
    actual_page = 1
    first = True

    def start_requests(self):
        return [FormRequest(self.start_urls[0], formdata=self.form_data, callback=self.parse)]

    def parse(self, response):
        jsonresponse = json.loads(response.text)
        self.actual_page = self.actual_page + 1

        if self.first:
            self.total_pages = int(jsonresponse['ResultCount'] / 20)

        for resp in jsonresponse['SearchResults']:
            property_id = resp['PropertyId']
            url = ("https://api.myhome.ie/brochure/%s?ApiKey=5f4bc74f-8d9a-41cb-ab85-a1b7cfc86622&SessionId=null&format=json" % (property_id,))
            self.first = False
            yield Request(url, callback=self.parse_link, priority=1)

        if self.actual_page <= self.total_pages:
            self.form_data['Page'] = str(self.actual_page)
            print(self.form_data)
            yield FormRequest(self.start_urls[0], formdata=self.form_data, callback=self.parse)

    def parse_link(self, response):
        loader = ItemLoader(item=HousecrawlerItem(), selector=response)

        jsonreponse = json.loads(response.text)
        brochure = jsonreponse['Brochure']
        property = jsonreponse['Brochure']['Property']
        type = brochure['SeoUrl'].split('/')[1]
        if type == 'residential':
            loader.add_value("provider", 'myhome.ie')
            loader.add_value("propertyId", str(property['PropertyId']))
            loader.add_value("title", property['OrderedDisplayAddress'])
            loader.add_value("url", 'https://myhome.ie' + brochure['SeoUrl'])
            loader.add_value("price", property['PriceAsString'])
            loader.add_value("photo", property['MainPhoto'])
            if 'BathString' in property.keys():
                loader.add_value("baths", property['BathString'].replace(
                    "baths", "").replace("bath", ""))
            if 'BedsString' in property.keys():
                loader.add_value("beds", property['BedsString'].replace(
                    "beds", "").replace("bed", ""))
            if property['Eircode'] != "":
                loader.add_value("eircode", property['Eircode'])
            if 'RefreshedOn' in property.keys():
                loader.add_value("date_renewed", property['RefreshedOn'])
            else:
                loader.add_value("date_renewed", property['ActivatedOn'])

            return loader.load_item()
