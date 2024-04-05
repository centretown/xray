package tools

// func Construct(rec *model.Record) (typ model.Recorder, err error) {
// 	var item any
// 	cat := category.Category(rec.Category)
// 	switch cat {
// 	case category.Texture:
// 		{
// 			txt := &Texture{}
// 			item = &txt.TextureItem
// 			err := json.Unmarshal([]byte(rec.Content), item)
// 			return txt, err
// 		}

// 	case category.Circle:
// 		{
// 			txt := &Circle{}
// 			item = &txt.CircleItem
// 			err := json.Unmarshal([]byte(rec.Content), item)
// 			return txt, err
// 		}

// 	case category.Motor:
// 		{
// 			txt := &Mover{}
// 			item = &txt.MoverItem
// 			err := json.Unmarshal([]byte(rec.Content), item)
// 			return txt, err
// 		}

// 	case category.Game:
// 		{
// 			txt := &Game{}
// 			item = &txt.GameItem
// 			err := json.Unmarshal([]byte(rec.Content), item)
// 			return txt, err
// 		}
// 	default:
// 		panic("attempting to construct undefined type")
// 	}

// 	err := json.Unmarshal([]byte(rec.Content), typ)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return
// }
