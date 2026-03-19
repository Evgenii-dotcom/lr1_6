import sys
import os
import unittest

sys.path.append(os.path.join("target", "debug"))

import fastcalc


class TestCounter(unittest.TestCase):

    def test_initial_value(self):
        counter = fastcalc.Counter(10)
        self.assertEqual(counter.get(), 10)

    def test_increment(self):
        counter = fastcalc.Counter(0)
        counter.increment()
        self.assertEqual(counter.get(), 1)

    def test_multiple_increment(self):
        counter = fastcalc.Counter(5)

        for _ in range(3):
            counter.increment()

        self.assertEqual(counter.get(), 8)

    def test_negative_initial_value(self):
        counter = fastcalc.Counter(-5)
        counter.increment()
        self.assertEqual(counter.get(), -4)


if __name__ == "__main__":
    unittest.main()