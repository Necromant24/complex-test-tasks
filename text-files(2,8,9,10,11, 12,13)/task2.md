- Для оптимизации производительности уже накинуты индексы для таблиц
- использование пула инстансов подключений к БД т.к. горм поддерживает пул запросов к своему db экземпляру ( и по факту можно использовать 1 его инстанс при небольшой нагрузке) [сейчас при каждом запросе создаю новое подключение]
- горизонтальное масштабирование с использованием nginx и инстансами нескольких подов в docker

