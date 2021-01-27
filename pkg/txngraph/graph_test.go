/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package txngraph

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/trustbloc/sidetree-core-go/pkg/mocks"

	"github.com/trustbloc/orb/pkg/api/txn"
)

const testDID = "did:method:abc"

func TestNew(t *testing.T) {
	log := New(mocks.NewMockCasClient(nil))
	require.NotNil(t, log)
}

func TestGraph_Add(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txnInfo := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString: "anchor",
				Namespace:    "namespace",
				Version:      1,
			},
		}

		cid, err := graph.Add(txnInfo)
		require.NoError(t, err)
		require.NotNil(t, cid)
	})
}

func TestGraph_Read(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txnInfo := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString: "anchor",
				Namespace:    "namespace",
				Version:      1,
			},
		}

		txnCID, err := graph.Add(txnInfo)
		require.NoError(t, err)
		require.NotNil(t, txnCID)

		txnNode, err := graph.Read(txnCID)
		require.NoError(t, err)
		require.Equal(t, txnInfo, txnNode)
	})

	t.Run("error - transaction (cid) not found", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txnNode, err := graph.Read("non-existent")
		require.Error(t, err)
		require.Nil(t, txnNode)
	})
}

func TestGraph_GetDidTransactions(t *testing.T) {
	t.Run("success - first did transaction (create), no previous did transaction", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txnInfo := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString: "anchor",
				Namespace:    "namespace",
				Version:      1,
			},
		}

		txnCID, err := graph.Add(txnInfo)
		require.NoError(t, err)
		require.NotNil(t, txnCID)

		didTxns, err := graph.GetDidTransactions(txnCID, testDID)
		require.NoError(t, err)
		require.Equal(t, 0, len(didTxns))
	})

	t.Run("success - previous transaction for did exists", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txn1 := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString: "anchor-1",
				Namespace:    "namespace",
				Version:      1,
			},
		}

		txn1CID, err := graph.Add(txn1)
		require.NoError(t, err)
		require.NotNil(t, txn1)

		testDID := "did:method:abc"

		previousDIDTxns := make(map[string]string)
		previousDIDTxns[testDID] = txn1CID

		txnInfo := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString:   "anchor-2",
				Namespace:      "namespace",
				Version:        1,
				PreviousDidTxn: previousDIDTxns,
			},
		}

		txnCID, err := graph.Add(txnInfo)
		require.NoError(t, err)
		require.NotNil(t, txnCID)

		didTxns, err := graph.GetDidTransactions(txnCID, testDID)
		require.NoError(t, err)
		require.Equal(t, 1, len(didTxns))
		require.Equal(t, txn1CID, didTxns[0])
	})

	t.Run("error - cid referenced in previous transaction not found", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		testDID := "did:method:abc"

		previousDIDTxns := make(map[string]string)
		previousDIDTxns[testDID] = "non-existent"

		txnInfo := &txn.OrbTransaction{
			Payload: txn.Payload{
				AnchorString:   "anchor-2",
				Namespace:      "namespace",
				Version:        1,
				PreviousDidTxn: previousDIDTxns,
			},
		}

		txnCID, err := graph.Add(txnInfo)
		require.NoError(t, err)
		require.NotNil(t, txnCID)

		didTxns, err := graph.GetDidTransactions(txnCID, testDID)
		require.Error(t, err)
		require.Nil(t, didTxns)
		require.Contains(t, err.Error(), "not found")
	})

	t.Run("error - head cid not found", func(t *testing.T) {
		graph := New(mocks.NewMockCasClient(nil))

		txnNode, err := graph.GetDidTransactions("non-existent", "did")
		require.Error(t, err)
		require.Nil(t, txnNode)
		require.Contains(t, err.Error(), "not found")
	})
}
