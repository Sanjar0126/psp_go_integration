package _const

// ErrorCode represents the platform-wide error codes that can be raised by PSP
type ErrorCode int64

const (
	// Success Успешный запрос
	Success ErrorCode = 0

	// SignCheckFailed Ошибка проверки подписи
	SignCheckFailed ErrorCode = -1

	// IncorrectParameterAmount Неверная сумма оплаты
	IncorrectParameterAmount ErrorCode = -2

	// ActionNotFound Запрашиваемое действие не найдено
	ActionNotFound ErrorCode = -3

	// AlreadyPaid Транзакция ранее была подтверждена
	// (при попытке подтвердить или отменить ранее подтвержденную транзакцию)
	AlreadyPaid ErrorCode = -4

	// UserNotExist Не найден пользователь/заказ (проверка параметра merchant_trans_id)
	UserNotExist ErrorCode = -5

	// TransactionNotExist Не найдена транзакция (проверка параметра merchant_prepare_id)
	TransactionNotExist ErrorCode = -6

	// FailedUpdateUser Ошибка при изменении данных пользователя (изменение баланса счета и т.п.)
	FailedUpdateUser ErrorCode = -7

	// ErrorInRequestFromPSP Ошибка в запросе от PSP (переданы не все параметры и т.п.)
	ErrorInRequestFromPSP ErrorCode = -8

	// TransactionCanceled Транзакция ранее была отменена
	// (При попытке подтвердить или отменить ранее отмененную транзакцию)
	TransactionCanceled ErrorCode = -9

	// VendorNotFound Поставщик не найден в системе
	VendorNotFound ErrorCode = -10

	// TransactionTypeNotCorrect Не правильный тип транзакции
	TransactionTypeNotCorrect ErrorCode = -11
)

var ErrorText = map[ErrorCode]string{
	Success:                   "Success",
	SignCheckFailed:           "SIGN CHECK FAILED!",
	IncorrectParameterAmount:  "Incorrect parameter amount",
	ActionNotFound:            "Action not found",
	AlreadyPaid:               "Already paid",
	UserNotExist:              "User does not exist",
	TransactionNotExist:       "Transaction does not exist",
	FailedUpdateUser:          "Failed to update user",
	ErrorInRequestFromPSP:     "Error in request from PSP",
	TransactionCanceled:       "Transaction cancelled",
	VendorNotFound:            "The vendor is not found",
	TransactionTypeNotCorrect: "Transaction type is not correct",
}

func GetErrorString(code ErrorCode) string {
	return ErrorText[code]
}
