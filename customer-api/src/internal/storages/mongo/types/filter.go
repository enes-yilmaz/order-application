package types

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type FindCustomersFilter struct {
	CustomerId     string `bson:"_id"`
	Sorting        string `query:"sorting"`
	Offset         int64
	Limit          int64
	IsCount        bool
	StartDate      time.Time
	EndDate        time.Time
	UpdatedAt      time.Time
	OrderBy        string
	OrderDirection int
}

func (f *FindCustomersFilter) BuildCustomersQueryPipeline() []bson.M {

	var pipelineBuild []bson.M

	//greaterThanFilter := bson.M{
	//	"$match": bson.M{
	//		"_id":         bson.M{"$gte": 4.5},
	//		"reviewCount": bson.M{"$gte": 1000},
	//	},
	//}
	//sample := bson.M{
	//	"$sample": bson.M{
	//		"size": f.Limit,
	//	},
	//}
	//pipelineBuild = append(pipelineBuild, greaterThanFilter)
	//pipelineBuild = append(pipelineBuild, sample)

	var facetStagePartitions bson.M

	var facetTotalCount []bson.M
	var facetItems []bson.M

	facetTotalCount = append(facetTotalCount, bson.M{"$count": "count"})

	if f.OrderBy == "name" {
		facetItems = append(facetItems, bson.M{
			"$sort": bson.M{
				"name": f.OrderDirection,
			},
		})
	} else {
		facetItems = append(facetItems, bson.M{
			"$sort": bson.M{
				"email": f.OrderDirection,
			},
		})
	}

	facetItems = append(facetItems,
		bson.M{
			"$skip": f.Offset,
		},
		bson.M{
			"$limit": f.Limit,
		},
	)

	if !f.IsCount {
		facetStagePartitions = bson.M{"items": facetItems, "totalCount": facetTotalCount}
	} else {
		facetStagePartitions = bson.M{"totalCount": facetTotalCount}
	}

	facetStage := bson.M{"$facet": facetStagePartitions}
	pipelineBuild = append(pipelineBuild, facetStage)

	return pipelineBuild
}
