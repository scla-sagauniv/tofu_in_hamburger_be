package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type RecipeOnDb struct {
	Id                 int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title              string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	RecipeUrl          string `protobuf:"bytes,3,opt,name=recipe_url,json=recipeUrl,proto3" json:"recipe_url,omitempty"`
	ImageUrl           string `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Pickup             bool   `protobuf:"varint,5,opt,name=pickup,proto3" json:"pickup,omitempty"`
	Nickname           string `protobuf:"bytes,6,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Materials          string `protobuf:"bytes,7,opt,name=materials,proto3" json:"materials,omitempty"`
	MaterialIds        string `protobuf:"bytes,8,opt,name=material_ids,json=materialIds,proto3" json:"material_ids,omitempty"`
	Publishday         string `protobuf:"bytes,9,opt,name=publishday,proto3" json:"publishday,omitempty"`
	Ranking            int64  `protobuf:"varint,10,opt,name=ranking,proto3" json:"ranking,omitempty"`
	RecipeIndicationId int64  `protobuf:"varint,11,opt,name=recipe_indication_id,json=recipeIndicationId,proto3" json:"recipe_indication_id,omitempty"`
	RecipeCostId       int64  `protobuf:"varint,12,opt,name=recipe_cost_id,json=recipeCostId,proto3" json:"recipe_cost_id,omitempty"`
	CreatedAt          string `protobuf:"bytes,13,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

type grpcReq struct {
	Recipes []*RecipeOnDb `json:"recipes"`
}

type apiResp struct {
	Result []RecipeOnApi `json:"result"`
}

type RecipeOnApi struct {
	FoodImageURL      string   `json:"foodImageUrl"`
	MediumImageURL    string   `json:"mediumImageUrl"`
	Nickname          string   `json:"nickname"`
	Pickup            int      `json:"pickup"`
	Rank              string   `json:"rank"`
	RecipeCost        string   `json:"recipeCost"`
	RecipeDescription string   `json:"recipeDescription"`
	RecipeID          int      `json:"recipeId"`
	RecipeIndication  string   `json:"recipeIndication"`
	RecipeMaterial    []string `json:"recipeMaterial"`
	RecipePublishday  string   `json:"recipePublishday"`
	RecipeTitle       string   `json:"recipeTitle"`
	RecipeURL         string   `json:"recipeUrl"`
	Shop              int      `json:"shop"`
	SmallImageURL     string   `json:"smallImageUrl"`
}

func (a RecipeOnApi) convertToRecipeIndicationId() int64 {
	switch a.RecipeIndication {
	case "指定なし":
		return 0
	case "5分以内":
		return 1
	case "約10分":
		return 2
	case "約15分":
		return 3
	case "約30分":
		return 4
	case "約1時間":
		return 5
	case "1時間以上":
		return 6
	}
	return 0
}

func (a RecipeOnApi) convertToRecipeCostId() int64 {
	switch a.RecipeCost {
	case "指定なし":
		return 0
	case "100円以下":
		return 1
	case "300円前後":
		return 2
	case "500円前後":
		return 3
	case "1,000円前後":
		return 4
	case "2,000円前後":
		return 5
	case "3,000円前後":
		return 6
	case "5,000円前後":
		return 7
	case "10,000円以上":
		return 8
	}
	return 0
}

func (a RecipeOnApi) convertToPickupBoolean() bool {
	if a.Pickup == 1 {
		return true
	} else {
		return false
	}
}

func (a RecipeOnApi) convertToRecipePublishday() string {
	layout := "2006/01/02 15:04:05"
	t, err := time.Parse(layout, a.RecipePublishday)
	if err != nil {
		panic(err)
	}
	pb := timestamppb.New(t)

	return pb.AsTime().Format(time.RFC3339)
}

func (a RecipeOnApi) convertToRanking() int64 {
	num, err := strconv.Atoi(a.Rank)
	if err != nil {
		panic(err)
	}
	return int64(num)
}

func main() {
	var recipes []*RecipeOnDb
	categoryList := []string{"30", "31", "32", "33", "14", "15", "16", "17", "23", "18", "22", "21", "10", "11", "12", "34", "19", "27", "35", "13", "20", "36", "37", "38", "39", "40", "26", "41", "42", "43", "44", "25", "46", "47", "48", "24", "49", "50", "51", "52", "53", "54", "55"}
	// categoryList := []int{275, 276, 277, 278, 68, 66, 67, 69, 70, 71, 72, 73, 74, 75, 76, 77, 443, 78, 80, 81, 79, 83, 444, 82, 445, 446, 447, 448, 449, 450, 97, 452, 98, 453, 454, 99, 456, 457, 455, 451, 96, 458, 95, 100, 101, 102, 103, 105, 107, 104, 478, 706, 479, 480, 481, 108, 109, 482, 483, 111, 112, 113, 114, 484, 115, 121, 131, 126, 124, 122, 123, 125, 127, 368, 128, 129, 130, 132, 133, 134, 135, 271, 687, 137, 676, 681, 369, 677, 683, 682, 678, 679, 684, 680, 138, 139, 140, 141, 142, 685, 686, 143, 145, 146, 144, 147, 151, 382, 152, 153, 154, 155, 156, 383, 384, 272, 385, 386, 158, 159, 161, 387, 160, 388, 169, 389, 171, 168, 167, 170, 164, 165, 166, 173, 390, 162, 415, 424, 421, 189, 187, 417, 416, 418, 722, 419, 420, 423, 190, 703, 184, 188, 185, 186, 191, 192, 193, 194, 195, 196, 675, 274, 463, 464, 700, 710, 711, 273, 485, 197, 486, 487, 488, 198, 199, 200, 201, 202, 203, 258, 204, 440, 205, 438, 439, 206, 215, 207, 208, 209, 210, 211, 216, 212, 441, 442, 214, 217, 218, 432, 433, 434, 435, 436, 229, 221, 220, 222, 219, 223, 227, 231, 437, 230, 392, 394, 391, 399, 395, 401, 404, 397, 393, 403, 400, 396, 405, 407, 412, 406, 398, 413, 411, 409, 410, 402, 698, 723, 408, 234, 631, 632, 633, 634, 635, 238, 244, 256, 701, 248, 255, 257, 262, 260, 261, 265, 266, 267, 268, 465, 269, 300, 301, 302, 307, 303, 304, 305, 306, 309, 310, 308, 311, 312, 313, 314, 315, 316, 317, 717, 318, 319, 320, 321, 323, 324, 325, 326, 327, 328, 329, 330, 331, 332, 333, 334, 322, 335, 718, 719, 720, 336, 337, 338, 339, 340, 341, 342, 343, 344, 345, 346, 347, 348, 349, 350, 351, 352, 353, 354, 355, 356, 357, 358, 359, 360, 361, 362, 363, 364, 365, 366, 367, 721, 688, 459, 460, 461, 697, 462, 690, 691, 702, 692, 693, 689, 695, 696, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 489, 490, 491, 492, 493, 494, 495, 496, 497, 708, 498, 499, 500, 501, 502, 503, 504, 505, 705, 699, 506, 507, 508, 509, 510, 511, 709, 724, 512, 513, 514, 707, 515, 516, 704, 517, 518, 519, 520, 521, 522, 523, 524, 525, 526, 712, 531, 532, 533, 534, 535, 536, 537, 539, 542, 713, 543, 538, 541, 546, 547, 548, 540, 544, 545, 549, 550, 551, 552, 553, 554, 555, 565, 556, 557, 558, 559, 560, 561, 714, 562, 563, 564, 566, 567, 568, 569, 570, 578, 571, 577, 572, 573, 574, 575, 576, 579, 580, 581, 582, 583, 584, 585, 586, 587, 588, 589, 590, 591, 596, 597, 598, 599, 602, 600, 601, 603, 604, 605, 606, 607, 608, 609, 610, 612, 613, 611, 614, 615, 616, 617, 618, 619, 620, 621, 622, 623, 624, 625, 626, 627, 628, 629, 715, 716, 630, 636, 637, 638, 639, 640, 641, 642, 643, 644, 645, 646, 648, 649, 650, 651, 652, 653, 654, 655, 656, 657, 658, 659, 660, 661, 662, 663, 664, 665, 666, 667, 668, 669, 670, 671, 672, 673, 674}
	now := time.Now().Format(time.RFC3339)

	for i, category := range categoryList {
		if i >= 50 {
			break
		}
		fmt.Printf("%d / %d\n", i, len(categoryList))
		// category := categoryList[0]
		// APIのURLを指定
		url := fmt.Sprintf("https://app.rakuten.co.jp/services/api/Recipe/CategoryRanking/20170426?format=json&categoryId=%s&applicationId=1009993429057254143", category)

		// GETリクエストを作成
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("リクエスト作成エラー:", err)
			return
		}

		// リクエストを送信
		time.Sleep(1000 * time.Millisecond)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("リクエスト送信エラー:", err)
			return
		}
		defer resp.Body.Close()

		// レスポンスのボディを読み取り
		body, err := ioutil.ReadAll(resp.Body)
		var response apiResp
		err = json.Unmarshal([]byte(body), &response)
		if err != nil {
			fmt.Println("JSONデコードエラー:", err)
			return
		}

		fmt.Printf("result len: %d\n", len(response.Result))
		if len(response.Result) == 0 {
			fmt.Println(string(body))
			fmt.Println(url)
		}
		for _, res := range response.Result {
			recipe := RecipeOnDb{
				Id:                 int64(res.RecipeID),
				Title:              res.RecipeTitle,
				RecipeUrl:          res.RecipeURL,
				ImageUrl:           res.FoodImageURL,
				Pickup:             res.convertToPickupBoolean(),
				Nickname:           res.Nickname,
				Materials:          strings.Join(res.RecipeMaterial, ","),
				MaterialIds:        strings.Join(res.RecipeMaterial, ","),
				Publishday:         res.convertToRecipePublishday(),
				Ranking:            res.convertToRanking(),
				RecipeIndicationId: res.convertToRecipeIndicationId(),
				RecipeCostId:       res.convertToRecipeCostId(),
				CreatedAt:          now,
				UpdatedAt:          now,
			}
			recipes = append(recipes, &recipe)
		}
	}

	grpcReq := grpcReq{
		Recipes: recipes,
	}
	jsonData, err := json.Marshal(&grpcReq)
	if err != nil {
		fmt.Println(err)
		return
	}

	// レスポンスを表示
	fmt.Println(string(jsonData))

	resp, err := http.Post("http://localhost:8080/rpc.ingredientRain.v1.RecipeService/CreateRecipesByBatch", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusCreated {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		jsonStr := string(body)
		fmt.Println("Response: ", jsonStr)
	} else {
		fmt.Println("Get failed with error: ", resp.Status)
	}
}
