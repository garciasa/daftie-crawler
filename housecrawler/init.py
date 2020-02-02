import logging
from scrapy.crawler import CrawlerProcess
from scrapy.utils.project import get_project_settings
from scrapy.utils.log import configure_logging
from housecrawler.spiders.daftie_spider import DaftieSpider
from housecrawler.spiders.myhome_spider import MyHomeSpider
from housecrawler.spiders.kenhoe_spider import KenhoeSpider
from housecrawler.spiders.keane_spider import KeaneSpider

configure_logging(install_root_handler=False)
logging.basicConfig(
    filename='log.txt',
    format='%(levelname)s: %(message)s',
    level=logging.INFO,
    filemode='w'

)


def start_sequentially(process: CrawlerProcess, crawlers: list):
    print('start crawler {}'.format(crawlers[0].__name__))
    deferred = process.crawl(crawlers[0])
    if len(crawlers) > 1:
        deferred.addCallback(
            lambda _: start_sequentially(process, crawlers[1:]))


crawlers = [MyHomeSpider, DaftieSpider, KenhoeSpider, KeaneSpider]
process = CrawlerProcess(settings=get_project_settings())
start_sequentially(process, crawlers)
process.start()
