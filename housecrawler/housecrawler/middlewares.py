# -*- coding: utf-8 -*-

# Define here the models for your spider middleware
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/spider-middleware.html


"""This module contains the ``SeleniumMiddleware`` scrapy middleware"""

import asyncio

from pyppeteer import launch
from scrapy import signals
from scrapy.http import HtmlResponse
from twisted.internet import defer

from housecrawler.pupeeteer import PuppeteerRequest


def _force_deferred(coro):
    dfd = defer.Deferred().addCallback(lambda f: f.result())
    future = asyncio.ensure_future(coro)
    future.add_done_callback(dfd.callback)

    return dfd


class PuppeteerMiddleware:
    """Scrapy middleware handling the requests using puppeteer"""

    browser = None

    @classmethod
    async def _from_crawler(cls, crawler):
        middleware = cls()

        middleware.browser = await launch(headless=True)
        crawler.signals.connect(
            middleware.spider_closed, signals.spider_closed)

        return middleware

    @classmethod
    def from_crawler(cls, crawler):
        """Initialize the middleware"""

        return _force_deferred(cls._from_crawler(crawler))

    async def _process_request(self, request, spider):
        """Process a request using puppeteer if applicable"""

        if not isinstance(request, PuppeteerRequest):
            return None

        page = await self.browser.newPage()
        await page.setCookie(request.cookies)
        await page.setViewport({'width': 1200, 'height': 900, 'deviceScaleFactor': 2})
        await page.goto(request.url, {'waitUntil': request.wait_until})

        if request.screenshot:
            request.meta['screenshot'] = await page.screenshot()

        body = await page.content()

        return HtmlResponse(
            page.url(),
            body=body,
            encoding='utf-8',
            request=request
        )

    def process_request(self, request, spider):
        return _force_deferred(self._process_request(request, spider))

    async def _spider_closed(self):
        await self.browser.close()

    def spider_closed(self):
        """Shutdown the driver when spider is closed"""

        return _force_deferred(self._spider_closed())
