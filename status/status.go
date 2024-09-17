package status

type Status = uint16

const (
  Success Status = iota // Успешно
  Any  // Любая ошибка
  NotFound  // Не найдено
  PermissionDenied  // Недостаточно прав
  TooFrequentRequests  // Слишком частые запросы
  InternalError  // Внутренняя ошибка сервера
  ManyConnections  // Много подключений
  NotEnoughArguments  // Недостаточно аргументов
  NotAuthorized  // Не авторизован
  Authorized  // Уже авторизован
  IncorrectValue  // Неверное значение
  Inactivity  // Бездействие
  Timeout  // Время ожидания истекло
  InvalidInput  // Неверный ввод
  ResourceUnavailable  // Ресурс недоступен
  OperationFailed  // Операция не удалась
)
var m = map[Status]string{
  Success:"Успешно",
  Any:"Любая ошибка",
  NotFound:"Не найдено",
  PermissionDenied:"Недостаточно прав",
  TooFrequentRequests:"Слишком частые запросы",
  InternalError:"Внутренняя ошибка сервера",
  ManyConnections:"Много подключений",
  NotEnoughArguments:"Недостаточно аргументов",
  NotAuthorized:"Не авторизован",
  Authorized:"Уже авторизован",
  IncorrectValue:"Неверное значение",
  Inactivity:"Бездействие",
  Timeout:"Время ожидания истекло",
  InvalidInput:"Неверный ввод",
  ResourceUnavailable:"Ресурс недоступен",
  OperationFailed:"Операция не удалась",
}

func readable(s Status) string {
  return m[s]
}