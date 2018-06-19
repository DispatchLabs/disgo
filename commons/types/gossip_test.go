/*
 *    This file is part of Disgo-Commons library.
 *
 *    The Disgo-Commons library is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    The Disgo-Commons library is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with the Disgo-Commons library.  If not, see <http://www.gnu.org/licenses/>.
 */
package types

import (
	"fmt"
	"reflect"
	"testing"
)

func testMockNewGossip(t *testing.T) (*Gossip, *Transaction) {
	tx := testMockTransaction(t)
	receipt := NewReceipt("test")
	gossip := NewGossip(*tx, *receipt)
	return gossip, tx
}

func TestNewGossip(t *testing.T) {
	gossip, tx := testMockNewGossip(t)
	testGossipStruct(t, gossip, tx)
}

func TestToGossipFromJson(t *testing.T) {
	receipt := NewReceipt(RequestNewTransaction)
	tx := testMockTransaction(t)
	s := fmt.Sprintf(`{"ReceiptId":"%s","Transaction":%s,"Rumors":[]}`, receipt.Id, tx.String())
	gossip, err := ToGossipFromJson([]byte(s))
	if err != nil {
		t.Fatalf("ToGossipFromJson returning error: %s", err)
	}
	if gossip.Transaction.String() != tx.String() {
		t.Errorf("gossip contains invalid %s value.\nG: %s\nE: %s", "Transaction", gossip.Transaction.String(), tx.String())
	}
	if gossip.Rumors == nil {
		t.Errorf("gossip contains invalid %s value: %s", "Rumors", gossip.Rumors)
	}
	//if gossip.Status != "Gossiping" {
	//	t.Errorf("gossip contains invalid %s value: %s", "Status", gossip.Status)
	//}
	if gossip.Key() != fmt.Sprintf("table-gossip-%s", tx.Hash) {
		t.Errorf("gossip.Key() returning invalid %s value: %s", "Key", gossip.Key())
	}
	if gossip.String() != s {
		t.Errorf("gossip.String() returning invalid value.\nG: %s\nE: %s", gossip.String(), s)
	}
}

func TestToJsonByGossip(t *testing.T) {
	gossip, tx := testMockNewGossip(t)
	bytes, err := ToJsonByGossip(*gossip)
	if err != nil {
		t.Fatalf("ToJsonByGossip returning error: %s", err)
	}
	s := fmt.Sprintf(`{"ReceiptId":"%s","Transaction":%s,"Rumors":[]}`, gossip.ReceiptId, tx.String())
	if string(bytes) != s {
		t.Errorf("ToJsonByGossip returning invalid value.\nG: %s\nE: %s", string(bytes), s)
	}
}

func TestContainsRumor(t *testing.T) {
	gossip, _ := testMockNewGossip(t)
	r1 := testMockRumor()
	r2 := NewRumor("0f86ea981203b26b5b8244c8f661e30e5104555068a4bd168d3e3015db9bb25a", "3ed25f42484d517cdfc72cafb7ebc9e8baa52c2b", "9c242afd4f2dcaedcfb0cff2bb9c38b5811ed29c249f5b49f7759642a473d5fb")
	gossip.Rumors = append(gossip.Rumors, *r1)
	if gossip.ContainsRumor(r1.Address) != true {
		t.Errorf("gossip.ContainsRumor returning invalid value.\nGot: %t\nExpected: %t", gossip.ContainsRumor(r1.Address), true)
	}
	if gossip.ContainsRumor(r2.Address) != false {
		t.Errorf("gossip.ContainsRumor returning invalid value.\nGot: %t\nExpected: %t", gossip.ContainsRumor(r2.Address), false)
	}
}

func TestToGossipByKey(t *testing.T) {
	// TODO: ToGossipByKey()
	t.Skip("Need a Badger DB mock")
}

func TestToGossips(t *testing.T) {
	// TODO: ToGossips()
	t.Skip("Need a Badger DB mock")
}

func TestToOldGossips(t *testing.T) {
	// TODO: ToOldGossips()
	t.Skip("Need a Badger DB mock")
}

func TestGossipSet(t *testing.T) {
	// TODO: Gossip.Set()
	t.Skip("Need a Badger DB mock")
}

func TestGossipRefresh(t *testing.T) {
	// TODO: Gossip.Refresh()
	t.Skip("Need a Badger DB mock")
}

func testGossipStruct(t *testing.T, gossip *Gossip, tx *Transaction) {
	if reflect.DeepEqual(gossip.Transaction, *tx) == false {
		t.Errorf("gossip contains invalid %s value.\nG: %s\nE: %s", "Transaction", gossip.Transaction, *tx)
	}
	if gossip.Rumors == nil {
		t.Errorf("gossip contains invalid %s value: %s", "Rumors", gossip.Rumors)
	}
	//if gossip.Status != "Gossiping" {
	//	t.Errorf("gossip contains invalid %s value: %s", "Status", gossip.Status)
	//}
	if gossip.Key() != fmt.Sprintf("table-gossip-%s", tx.Hash) {
		t.Errorf("gossip.Key() returning invalid %s value: %s", "Key", gossip.Key())
	}
	s := fmt.Sprintf(`{"ReceiptId":"%s","Transaction":%s,"Rumors":[]}`, gossip.ReceiptId, tx.String())
	if gossip.String() != s {
		t.Errorf("gossip.String() returning invalid value.\nG: %s\nE: %s", gossip.String(), s)
	}
}
