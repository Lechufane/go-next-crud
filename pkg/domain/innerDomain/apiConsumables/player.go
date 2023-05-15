package apiConsumables

type Player struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Level   int     `json:"level"`
	Rank    int     `json:"rank"`
	Winrate float64 `json:"winrate"`
	ClanId  int     `json:"clanId"`
}
