package productsHandler

import (
	"net/http"

	"github.com/Lechufane/go-next-crud/pkg/constants"
	"github.com/Lechufane/go-next-crud/pkg/domain/innerDomain/apiConsumables"
	"github.com/Lechufane/go-next-crud/pkg/domain/innerDomain/response"
	requestHelper "github.com/Lechufane/go-next-crud/pkg/useCases/helpers/request"
	responseHelper "github.com/Lechufane/go-next-crud/pkg/useCases/helpers/response"
)

func GetPlayerById(productId string) (apiConsumables.Player, response.Status) {
	var player apiConsumables.Player
	resp, status := requestHelper.GetRequest(constants.BASE_PATH + constants.PLAYERS_API_URL + "/" + productId, nil)
	if status != response.SuccesfulRequest {
		return player, response.InternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return player, response.ConsumablesApiError
	}
	status = responseHelper.ParseResponseStruct(&player, resp, response.ProductFound)
	return player, status
}
