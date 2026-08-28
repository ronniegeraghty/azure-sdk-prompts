import unittest

from order_processing.model import Order, OrderStatus


class OrderTests(unittest.TestCase):
    def test_json_round_trip(self) -> None:
        order = Order("o-1", "Ada", "Keyboard", 2, 199.98)

        restored = Order.from_json(order.to_json())

        self.assertEqual(restored, order)
        self.assertEqual(restored.status, OrderStatus.PENDING)

    def test_rejects_invalid_status(self) -> None:
        payload = (
            '{"order_id":"o-1","customer_name":"Ada","product":"Keyboard",'
            '"quantity":1,"total_price":99.0,"status":"unknown"}'
        )

        with self.assertRaisesRegex(ValueError, "invalid order status"):
            Order.from_json(payload)

    def test_rejects_missing_fields(self) -> None:
        with self.assertRaisesRegex(ValueError, "missing order fields"):
            Order.from_json('{"order_id":"o-1"}')


if __name__ == "__main__":
    unittest.main()
