package migrations

import (
	"context"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateMigrations() []migrate.Migration {
	return []migrate.Migration{
		migration1,
		migration2,
		migration3,
		migration4,
		migration5,
		migration6,
		migration7,
		migration8,
		migration9,
		migration10,
		migration11,
		migration12,
		migration13,
		migration14,
		migration15,
		migration16,
		migration17,
		migration18,
		migration19,
		migration20,
		migration21,
		migration22,
		migration23,
		migration24,
		migration25,
		migration26,
		migration27,
		migration28,
		migration29,
		migration30,
		migration31,
		migration32,
		migration33,
		migration34,
		migration35,
		migration36,
		migration37,
		migration38,
		migration39,
		migration40,
		migration41,
		migration42,
		migration43,
		migration44,
		migration45,
		migration46,
		migration47,
		migration48,
		migration49,
		migration50,
		migration51,
		migration52,
		migration53,
		migration54,
		migration55,
		migration56,
		migration57,
		migration58,
		migration59,
		migration60,
		migration61,
		migration62,
		migration63,
		migration64,
		migration65,
		migration66,
		migration67,
		migration68,
		migration69,
		migration70,
		migration71,
		migration72,
		migration73,
		migration74,
		migration75,
		migration76,
		migration77,
		migration78,
		migration79,
		migration80,
		migration81,
		migration82,
	}
}

func renameField(db *mongo.Database, coll, from, to string) error {
	_, err := db.Collection(coll).UpdateMany(context.Background(), bson.M{}, bson.M{"$rename": bson.M{from: to}})

	return err
}
