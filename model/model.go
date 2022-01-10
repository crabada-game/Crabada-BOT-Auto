package model

type TeamList struct {
	ErrorCode interface{} `json:"error_code"`
	Message   interface{} `json:"message"`
	Result    struct {
		TotalRecord int `json:"totalRecord"`
		TotalPages  int `json:"totalPages"`
		Page        int `json:"page"`
		Limit       int `json:"limit"`
		Data        []struct {
			TeamID               int64  `json:"team_id"`
			Owner                string `json:"owner"`
			CrabadaID1           int    `json:"crabada_id_1"`
			Crabada1Photo        string `json:"crabada_1_photo"`
			Crabada1Hp           int    `json:"crabada_1_hp"`
			Crabada1Speed        int    `json:"crabada_1_speed"`
			Crabada1Armor        int    `json:"crabada_1_armor"`
			Crabada1Damage       int    `json:"crabada_1_damage"`
			Crabada1Critical     int    `json:"crabada_1_critical"`
			Crabada1IsOrigin     int    `json:"crabada_1_is_origin"`
			Crabada1IsGenesis    int    `json:"crabada_1_is_genesis"`
			Crabada1LegendNumber int    `json:"crabada_1_legend_number"`
			CrabadaID2           int    `json:"crabada_id_2"`
			Crabada2Photo        string `json:"crabada_2_photo"`
			Crabada2Hp           int    `json:"crabada_2_hp"`
			Crabada2Speed        int    `json:"crabada_2_speed"`
			Crabada2Armor        int    `json:"crabada_2_armor"`
			Crabada2Damage       int    `json:"crabada_2_damage"`
			Crabada2Critical     int    `json:"crabada_2_critical"`
			Crabada2IsOrigin     int    `json:"crabada_2_is_origin"`
			Crabada2IsGenesis    int    `json:"crabada_2_is_genesis"`
			Crabada2LegendNumber int    `json:"crabada_2_legend_number"`
			CrabadaID3           int    `json:"crabada_id_3"`
			Crabada3Photo        string `json:"crabada_3_photo"`
			Crabada3Hp           int    `json:"crabada_3_hp"`
			Crabada3Speed        int    `json:"crabada_3_speed"`
			Crabada3Armor        int    `json:"crabada_3_armor"`
			Crabada3Damage       int    `json:"crabada_3_damage"`
			Crabada3Critical     int    `json:"crabada_3_critical"`
			Crabada3IsOrigin     int    `json:"crabada_3_is_origin"`
			Crabada3IsGenesis    int    `json:"crabada_3_is_genesis"`
			Crabada3LegendNumber int    `json:"crabada_3_legend_number"`
			BattlePoint          int    `json:"battle_point"`
			TimePoint            int    `json:"time_point"`
			MinePoint            int    `json:"mine_point"`
			GameType             string `json:"game_type"`
			MineStartTime        int    `json:"mine_start_time"`
			MineEndTime          int64  `json:"mine_end_time"`
			GameID               int64  `json:"game_id"`
			GameStartTime        int    `json:"game_start_time"`
			GameEndTime          int    `json:"game_end_time"`
			ProcessStatus        string `json:"process_status"`
			GameRound            int    `json:"game_round"`
			Status               string `json:"status"`
			Crabada1Class        int    `json:"crabada_1_class"`
			Crabada2Class        int    `json:"crabada_2_class"`
			Crabada3Class        int    `json:"crabada_3_class"`
			Crabada1Type         int    `json:"crabada_1_type"`
			Crabada2Type         int    `json:"crabada_2_type"`
			Crabada3Type         int    `json:"crabada_3_type"`
		} `json:"data"`
		TeamSize int `json:"team_size"`
	} `json:"result"`
}
