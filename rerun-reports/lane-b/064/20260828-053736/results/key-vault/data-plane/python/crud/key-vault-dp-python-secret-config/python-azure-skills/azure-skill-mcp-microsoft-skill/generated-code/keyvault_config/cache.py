from __future__ import annotations

from collections.abc import Callable, Iterable, Mapping
from datetime import datetime, timedelta, timezone

from .models import SecretInfo
from .provider import SecretProvider

RequiredSecrets = Iterable[str] | Mapping[str, str | None]


class _Unset:
    pass


_UNSET = _Unset()


class SecretCache:
    def __init__(
        self,
        provider: SecretProvider,
        *,
        warning_window: timedelta = timedelta(days=7),
        clock: Callable[[], datetime] = lambda: datetime.now(timezone.utc),
    ) -> None:
        if warning_window < timedelta(0):
            raise ValueError("warning_window cannot be negative")
        self._provider = provider
        self._warning_window = warning_window
        self._clock = clock
        self._entries: dict[str, SecretInfo] = {}
        self._defaults: dict[str, str | None] = {}

    def load_required(self, required: RequiredSecrets) -> None:
        items = required.items() if isinstance(required, Mapping) else (
            (name, None) for name in required
        )
        for name, default in items:
            self.refresh(name, default=default)

    def get(self, name: str, default: str | None = None) -> str | None:
        entry = self._entries.get(name)
        if entry is None or self._is_near_expiry(entry):
            entry = self.refresh(name, default=default)
        return entry.value

    def refresh(
        self, name: str, default: str | None | _Unset = _UNSET
    ) -> SecretInfo:
        if not isinstance(default, _Unset):
            self._defaults[name] = default
        fallback = self._defaults.get(name)
        entry = self._provider.get_info(name, fallback)
        self._entries[name] = entry
        return entry

    def refresh_near_expiry(self) -> list[str]:
        refreshed: list[str] = []
        for name, entry in list(self._entries.items()):
            if self._is_near_expiry(entry):
                self.refresh(name)
                refreshed.append(name)
        return refreshed

    def expiring_keys(self) -> list[str]:
        return sorted(
            name for name, entry in self._entries.items()
            if self._is_near_expiry(entry)
        )

    def _is_near_expiry(self, entry: SecretInfo) -> bool:
        return entry.expires_within(self._clock() + self._warning_window)
