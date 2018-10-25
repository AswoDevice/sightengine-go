package sightengine_go

type Response struct {
	Status  string `json:"status"`
	Request struct {
		ID         string  `json:"id"`
		Timestamp  float64 `json:"timestamp"`
		Operations int     `json:"operations"`
	} `json:"request"`
	Weapon  float64 `json:"weapon"`
	Alcohol float64 `json:"alcohol"`
	Drugs   float64 `json:"drugs"`
	Scam    struct {
		Prob float64 `json:"prob"`
	} `json:"scam"`
	Nudity struct {
		Raw     float64 `json:"raw"`
		Safe    float64 `json:"safe"`
		Partial float64 `json:"partial"`
	} `json:"nudity"`
	Type struct {
		Photo        float64 `json:"photo"`
		Illustration float64 `json:"illustration"`
	} `json:"type"`
	Face struct {
		Single   float64 `json:"single"`
		Multiple float64 `json:"multiple"`
	} `json:"face"`
	Faces []struct {
		X1       float64 `json:"x1"`
		Y1       float64 `json:"y1"`
		X2       float64 `json:"x2"`
		Y2       float64 `json:"y2"`
		Features struct {
			LeftEye struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"left_eye"`
			RightEye struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"right_eye"`
			NoseTip struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"nose_tip"`
			LeftMouthCorner struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"left_mouth_corner"`
			RightMouthCorner struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"right_mouth_corner"`
		} `json:"features"`
		Attributes struct {
			Female     float64 `json:"female"`
			Male       float64 `json:"male"`
			Minor      float64 `json:"minor"`
			Sunglasses float64 `json:"sunglasses"`
		} `json:"attributes"`
		Celebrity []struct {
			Name string  `json:"name"`
			Prob float64 `json:"prob"`
		} `json:"celebrity"`
	} `json:"faces"`
	Sharpness  float64 `json:"sharpness"`
	Brightness float64 `json:"brightness"`
	Contrast   float64 `json:"contrast"`
	Colors     struct {
		Dominant struct {
			R   int    `json:"r"`
			G   int    `json:"g"`
			B   int    `json:"b"`
			Hex string `json:"hex"`
		} `json:"dominant"`
		Other []struct {
			R   int    `json:"r"`
			G   int    `json:"g"`
			B   int    `json:"b"`
			Hex string `json:"hex"`
		} `json:"other"`
	} `json:"colors"`
	Text struct {
		HasArtificial float64 `json:"has_artificial"`
		HasNatural    float64 `json:"has_natural"`
		Boxes         []struct {
			X1    float64 `json:"x1"`
			Y1    float64 `json:"y1"`
			X2    float64 `json:"x2"`
			Y2    float64 `json:"y2"`
			Label string  `json:"label"`
			Prob  float64 `json:"prob"`
		} `json:"boxes"`
	} `json:"text"`
	Offensive struct {
		Prob  float64 `json:"prob"`
		Boxes []struct {
			X1    float64 `json:"x1"`
			Y1    float64 `json:"y1"`
			X2    float64 `json:"x2"`
			Y2    float64 `json:"y2"`
			Label string  `json:"label"`
			Prob  float64 `json:"prob"`
		} `json:"boxes"`
	} `json:"offensive"`
	Media struct {
		ID  string `json:"id"`
		URI string `json:"uri"`
	} `json:"media"`
}