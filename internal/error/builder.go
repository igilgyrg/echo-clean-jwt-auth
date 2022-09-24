package error

import "net/http"

// NewUnauthorizedError New Unauthorized Error
func NewUnauthorizedError(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  Unauthorized.Error(),
		ErrCauses: causes,
	}
}

func NewBadRequestError(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  BadRequest.Error(),
		ErrCauses: causes,
	}
}

func NewSystemError(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  InternalSystem.Error(),
		ErrCauses: causes,
	}
}

func NewItemIsExists(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ItemIsExists.Error(),
		ErrCauses: causes,
	}
}

func NewUserIsExistsWithEmail(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  UserIsExistsWithEmail.Error(),
		ErrCauses: causes,
	}
}

func NewRefreshTokenInvalid(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  InvalidRefreshToken.Error(),
		ErrCauses: causes,
	}
}

func NewAccessTokenInvalid(causes interface{}) HttpRestError {
	return &HttpError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  InvalidAccessToken.Error(),
		ErrCauses: causes,
	}
}
