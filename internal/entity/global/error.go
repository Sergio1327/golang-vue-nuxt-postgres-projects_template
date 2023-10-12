package global

import "errors"

var (
	// ErrDBUnvailable база данных недоступна
	ErrDBUnvailable = errors.New("база данных недоступна")

	// ErrNeedAuth необходимо предварительно авторизоваться
	ErrNeedAuth = errors.New("необходима авторизация")

	// ErrParamsIncorrect неверные параметры запроса
	ErrParamsIncorrect = errors.New("неверные параметры запроса")

	// ErrTooManyRequest слишком частые неверные запросы, подождите немного
	ErrTooManyRequest = errors.New("слишком частые неверные запросы, подождите немного")

	// ErrInternalError внутряя ошибка
	ErrInternalError = errors.New("произошла внутренняя ошибка, пожалуйста попробуйте выполнить действие позже")

	// ErrNoData данные не найдены"
	ErrNoData = errors.New("данные не найдены")

	// ErrAccessRight ошибка при проверке прав доступа
	ErrAccessRight = errors.New("у вас нет нужных прав доступа")

	// ErrUniqueConstraintViolated нарушено ограничение уникальности
	ErrUniqueConstraintViolated = errors.New("нарушено ограничение уникальности")
)
