package handler

import ("github.com/labstack/echo"
	"github.com/kelvins/geocoder"
	"strconv"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"net/http"
	"time"
	"log"
)
type State struct {
	_id string
	state string
	active string
	confirmed string
	death string
	last_updated_time string
}
func(h *Handler) GetCasesFromGeoLocation(c echo.Context)(err error) {

	fmt.Printf(c.QueryParam("latitude"))
	fmt.Printf(c.QueryParam("longitude"))


	f1, err := strconv.ParseFloat(c.QueryParam("latitude"),8)
	f2, err := strconv.ParseFloat(c.QueryParam("longitude"),8)

	var location = geocoder.Location{
		Latitude:  f1,
		Longitude: f2,
	}

	addresses, err := geocoder.GeocodingReverse(location)

	if err != nil {
		fmt.Println("Could not get the addresses: ", err)
	} else {
		var address = addresses[0]
		fmt.Println(address.FormatAddress())
		fmt.Println(address.FormattedAddress)
		fmt.Println(address.Types)
	}

	return


}


func(h *Handler) GetCasesFromStateCode(c echo.Context)(err error) {

	var result State
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	print(c.QueryParam("stateCode"))
	fmt.Println("find one result:", result)
	filterCursor, err := h.DB.Collection("state").Find(ctx, bson.M{"_id": c.QueryParam("stateCode")})
	if err != nil {
		log.Fatal(err)
	}
	var filteredDoc []bson.M
	if err = filterCursor.All(ctx, &filteredDoc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("requested doc:  ", filteredDoc)


	//     TO traverse over the whole collection

	/*


	cursor, err := h.DB.Collection("state").Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			} else {
				fmt.Println("\nresult type:", reflect.TypeOf(result))
				fmt.Println("result:", result)
			}
		}
	}


	*/



	return c.JSON(http.StatusOK, filteredDoc)
}
