# Заголовок первого уровня

Документ для проверки рендеринга всех основных блоков Markdown (CommonMark + GFM).

## Заголовок второго уровня

### Заголовок третьего уровня

#### Заголовок четвёртого уровня

##### Заголовок пятого уровня

###### Заголовок шестого уровня

---

## Текст и inline-разметка

Обычный абзац. Mixed text: English and русский в одной строке. Второе предложение, чтобы абзац переносился на несколько строк и было видно межстрочный интервал и ширину колонки текста.

Выделения: **жирный**, *курсив*, ***жирный курсив***, ~~зачёркнутый~~, `inline code`, **жирный с `кодом` внутри**, *курсив со [ссылкой](https://example.com)*.

Ссылки: [обычная](https://example.com), [с title](https://example.com "Подсказка"), [reference-ссылка][ref], автоссылка <https://go.dev>, голый URL https://github.com/yuin/goldmark и email <user@example.com>.

[ref]: https://commonmark.org "CommonMark"

Жёсткий перенос строки через два пробела в конце  
вот эта строка начинается с новой строки.  
И ещё одна через обратный слеш\
в конце строки.

Экранирование: \*не курсив\*, \`не код\`, \# не заголовок, 2 \* 3 = 6.

Длинное слово без пробелов: Donaudampfschifffahrtselektrizitätenhauptbetriebswerkbauunterbeamtengesellschaft_and_some_more_characters_to_force_wrapping.

Длинный URL: https://example.com/api/v3/payments/refunds/very/long/path/that/should/wrap/properly/in/the/pdf/output?param1=value1&param2=value2

Типографика: «ёлочки», „лапки“, тире — и дефис -, многоточие…, спецсимволы © ® ™ € ± ≠ ≤ ≥ → ∞.

## Цитаты

> Простая цитата в один абзац.

> Цитата из нескольких абзацев.
>
> Второй абзац с **выделением** и `кодом`.
>
> > Вложенная цитата второго уровня.
> >
> > - список внутри цитаты
> > - второй пункт

## Списки

Маркированный:

- Первый пункт
- Второй пункт с длинным текстом, который должен переноситься на следующую строку и сохранять отступ относительно маркера списка
  - Вложенный пункт
  - Ещё один
    - Третий уровень
    - И ещё
- Третий пункт

Нумерованный:

1. Первый
2. Второй
   1. Вложенный нумерованный
   2. Ещё один
3. Третий

Нумерация с произвольного числа:

7. Седьмой
8. Восьмой
9. Девятый
10. Десятый

Смешанный с абзацами и кодом внутри:

1. Пункт с абзацем.

   Второй абзац внутри того же пункта.

2. Пункт с блоком кода:

   ```bash
   go build -o downprint .
   ```

3. Пункт с цитатой:

   > Цитата внутри списка.

Список задач:

- [x] Парсинг Markdown
- [x] Рендеринг в HTML
- [ ] Оглавление
  - [x] Вложенная выполненная задача
  - [ ] Вложенная невыполненная

## Код

Fenced-блок без языка:

```
plain text block
    с отступами   и   пробелами
```

Go:

```go
package main

import "fmt"

// Greeter says hello.
type Greeter struct {
	Name string
}

func (g Greeter) Greet() string {
	return fmt.Sprintf("Привет, %s!", g.Name)
}

func main() {
	fmt.Println(Greeter{Name: "мир"}.Greet())
}
```

JSON:

```json
{
  "id": 42,
  "name": "refund",
  "active": true,
  "amount": 1.45,
  "tags": ["visa", "acs"],
  "meta": null
}
```

Bash:

```bash
#!/usr/bin/env bash
set -euo pipefail

for f in *.md; do
  downprint "$f" && echo "ok: $f"
done
```

SQL:

```sql
SELECT id, amount, currency
FROM payments
WHERE status = 'refunded'
  AND created_at >= NOW() - INTERVAL '7 days'
ORDER BY created_at DESC;
```

YAML:

```yaml
server:
  host: 0.0.0.0
  port: 8080
features:
  - pdf
  - html
```

Diff:

```diff
- old line
+ new line
  unchanged line
```

Очень длинная строка в коде (должна переноситься):

```
curl -X POST "https://cert.api.visa.com/acs/v3/payments/refunds" -H "Content-Type: application/json" -H "Authorization: Basic abc...xyz==" -d '{"amount":"1.45","currency":"840"}'
```

Код с отступом (indented code block):

    func indented() {
        return
    }

## Таблицы

Выравнивание колонок:

| По левому краю | По центру | По правому краю |
|:---------------|:---------:|----------------:|
| Amt            | string    | 1.45            |
| Ccy            | string    | 840             |
| XchgRate       | number    | 0.8758621       |

Inline-разметка в ячейках:

| Поле | Описание | Пример |
|------|----------|--------|
| `PAN` | **Номер карты**, маскируется | `4761********0039` |
| `ShrtNm` | *Короткое имя* мерчанта | [ссылка](https://example.com) |
| `Tp` | Тип с экранированным \| пайпом | ~~old~~ new |

Широкая таблица с длинным текстом:

| Параметр | Обязательный | Тип | Описание |
|----------|:------------:|-----|----------|
| correlatnId | да | string | Идентификатор корреляции, который передаётся клиентом и возвращается в ответе без изменений для сопоставления запросов и ответов |
| clientId | да | string | Идентификатор клиента, выданный при регистрации в системе |
| LclDtTm | нет | datetime | Локальное время транзакции в формате ISO 8601 без указания часового пояса |

## Изображения

Изображение по относительному пути:

![Схема конвертации](image.svg)

Изображение со ссылкой:

[![Схема](image.svg "Кликабельная картинка")](https://example.com)

## Горизонтальные линии

Три варианта синтаксиса:

***

___

- - -

## Raw HTML

<details>
<summary>Раскрывающийся блок</summary>

Содержимое внутри `<details>`. В PDF должно быть видно хотя бы summary.

</details>

Клавиши: <kbd>Ctrl</kbd> + <kbd>C</kbd>. Индексы: H<sub>2</sub>O, x<sup>2</sup>. <mark>Выделенный маркером текст</mark>.

<p align="center">Абзац с выравниванием по центру через HTML</p>

<!-- Этот комментарий не должен попасть в PDF -->

## Разрыв страницы

Длинный блок ниже проверяет перенос кода между страницами:

```go
func longFunction() {
	step01()
	step02()
	step03()
	step04()
	step05()
	step06()
	step07()
	step08()
	step09()
	step10()
	step11()
	step12()
	step13()
	step14()
	step15()
	step16()
	step17()
	step18()
	step19()
	step20()
	step21()
	step22()
	step23()
	step24()
	step25()
	step26()
	step27()
	step28()
	step29()
	step30()
}
```

Конец документа.
