package main

import (
	"context"
	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	url := "mongodb://localhost:27017"
	db := "solarland_test"

	defaultOptions := []*options.ClientOptions{
		options.Client().ApplyURI(url).SetConnectTimeout(time.Second * 10),
	}
	client, err := mongo.Connect(context.Background(), defaultOptions...)
	if err != nil {
		panic(err)
	}

	filter := bson.M{"user.userID": "test_user_id_111"}

	account := &Account{
		User: &User{
			Name:   "t8",
			UserID: "test_user_id_111",
		},
		IsOfficialAccount: false,
	}

	ctx := context.Background()
	client.Database(db).Drop(ctx)
	col := client.Database(db).Collection("Account2")
	_, err = col.InsertOne(ctx, account)
	check(err)

	_, err = col.UpdateOne(ctx, filter, bson.M{"$set": bson.M{"isOfficialAccount": true}}, options.Update().SetUpsert(true))
	check(err)



	_, err = col.UpdateOne(ctx, filter, bson.M{"$set": &Account{IsOfficialAccount: false}}, options.Update().SetHint())
	check(err)
	//col.UpdateOne(ctx, filter, bson.M{"$set": Updater{Account{IsOfficialAccount: false}}})
}

type Account struct {
	Phone             string    `json:"phone,omitempty" bson:"phone,omitempty"`
	Email             string    `json:"email,omitempty" bson:"email,omitempty"`
	DeviceID          string    `json:"deviceID,omitempty" bson:"deviceID,omitempty"`
	Password          []byte    `json:"password,omitempty" bson:"password,omitempty"`
	GoogleID          string    `json:"googleID,omitempty" bson:"googleID,omitempty"`
	CreatorID         string    `json:"creatorID,omitempty" bson:"creatorID,omitempty"`
	CreateTime        time.Time `json:"createTime,omitempty" bson:"createTime,omitempty"`
	User              *User     `json:"user,omitempty" bson:"user,omitempty"`
	AuthorityRoles    []string  `json:"authorityRoles,omitempty" bson:"authorityRoles,omitempty"`
	IsOfficialAccount bool      `json:"isOfficialAccount,omitempty" bson:"isOfficialAccount,omitempty"`
}

func (a *Account) MarshalBSON() ([]byte, error) {
	flat, err := flatbson.Flatten(a)
	if err != nil {
		return nil, err
	}

	return bson.Marshal(flat)
}

type User struct {
	UserID            string    `bson:"userID,omitempty" json:"userID,omitempty"`
	NumberID          int64     `bson:"numberID,omitempty" json:"numberID,omitempty"`
	Name              string    `bson:"name,omitempty" json:"name,omitempty"`
	NameSeq           int       `json:"nameSeq,omitempty" bson:"nameSeq,omitempty"`
	Gender            int8      `bson:"gender,omitempty" json:"gender,omitempty"` // Community Gender, different from avatarBox gender
	Birthday          string    `bson:"birthday,omitempty" json:"birthday,omitempty"`
	IconStoreID       string    `json:"iconStoreID,omitempty" bson:"iconStoreID,omitempty"`
	BodyStoreID       string    `json:"bodyStoreID,omitempty" bson:"bodyStoreID,omitempty"`
	Profile           string    `json:"profile,omitempty" bson:"profile,omitempty"`
	AvatarJSON        string    `json:"avatarJSON,omitempty" bson:"avatarJSON,omitempty"`
	FollowingCount    int32     `json:"followingCount,omitempty" bson:"followingCount,omitempty"`
	FollowerCount     int32     `json:"followerCount,omitempty" bson:"followerCount,omitempty"`
	IsOnline          bool      `json:"user.isOnline" bson:"isOnline"`
	ActiveTime        time.Time `json:"activeTime" bson:"activeTime,omitempty"`
	LoginTime         time.Time `json:"loginTime" bson:"loginTime,omitempty"`
	QuestionnaireDone bool      `json:"questionnaireDone" bson:"questionnaireDone,omitempty"`
	TotalGameSeconds  int64     `json:"totalGameSeconds" bson:"totalGameSeconds,omitempty"`
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
