package orders

import (
	"context"

	"github.com/campbel/terpcatalog/shared/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id"`
	Information OrderInfo          `bson:"information"`
	Items       []OrderItem        `bson:"items"`
}

type OrderInfo struct {
	Company       string       `bson:"company_name"`
	LicenseNumber string       `bson:"license_number"`
	Email         string       `bson:"email"`
	Phone         string       `bson:"phone"`
	Address       OrderAddress `bson:"address"`
}

type OrderAddress struct {
	Street string `bson:"street"`
	City   string `bson:"city"`
	State  string `bson:"state"`
	Postal string `bson:"postal"`
}

type OrderItem struct {
	StrainID string `bson:"strain_id"`
	Quantity int    `bson:"quantity"`
}

type Store interface {
	GetOrders(ctx context.Context) ([]types.Order, error)
	GetOrder(ctx context.Context, id string) (types.Order, error)
	CreateOrder(ctx context.Context, order types.Order) (types.Order, error)
	UpdateOrder(ctx context.Context, id string, order types.Order) (types.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}

type Collection struct {
	collection *mongo.Collection
}

func NewStore(collection *mongo.Collection) *Collection {
	return &Collection{collection: collection}
}

func (db *Collection) GetOrders(ctx context.Context) ([]types.Order, error) {
	var orders []Order
	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &orders)
	if err != nil {
		return nil, err
	}
	return translateToCommonSlice(orders), nil
}

func (db *Collection) GetOrder(ctx context.Context, id string) (types.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Order{}, err
	}
	var order Order
	err = db.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		return types.Order{}, err
	}
	return translateToCommon(order), nil
}

func (db *Collection) CreateOrder(ctx context.Context, order types.Order) (types.Order, error) {
	order.ID = primitive.NewObjectID().Hex()
	_, err := db.collection.InsertOne(ctx, translateFromCommon(order))
	if err != nil {
		return types.Order{}, err
	}
	return order, nil
}

func (db *Collection) UpdateOrder(ctx context.Context, id string, order types.Order) (types.Order, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return types.Order{}, err
	}
	_, err = db.collection.UpdateOne(ctx, bson.M{"_id": objectID}, bson.M{"$set": translateFromCommon(order)})
	if err != nil {
		return types.Order{}, err
	}
	return order, nil
}

func (db *Collection) DeleteOrder(ctx context.Context, id string) error {
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

func translateToCommon(order Order) types.Order {
	return types.Order{
		ID:          order.ID.Hex(),
		Information: translateToCommonInfo(order.Information),
		Items:       translateToCommonItems(order.Items),
	}
}

func translateToCommonInfo(info OrderInfo) types.OrderInfo {
	return types.OrderInfo{
		Company:       info.Company,
		LicenseNumber: info.LicenseNumber,
		Email:         info.Email,
		Phone:         info.Phone,
		Address: types.OrderAddress{
			Street: info.Address.Street,
			City:   info.Address.City,
			State:  info.Address.State,
			Postal: info.Address.Postal,
		},
	}
}

func translateToCommonItems(items []OrderItem) []types.OrderItem {
	commonItems := make([]types.OrderItem, len(items))
	for i, item := range items {
		commonItems[i] = types.OrderItem{
			StrainID: item.StrainID,
			Quantity: item.Quantity,
		}
	}
	return commonItems
}

func translateToCommonSlice(orders []Order) []types.Order {
	commonOrders := make([]types.Order, len(orders))
	for i, order := range orders {
		commonOrders[i] = translateToCommon(order)
	}
	return commonOrders
}

func translateFromCommon(order types.Order) Order {
	id, _ := primitive.ObjectIDFromHex(order.ID)
	return Order{
		ID:          id,
		Information: translateFromCommonInfo(order.Information),
		Items:       translateFromCommonItems(order.Items),
	}
}

func translateFromCommonInfo(info types.OrderInfo) OrderInfo {
	return OrderInfo{
		Company:       info.Company,
		LicenseNumber: info.LicenseNumber,
		Email:         info.Email,
		Phone:         info.Phone,
		Address: OrderAddress{
			Street: info.Address.Street,
			City:   info.Address.City,
			State:  info.Address.State,
			Postal: info.Address.Postal,
		},
	}
}

func translateFromCommonItems(items []types.OrderItem) []OrderItem {
	commonItems := make([]OrderItem, len(items))
	for i, item := range items {
		commonItems[i] = OrderItem{
			StrainID: item.StrainID,
			Quantity: item.Quantity,
		}
	}
	return commonItems
}
