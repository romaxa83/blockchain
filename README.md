### Реализация blockchain на golang

Используется БД <a href="https://github.com/dgraph-io/badger" >BadgerDB</a>

используемые ключи в бд
* lh - lastHash


##### Основные понятия

<dl>
  <dt>Блокчейн</dt>
  <dd>
    <p>
По сути, блокчейн это база данных определенной структуры: упорядоченный связанный список. 
Это означает, что блоки хранятся в порядке вставки, при чем каждый блок связан с предыдущем.
Такая структура позволяет получить последний блок в цепочке и эффективно получить
блок по его хэшу.
    </p>
    <p>
В данном блокчейне добавление новых блоков проходит очень легко и быстро, в реальном же
блокчейне добавление новых блоков требует определенной работы: выполнение сложных вычислений,
перед получением права на добавления блока (этот механизм называется <u>Prof-of-Work</u>).
Кроме того, блокчейн — распределенная база, которая не имеет единого центра принятия решения.
Таким образом, новый блок должен быть подтвержден и одобрен другими участниками сети
(данный механизм называется <u>консенсусом</u>)
    </p>
  </dd>
  <dt>Блок</dt>
  <dd>
    <p>
В блокчейне, блоки хранят полезную информацию. Например, в bitcoin блоки хранят транзакции,
суть любой криптовалюты. Помимо полезной информации, в блоке содержится служебная информация:
версия, дата создания в виде timestamp и хеш предыдущего блока.
    </p>
  </dd>
  <dt>Proof-of-Work(PoW)</dt>
  <dd>
    <p>
    При добавления нового блока в блокчейн необходимо проделать некоторую сложную работу.
Именно эта сложная работа делает блокчейн надежным и целостным. В блокчейне некоторые
участники (майнеры) сети работают над поддержанием сети, добавлением в блокчейн новых
блоков и получают вознаграждение за свою работу. В результате их работы блок встраивается
в блокчейн надежным способом, что обеспечивает стабильность всей базы данных блокчейна.
Стоит отметить, что тот, кто выполнил работу, должен также доказать её выполнение.Этот
весь «сделай сложную работу и докажи её»-механизм называется <u>Proof-of-Work (доказательство работы)</u>.
    </p>
    <p>
К примеру в Биткоине цель такой работы — это нахождение хеша блока, который удовлетворяет
определенным требованиям. Данный хеш и служит доказательством. Таким образом, поиск 
доказательства и есть фактическая работа.
    </p>
    <p>
Важно Proof-of-Work алгоритмы должны соответствовать следующему требованию: выполнение
работы должно быть сложным, но проверка доказательства должна быть простой.
Проверка доказательства обычно передается кому-то стороннему, поэтому у них данная 
проверка не должна занимать много времени.
    </p>
  </dd>
</dl>