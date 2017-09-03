/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package awsSDK

import (
	"encoding/xml"
)

type EC2 struct {
	XMLName        xml.Name      `xml:"Response"`
	RequestID      string        `xml:"RequestID"`
	NextToken      string        `xml:"nextToken"`
	ReservationSet []Reservation `xml:"reservationSet>item"`
}

type Reservation struct {
	OwnerId       string `xml:"ownerId"`
	ReservationId string `xml:"reservationId"`
	// InstancesSet  []Instance `xml:"instancesSet>item"`
}
