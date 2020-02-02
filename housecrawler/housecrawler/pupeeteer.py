"""This module contains the ``SeleniumRequest`` class"""

from scrapy import Request


class PuppeteerRequest(Request):
    """Scrapy ``Request`` subclass providing additional arguments"""

    def __init__(self, screenshot=False, *args, **kwargs):
        """Initialize a new Puppeteer request

        Parameters
        ----------
        screenshot: bool
            If True, a screenshot of the page will be taken and the data of the screenshot
            will be returned in the response "meta" attribute.

        """

        self.screenshot = screenshot

        super().__init__(*args, **kwargs)
