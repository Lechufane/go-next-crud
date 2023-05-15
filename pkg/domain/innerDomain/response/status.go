package response

import "net/http"

type StatusBody struct {
	Index int    `json:"index"`
	Name  string `json:"name"`
	Http  int    `json:"http"`
}

type Status int

const (
	// GeneralExceptions
	InternalServerError Status = iota
	NotFound
	BadRequest
	Conflict
	Unknown

	// UserExceptions
	InvalidEmail
	InvalidPassword
	InvalidUsername
	EmailAlreadyExists
	UsernameAlreadyExists
	Unauthorized
	Forbidden

	// DBExepctions
	DBQueryError
	DBExecutionError
	DBRowsError
	DBLastRowIdError
	DBScanError
	DBTransactionError
	DBTransactionClosed
	DBCommitError
	DBItemAlreadyExists

	// JsonExceptions
	JsonDecodingError
	JsonEncodingError

	// SuccessfulCodes
	SuccessfulCreation
	SuccessfulDeletion
	SuccessfulUpdate
	SuccessfulSearch

	// FailureCodes
	FailedCreation
	FailedDeletion
	FailedUpdation
	FailedSearch

	// EncryptionExceptions
	EncryptionError
	DecryptionError

	ViewParsedSuccesfully

	// RequestExceptions
	ErrorInRequest
	SuccesfulRequest

	UserFound

	// ApiExceptions
	ClientsApiError
	MissionsApiError
	ConsumablesApiError

	// CustomerExceptions
	CustomerFound
	CustomerNotFound

	// ProductExceptions
	ProductFound
	ProductNotFound

	// UnitsExceptions
	UnitFound
	UnitNotFound

	// MissionExceptions
	MissionFound

	// DronesExceptions
	DronesFound
	DronesNotFound

	// TerrainExceptions
	TerrainFound

	// CropExceptions
	CropsFound
	CropsNotFound
)

var (
	statusMap = map[Status]StatusBody{
		InternalServerError: {Index: 0, Name: "InternalServerError", Http: http.StatusInternalServerError},
		NotFound:            {Index: 1, Name: "NotFound", Http: http.StatusNotFound},
		BadRequest:          {Index: 2, Name: "BadRequest", Http: http.StatusBadRequest},
		Conflict:            {Index: 3, Name: "Conflict", Http: http.StatusConflict},
		Unknown:             {Index: 4, Name: "Unknown", Http: http.StatusNotImplemented},

		InvalidEmail:          {Index: 5, Name: "InvalidEmail", Http: http.StatusBadRequest},
		InvalidPassword:       {Index: 6, Name: "InvalidPassword", Http: http.StatusBadRequest},
		InvalidUsername:       {Index: 7, Name: "InvalidUsername", Http: http.StatusBadRequest},
		EmailAlreadyExists:    {Index: 8, Name: "EmailAlreadyExists", Http: http.StatusConflict},
		UsernameAlreadyExists: {Index: 9, Name: "UsernameAlreadyExists", Http: http.StatusConflict},
		Unauthorized:          {Index: 10, Name: "Unauthorized", Http: http.StatusUnauthorized},
		Forbidden:             {Index: 11, Name: "Forbidden", Http: http.StatusForbidden},

		DBQueryError:        {Index: 12, Name: "DBQueryError", Http: http.StatusInternalServerError},
		DBExecutionError:    {Index: 13, Name: "DBExecutionError", Http: http.StatusInternalServerError},
		DBRowsError:         {Index: 14, Name: "DBRowsError", Http: http.StatusInternalServerError},
		DBLastRowIdError:    {Index: 15, Name: "DBLastRowIdError", Http: http.StatusInternalServerError},
		DBScanError:         {Index: 16, Name: "DBScanError", Http: http.StatusInternalServerError},
		DBTransactionError:  {Index: 17, Name: "DBTransactionError", Http: http.StatusInternalServerError},
		DBTransactionClosed: {Index: 18, Name: "DBTransactionClosed", Http: http.StatusInternalServerError},
		DBCommitError:       {Index: 19, Name: "DBCommitError", Http: http.StatusInternalServerError},
		DBItemAlreadyExists: {Index: 20, Name: "DBItemAlreadyExists", Http: http.StatusConflict},

		JsonDecodingError: {Index: 21, Name: "JsonDecodingError", Http: http.StatusInternalServerError},
		JsonEncodingError: {Index: 22, Name: "JsonEncodingError", Http: http.StatusInternalServerError},

		SuccessfulCreation: {Index: 23, Name: "SuccessfulCreation", Http: http.StatusCreated},
		SuccessfulDeletion: {Index: 24, Name: "SuccessfulDeletion", Http: http.StatusOK},
		SuccessfulUpdate:   {Index: 25, Name: "SuccessfulUpdate", Http: http.StatusOK},
		SuccessfulSearch:   {Index: 26, Name: "SuccessfulSearch", Http: http.StatusOK},

		FailedCreation: {Index: 27, Name: "FailedCreation", Http: http.StatusConflict},
		FailedDeletion: {Index: 28, Name: "FailedDeletion", Http: http.StatusConflict},
		FailedUpdation: {Index: 29, Name: "FailedUpdation", Http: http.StatusConflict},
		FailedSearch:   {Index: 30, Name: "FailedSearch", Http: http.StatusConflict},

		EncryptionError: {Index: 31, Name: "EncryptionError", Http: http.StatusInternalServerError},
		DecryptionError: {Index: 32, Name: "DecryptionError", Http: http.StatusInternalServerError},

		ViewParsedSuccesfully: {Index: 33, Name: "ViewParsedSuccesfully", Http: http.StatusOK},

		ErrorInRequest:   {Index: 34, Name: "InvalidRequest", Http: http.StatusInternalServerError},
		SuccesfulRequest: {Index: 35, Name: "SuccesfulRequest", Http: http.StatusOK},

		UserFound: {Index: 36, Name: "UserFound", Http: http.StatusOK},

		ClientsApiError:     {Index: 37, Name: "ClientsApiError", Http: http.StatusInternalServerError},
		MissionsApiError:    {Index: 38, Name: "MissionsApiError", Http: http.StatusInternalServerError},
		ConsumablesApiError: {Index: 39, Name: "ConsumablesApiError", Http: http.StatusInternalServerError},

		CustomerFound:    {Index: 40, Name: "CustomerFound", Http: http.StatusOK},
		CustomerNotFound: {Index: 41, Name: "CustomerNotFound", Http: http.StatusNotFound},

		ProductFound:    {Index: 42, Name: "ProductFound", Http: http.StatusOK},
		ProductNotFound: {Index: 43, Name: "ProductNotFound", Http: http.StatusNotFound},

		UnitFound:    {Index: 44, Name: "UnitFound", Http: http.StatusOK},
		UnitNotFound: {Index: 45, Name: "UnitNotFound", Http: http.StatusNotFound},

		MissionFound: {Index: 46, Name: "MissionFound", Http: http.StatusOK},

		DronesFound:    {Index: 47, Name: "DronesFound", Http: http.StatusOK},
		DronesNotFound: {Index: 48, Name: "DronesNotFound", Http: http.StatusNotFound},
		TerrainFound:   {Index: 49, Name: "TerrainFound", Http: http.StatusOK},

		CropsFound:    {Index: 50, Name: "CropsFound", Http: http.StatusOK},
		CropsNotFound: {Index: 51, Name: "CropsNotFound", Http: http.StatusNotFound},
	}
)

func (s Status) String() string {
	return statusMap[s].Name
}

func (s Status) Index() int {
	return statusMap[s].Index
}

func (s Status) StatusCode() int {
	return statusMap[s].Http
}

func GetStatusByName(name string) Status {
	for key, value := range statusMap {
		if value.Name == name {
			return key
		}
	}
	return Unknown
}
