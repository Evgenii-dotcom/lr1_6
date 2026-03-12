import sys
import os

# Добавляем правильный путь для Windows debug сборки
sys.path.append(os.path.join("target", "debug"))

# Импортируем модуль (без lib и без расширения)
import fastcalc

counter = fastcalc.Counter(10)

counter.increment()
counter.increment()

print(counter.get())