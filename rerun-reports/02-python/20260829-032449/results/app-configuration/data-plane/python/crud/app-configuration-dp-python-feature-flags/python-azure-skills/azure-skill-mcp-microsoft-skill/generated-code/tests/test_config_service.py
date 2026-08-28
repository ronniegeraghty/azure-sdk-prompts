import asyncio
import unittest
from types import SimpleNamespace
from unittest.mock import MagicMock

from azure.core import MatchConditions

from config_service import AsyncConfigurationService, ConfigurationService
from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher


def setting(key, value, etag):
    return SimpleNamespace(key=key, value=value, etag=etag)


class AsyncItemPager:
    def __init__(self, items):
        self._items = items

    def __aiter__(self):
        async def iterate():
            for item in self._items:
                yield item

        return iterate()


class AsyncPagePager(AsyncItemPager):
    def by_page(self):
        return self


class AsyncClient:
    def __init__(self):
        self.responses = []
        self.get_calls = []

    async def get_configuration_setting(self, **kwargs):
        self.get_calls.append(kwargs)
        return self.responses.pop(0)


class ConfigurationServiceTests(unittest.TestCase):
    def test_single_setting_uses_etag_after_first_read(self):
        client = MagicMock()
        client.get_configuration_setting.side_effect = [
            setting("app:key", "one", "etag-1"),
            None,
        ]
        service = ConfigurationService(client)

        self.assertEqual("one", service.get_setting("app:key"))
        self.assertEqual("one", service.get_setting("app:key"))
        second_call = client.get_configuration_setting.call_args_list[1].kwargs
        self.assertEqual("etag-1", second_call["etag"])
        self.assertEqual(MatchConditions.IfModified, second_call["match_condition"])

    def test_prefix_uses_head_etags_to_avoid_redownload(self):
        client = MagicMock()
        list_page = MagicMock()
        list_page.etag = "page-etag"
        list_page.__iter__.return_value = iter([setting("app:a", "A", "a")])
        list_pager = MagicMock()
        list_pager.by_page.return_value = [list_page]
        client.list_configuration_settings.return_value = list_pager
        check_page = SimpleNamespace(etag="page-etag")
        client.check_configuration_settings.return_value.by_page.return_value = [
            check_page
        ]
        service = ConfigurationService(client)

        expected = {"app:a": "A"}
        self.assertEqual(expected, service.list_settings("app:"))
        self.assertEqual(expected, service.list_settings("app:"))
        client.list_configuration_settings.assert_called_once()

    def test_watcher_refreshes_after_sentinel_change(self):
        service = MagicMock()
        service.get_setting.side_effect = ["v1", "v2"]
        watcher = ConfigurationWatcher(service, ["app:sentinel"], 1)

        self.assertFalse(watcher.poll_once())
        self.assertTrue(watcher.poll_once())
        service.refresh_all.assert_called_once()

    def test_async_single_setting_uses_etag(self):
        async def run():
            client = AsyncClient()
            client.responses = [
                setting("app:key", "one", "etag-1"),
                None,
            ]
            service = AsyncConfigurationService(client)
            self.assertEqual("one", await service.get_setting("app:key"))
            self.assertEqual("one", await service.get_setting("app:key"))
            self.assertEqual("etag-1", client.get_calls[1]["etag"])

        asyncio.run(run())

    def test_async_watcher_refreshes_after_sentinel_change(self):
        class Service:
            def __init__(self):
                self.values = iter(("v1", "v2"))
                self.refresh_count = 0

            async def get_setting(self, key, label=None):
                return next(self.values)

            async def refresh_all(self):
                self.refresh_count += 1

        async def run():
            service = Service()
            watcher = AsyncConfigurationWatcher(service, ["app:sentinel"], 1)
            self.assertFalse(await watcher.poll_once())
            self.assertTrue(await watcher.poll_once())
            self.assertEqual(1, service.refresh_count)

        asyncio.run(run())


if __name__ == "__main__":
    unittest.main()
