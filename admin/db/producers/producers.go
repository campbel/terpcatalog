package producers

import (
	"context"

	"github.com/campbel/terpcatalog/admin/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Producer struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Website string             `bson:"website"`
	Email   string             `bson:"email"`
	Phone   string             `bson:"phone"`
	Address string             `bson:"address"`
	Contact ProducerContact    `bson:"contact"`
}

type ProducerContact struct {
	FirstName string `bson:"first_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
	Phone     string `bson:"phone"`
}

type Store interface {
	GetProducers(ctx context.Context) ([]types.Producer, error)
	GetProducer(ctx context.Context, id string) (types.Producer, error)
	CreateProducer(ctx context.Context, producer types.Producer) (types.Producer, error)
	UpdateProducer(ctx context.Context, id string, producer types.Producer) (types.Producer, error)
	DeleteProducer(ctx context.Context, id string) error
}

type Collection struct {
	collection *mongo.Collection
}

func NewStore(collection *mongo.Collection) *Collection {
	return &Collection{collection: collection}
}

func (db *Collection) GetProducers(ctx context.Context) ([]types.Producer, error) {
	var producers []Producer
	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &producers)
	if err != nil {
		return nil, err
	}
	return translateToCommonSlice(producers), nil
}

func (db *Collection) GetProducer(ctx context.Context, id string) (types.Producer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Producer{}, err
	}
	var producer Producer
	err = db.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&producer)
	if err != nil {
		return types.Producer{}, err
	}
	return translateToCommon(producer), nil
}

func (db *Collection) CreateProducer(ctx context.Context, producer types.Producer) (types.Producer, error) {
	producer.ID = primitive.NewObjectID().Hex()
	_, err := db.collection.InsertOne(ctx, translateFromCommon(producer))
	if err != nil {
		return types.Producer{}, err
	}
	return producer, nil
}

func (db *Collection) UpdateProducer(ctx context.Context, id string, producer types.Producer) (types.Producer, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Producer{}, err
	}
	_, err = db.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": translateFromCommon(producer)})
	if err != nil {
		return types.Producer{}, err
	}
	return producer, nil
}

func (db *Collection) DeleteProducer(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = db.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	return nil
}

func translateToCommon(producer Producer) types.Producer {
	return types.Producer{
		ID:      producer.ID.Hex(),
		Name:    producer.Name,
		Website: producer.Website,
		Email:   producer.Email,
		Phone:   producer.Phone,
		Address: producer.Address,
		Contact: types.ProducerContact{
			FirstName: producer.Contact.FirstName,
			LastName:  producer.Contact.LastName,
			Email:     producer.Contact.Email,
			Phone:     producer.Contact.Phone,
		},
	}
}

func translateToCommonSlice(producers []Producer) []types.Producer {
	commonProducers := make([]types.Producer, len(producers))
	for i, producer := range producers {
		commonProducers[i] = translateToCommon(producer)
	}
	return commonProducers
}

func translateFromCommon(producer types.Producer) Producer {
	id, _ := primitive.ObjectIDFromHex(producer.ID)
	return Producer{
		ID:      id,
		Name:    producer.Name,
		Website: producer.Website,
		Email:   producer.Email,
		Phone:   producer.Phone,
		Address: producer.Address,
		Contact: ProducerContact{
			FirstName: producer.Contact.FirstName,
			LastName:  producer.Contact.LastName,
			Email:     producer.Contact.Email,
			Phone:     producer.Contact.Phone,
		},
	}
}
