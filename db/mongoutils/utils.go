package mongoutils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HexToObjectId(s string) primitive.ObjectID {

	objId, _ := primitive.ObjectIDFromHex(s)

	return objId
}

func CalcPageOffset(rowsPerPage, pageNo int64) int64 {

	if pageNo < 1 {
		return 0
	}

	return rowsPerPage * (pageNo - 1)
}
