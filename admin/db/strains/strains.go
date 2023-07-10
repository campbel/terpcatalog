package strains

import (
	"context"

	"github.com/campbel/terpcatalog/admin/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Strain struct {
	ID          primitive.ObjectID `bson:"_id"`
	ProducerID  primitive.ObjectID `bson:"producer_id"`
	Name        string             `bson:"name"`
	Category    string             `bson:"category"`
	Genetics    string             `bson:"genetics"`
	THC         float64            `bson:"thc"`
	Terpenes    float64            `bson:"terpenes"`
	Price       float64            `bson:"price"`
	HarvestDate string             `bson:"harvest_date"`
	Images      []string           `bson:"images"`
}

type Store interface {
	GetStrains(ctx context.Context) ([]types.Strain, error)
	GetStrain(ctx context.Context, id string) (types.Strain, error)
	CreateStrain(ctx context.Context, strain types.Strain) (types.Strain, error)
	UpdateStrain(ctx context.Context, id string, strain types.Strain) (types.Strain, error)
	DeleteStrain(ctx context.Context, id string) error
}

type Collection struct {
	collection *mongo.Collection
}

func NewStore(collection *mongo.Collection) *Collection {
	return &Collection{collection: collection}
}

func (db *Collection) GetStrains(ctx context.Context) ([]types.Strain, error) {
	var strains []Strain
	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &strains)
	if err != nil {
		return nil, err
	}
	return translateToCommonSlice(strains), nil
}

func (db *Collection) GetStrain(ctx context.Context, id string) (types.Strain, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Strain{}, err
	}
	var strain Strain
	err = db.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&strain)
	if err != nil {
		return types.Strain{}, err
	}
	return translateToCommon(strain), nil
}

func (db *Collection) CreateStrain(ctx context.Context, strain types.Strain) (types.Strain, error) {
	strain.ID = primitive.NewObjectID().Hex()
	_, err := db.collection.InsertOne(ctx, translateFromCommon(strain))
	if err != nil {
		return types.Strain{}, err
	}
	return strain, nil
}

func (db *Collection) UpdateStrain(ctx context.Context, id string, strain types.Strain) (types.Strain, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Strain{}, err
	}
	_, err = db.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": translateFromCommon(strain)})
	if err != nil {
		return types.Strain{}, err
	}
	return strain, nil
}

func (db *Collection) DeleteStrain(ctx context.Context, id string) error {
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

func translateToCommon(strain Strain) types.Strain {
	return types.Strain{
		ID:          strain.ID.Hex(),
		ProducerID:  strain.ProducerID.Hex(),
		Name:        strain.Name,
		Category:    strain.Category,
		Genetics:    strain.Genetics,
		THC:         strain.THC,
		Terpenes:    strain.Terpenes,
		Price:       strain.Price,
		HarvestDate: strain.HarvestDate,
		Images:      strain.Images,
	}
}

func translateToCommonSlice(strains []Strain) []types.Strain {
	commonStrains := make([]types.Strain, len(strains))
	for i, strain := range strains {
		commonStrains[i] = translateToCommon(strain)
	}
	return commonStrains
}

func translateFromCommon(strain types.Strain) Strain {
	id, _ := primitive.ObjectIDFromHex(strain.ID)
	pid, _ := primitive.ObjectIDFromHex(strain.ProducerID)
	return Strain{
		ID:          id,
		ProducerID:  pid,
		Name:        strain.Name,
		Category:    strain.Category,
		Genetics:    strain.Genetics,
		THC:         strain.THC,
		Terpenes:    strain.Terpenes,
		Price:       strain.Price,
		HarvestDate: strain.HarvestDate,
		Images:      strain.Images,
	}
}
