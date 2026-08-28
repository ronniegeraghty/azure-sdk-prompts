import asyncio
import unittest

from config_watcher import AsyncConfigurationWatcher, ConfigurationWatcher


class _SyncConfig:
    def __init__(self):
        self.values = iter(("v1", "v2"))
        self.refreshes = 0

    def get_setting(self, key, label=None):
        return next(self.values)

    def refresh_all(self):
        self.refreshes += 1


class _AsyncConfig:
    def __init__(self):
        self.values = iter(("v1", "v2"))
        self.refreshes = 0

    async def get_setting(self, key, label=None):
        return next(self.values)

    async def refresh_all(self):
        self.refreshes += 1


class WatcherTests(unittest.TestCase):
    def test_sync_change_refreshes_cache(self):
        config = _SyncConfig()
        watcher = ConfigurationWatcher(config, ["Sentinel"], 0.001)
        watcher.run(max_polls=1)
        self.assertEqual(config.refreshes, 1)

    def test_async_change_refreshes_cache(self):
        async def run_test():
            config = _AsyncConfig()
            watcher = AsyncConfigurationWatcher(config, ["Sentinel"], 0.001)
            await watcher.run(max_polls=1)
            self.assertEqual(config.refreshes, 1)

        asyncio.run(run_test())


if __name__ == "__main__":
    unittest.main()
