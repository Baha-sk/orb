/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package proof

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/hyperledger/aries-framework-go/component/storageutil/mem"
	"github.com/stretchr/testify/require"

	"github.com/trustbloc/orb/pkg/activitypub/vocab"
	"github.com/trustbloc/orb/pkg/anchor/handler/mocks"
	"github.com/trustbloc/orb/pkg/anchor/policy"
	proofapi "github.com/trustbloc/orb/pkg/anchor/proof"
	"github.com/trustbloc/orb/pkg/internal/testutil"
	orbmocks "github.com/trustbloc/orb/pkg/mocks"
	"github.com/trustbloc/orb/pkg/pubsub/mempubsub"
	storemocks "github.com/trustbloc/orb/pkg/store/mocks"
	"github.com/trustbloc/orb/pkg/store/vcstatus"
	vcstore "github.com/trustbloc/orb/pkg/store/verifiable"
	"github.com/trustbloc/orb/pkg/store/witness"
)

//go:generate counterfeiter -o ../mocks/monitoring.gen.go --fake-name MonitoringService . monitoringSvc
//go:generate counterfeiter -o ../mocks/vcstatus.gen.go --fake-name VCStatusStore . vcStatusStore

const (
	vcID                     = "http://peer1.com/vc/62c153d1-a6be-400e-a6a6-5b700b596d9d"
	witnessURL               = "http://example.com/orb/services"
	configStoreName          = "orb-config"
	defaultPolicyCacheExpiry = 5 * time.Second
)

func TestNew(t *testing.T) {
	ps := mempubsub.New(mempubsub.Config{})

	store, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
	require.NoError(t, err)

	providers := &Providers{
		AnchorEventStore: store,
	}

	c := New(providers, ps)
	require.NotNil(t, c)
}

func TestWitnessProofHandler(t *testing.T) {
	ps := mempubsub.New(mempubsub.Config{})
	defer ps.Stop()

	witnessIRI, outerErr := url.Parse(witnessURL)
	require.NoError(t, outerErr)

	configStore, outerErr := mem.NewProvider().OpenStore(configStoreName)
	require.NoError(t, outerErr)

	expiryTime := time.Now().Add(60 * time.Second)

	t.Run("success - witness policy not satisfied", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEvent), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    &mockWitnessPolicy{eval: false},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(), expiryTime, []byte(witnessProof))
		require.NoError(t, err)
	})

	t.Run("success - proof expired", func(t *testing.T) {
		proofHandler := New(&Providers{}, ps)

		expiredTime := time.Now().Add(-60 * time.Second)

		err := proofHandler.HandleProof(witnessIRI, vcID, expiredTime, nil)
		require.NoError(t, err)
	})

	t.Run("success - witness policy satisfied", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		witnessPolicy, err := policy.New(configStore, defaultPolicyCacheExpiry)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    witnessPolicy,
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.NoError(t, err)
	})

	t.Run("success - vc status is completed", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusCompleted)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{eval: true},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.NoError(t, err)
	})

	t.Run("success - policy satisfied but some witness proofs are empty", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore: &mockWitnessStore{WitnessProof: []*proofapi.WitnessProof{{
				Type:    proofapi.WitnessTypeSystem,
				Witness: "witness",
			}}},
			WitnessPolicy: &mockWitnessPolicy{eval: true},
			Metrics:       &orbmocks.MetricsProvider{},
			DocLoader:     testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.NoError(t, err)
	})

	t.Run("error - get vc status error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		witnessPolicy, err := policy.New(configStore, defaultPolicyCacheExpiry)
		require.NoError(t, err)

		mockVCStatusStore := &mocks.VCStatusStore{}
		mockVCStatusStore.GetStatusReturns("", fmt.Errorf("get vc status error"))

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      mockVCStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    witnessPolicy,
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), fmt.Sprintf(
			"failed to get status for anchor event [%s]: get vc status error", ae.Anchors().String()))
	})

	t.Run("error - second get vc status error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		witnessPolicy, err := policy.New(configStore, defaultPolicyCacheExpiry)
		require.NoError(t, err)

		mockVCStatusStore := &mocks.VCStatusStore{}
		mockVCStatusStore.GetStatusReturnsOnCall(0, proofapi.VCStatusInProcess, nil)
		mockVCStatusStore.GetStatusReturnsOnCall(1, "", fmt.Errorf("second get vc status error"))

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      mockVCStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    witnessPolicy,
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), fmt.Sprintf(
			"failed to get status for anchor event [%s]: second get vc status error", ae.Anchors().String()))
	})

	t.Run("error - set vc status to complete error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		witnessPolicy, err := policy.New(configStore, defaultPolicyCacheExpiry)
		require.NoError(t, err)

		mockVCStatusStore := &mocks.VCStatusStore{}
		mockVCStatusStore.AddStatusReturns(fmt.Errorf("add vc status error"))

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      mockVCStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    witnessPolicy,
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), fmt.Sprintf(
			"failed to change status to 'completed' for anchor event [%s]: add vc status error", ae.Anchors().String()))
	})

	t.Run("VC status already completed", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		witnessStore, err := witness.New(mem.NewProvider())
		require.NoError(t, err)

		// prepare witness store with 'empty' witness proofs
		emptyWitnessProofs := []*proofapi.WitnessProof{{Type: proofapi.WitnessTypeSystem, Witness: witnessIRI.String()}}
		err = witnessStore.Put(ae.Anchors().String(), emptyWitnessProofs)
		require.NoError(t, err)

		witnessPolicy, err := policy.New(configStore, defaultPolicyCacheExpiry)
		require.NoError(t, err)

		mockVCStatusStore := &mocks.VCStatusStore{}
		mockVCStatusStore.GetStatusReturnsOnCall(0, proofapi.VCStatusInProcess, nil)
		mockVCStatusStore.GetStatusReturnsOnCall(1, proofapi.VCStatusCompleted, nil)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      mockVCStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     witnessStore,
			WitnessPolicy:    witnessPolicy,
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.NoError(t, err)
	})

	t.Run("error - witness policy error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{Err: fmt.Errorf("witness policy error")},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(),
			expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), fmt.Sprintf(
			"failed to evaluate witness policy for anchor event [%s]: witness policy error", ae.Anchors().String()))
	})

	t.Run("error - vc status not found store error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore, // error will be returned b/c we didn't set "in-process" status for vc
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, "testVC",
			expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), "status not found for anchor event[testVC]")
	})

	t.Run("error - store error", func(t *testing.T) {
		store := &storemocks.Store{}
		store.GetReturns(nil, fmt.Errorf("get error"))

		provider := &storemocks.Provider{}
		provider.OpenStoreReturns(store, nil)

		vcStore, err := vcstore.New(provider, testutil.GetLoader(t))
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(vcID, proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, vcID, expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get anchor event: get error")
	})

	t.Run("error - witness store add proof error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{AddProofErr: fmt.Errorf("witness store error")},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(), expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), "witness store error")
	})

	t.Run("error - witness store add proof error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEventTwoProofs), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{GetErr: fmt.Errorf("witness store error")},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(), expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), "witness store error")
	})

	t.Run("error - unmarshal witness proof", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(vcID, proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    &mocks.MonitoringService{},
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, vcID, expiryTime, []byte(""))
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to unmarshal incoming witness proof for anchor event")
	})

	t.Run("error - monitoring error", func(t *testing.T) {
		vcStore, err := vcstore.New(mem.NewProvider(), testutil.GetLoader(t))
		require.NoError(t, err)

		ae := &vocab.AnchorEventType{}
		require.NoError(t, json.Unmarshal([]byte(anchorEvent), ae))

		err = vcStore.Put(ae)
		require.NoError(t, err)

		monitoringSvc := &mocks.MonitoringService{}
		monitoringSvc.WatchReturns(fmt.Errorf("monitoring error"))

		vcStatusStore, err := vcstatus.New(mem.NewProvider())
		require.NoError(t, err)

		err = vcStatusStore.AddStatus(ae.Anchors().String(), proofapi.VCStatusInProcess)
		require.NoError(t, err)

		providers := &Providers{
			AnchorEventStore: vcStore,
			StatusStore:      vcStatusStore,
			MonitoringSvc:    monitoringSvc,
			WitnessStore:     &mockWitnessStore{},
			WitnessPolicy:    &mockWitnessPolicy{},
			Metrics:          &orbmocks.MetricsProvider{},
			DocLoader:        testutil.GetLoader(t),
		}

		proofHandler := New(providers, ps)

		err = proofHandler.HandleProof(witnessIRI, ae.Anchors().String(), expiryTime, []byte(witnessProof))
		require.Error(t, err)
		require.Contains(t, err.Error(), "monitoring error")
	})
}

type mockWitnessStore struct {
	WitnessProof []*proofapi.WitnessProof
	AddProofErr  error
	GetErr       error
}

func (w *mockWitnessStore) AddProof(vcID, witnessID string, proof []byte) error {
	if w.AddProofErr != nil {
		return w.AddProofErr
	}

	return nil
}

func (w *mockWitnessStore) Get(vcID string) ([]*proofapi.WitnessProof, error) {
	if w.GetErr != nil {
		return nil, w.GetErr
	}

	return w.WitnessProof, nil
}

type mockWitnessPolicy struct {
	eval bool
	Err  error
}

func (wp *mockWitnessPolicy) Evaluate(_ []*proofapi.WitnessProof) (bool, error) {
	if wp.Err != nil {
		return false, wp.Err
	}

	return wp.eval, nil
}

//nolint:lll
const anchorEvent = `{
  "@context": "https://w3id.org/activityanchors/v1",
  "anchors": "hl:uEiBL1RVIr2DdyRE5h6b8bPys-PuVs5mMPPC778OtklPa-w",
  "attachment": [
    {
      "contentObject": {
        "properties": {
          "https://w3id.org/activityanchors#generator": "https://w3id.org/orb#v0",
          "https://w3id.org/activityanchors#resources": [
            {
              "id": "did:orb:uAAA:EiD6mH7iCLGjm9mhBr2TP_5_vRz6nyLYZ5E74xbZzrlmLg"
            }
          ]
        },
        "subject": "hl:uEiB1miJeUsG7PiLvFel8DKoluzDVl3OnpjKgAGZS588PXQ:uoQ-BeEJpcGZzOi8vYmFma3JlaWR2dGlyZjR1d2J4bTdjZjN5djVmNmF6a3JmeG15bmxmM3R1NnRkZmlhYW16am9wdHlwbHU"
      },
      "type": "AnchorObject",
      "generator": "https://w3id.org/orb#v0",
      "url": "hl:uEiBL1RVIr2DdyRE5h6b8bPys-PuVs5mMPPC778OtklPa-w",
      "witness": {
        "@context": "https://www.w3.org/2018/credentials/v1",
        "credentialSubject": {
          "id": "hl:uEiBy8pPgN9eS3hpQAwpSwJJvm6Awpsnc8kR_fkbUPotehg"
        },
        "issuanceDate": "2021-01-27T09:30:10Z",
        "issuer": "https://sally.example.com/services/anchor",
        "proof": [
          {
            "created": "2021-01-27T09:30:05Z",
            "domain": "https://witness1.example.com/ledgers/maple2021",
            "jws": "eyJ...",
            "proofPurpose": "assertionMethod",
            "type": "JsonWebSignature2020",
            "verificationMethod": "did:example:abcd#key"
          }
        ],
        "type": "VerifiableCredential"
      }
    }
  ],
  "attributedTo": "https://orb.domain1.com/services/orb",
  "parent": [
    "hl:uEiAsiwjaXOYDmOHxmvDl3Mx0TfJ0uCar5YXqumjFJUNIBg:uoQ-CeEdodHRwczovL2V4YW1wbGUuY29tL2Nhcy91RWlBc2l3amFYT1lEbU9IeG12RGwzTXgwVGZKMHVDYXI1WVhxdW1qRkpVTklCZ3hCaXBmczovL2JhZmtyZWlibXJtZW51eGhnYW9tb2Q0bTI2ZHM1enRkdWp4emhqb2JndnBzeWwydjJuZGNza3EyaWF5",
    "hl:uEiAn3Y7USoP_lNVX-f0EEu1ajLymnqBJItiMARhKBzAKWg:uoQ-CeEdodHRwczovL2V4YW1wbGUuY29tL2Nhcy91RWlBbjNZN1VTb1BfbE5WWC1mMEVFdTFhakx5bW5xQkpJdGlNQVJoS0J6QUtXZ3hCaXBmczovL2JhZmtyZWliaDN3aG5pc3VkNzZrbmt2N3o3dWNiZjNrMnJzNmtuaHZhamVybnJkYWJkYmZhb21ha2xp"
  ],
  "published": "2021-01-27T09:30:10Z",
  "type": "AnchorEvent",
  "url": "hl:uEiCJWrCq8ttsWob5UVueRQiQ_QUrocJY6ZA8BDgzgakuhg:uoQ-BeEJpcGZzOi8vYmFma3JlaWVqbGt5a3Y0dzNucm5pbjZrcmxvcGVrY2VxN3Vjc3hpb2NsZHV6YXBhZWhhenlka2pvcXk"
}`

//nolint:lll
const anchorEventTwoProofs = `{
  "@context": "https://w3id.org/activityanchors/v1",
  "anchors": "hl:uEiBL1RVIr2DdyRE5h6b8bPys-PuVs5mMPPC778OtklPa-w",
  "attachment": [
    {
      "contentObject": {
        "properties": {
          "https://w3id.org/activityanchors#generator": "https://w3id.org/orb#v0",
          "https://w3id.org/activityanchors#resources": [
            {
              "id": "did:orb:uAAA:EiD6mH7iCLGjm9mhBr2TP_5_vRz6nyLYZ5E74xbZzrlmLg"
            }
          ]
        },
        "subject": "hl:uEiB1miJeUsG7PiLvFel8DKoluzDVl3OnpjKgAGZS588PXQ:uoQ-BeEJpcGZzOi8vYmFma3JlaWR2dGlyZjR1d2J4bTdjZjN5djVmNmF6a3JmeG15bmxmM3R1NnRkZmlhYW16am9wdHlwbHU"
      },
      "type": "AnchorObject",
      "generator": "https://w3id.org/orb#v0",
      "url": "hl:uEiBL1RVIr2DdyRE5h6b8bPys-PuVs5mMPPC778OtklPa-w",
      "witness": {
        "@context": "https://www.w3.org/2018/credentials/v1",
        "credentialSubject": {
          "id": "hl:uEiBy8pPgN9eS3hpQAwpSwJJvm6Awpsnc8kR_fkbUPotehg"
        },
        "issuanceDate": "2021-01-27T09:30:10Z",
        "issuer": "https://sally.example.com/services/anchor",
        "proof": [
          {
            "created": "2021-01-27T09:30:00Z",
            "domain": "sally.example.com",
            "jws": "eyJ...",
            "proofPurpose": "assertionMethod",
            "type": "JsonWebSignature2020",
            "verificationMethod": "did:example:abcd#key"
          },
          {
            "created": "2021-01-27T09:30:05Z",
            "domain": "https://witness1.example.com/ledgers/maple2021",
            "jws": "eyJ...",
            "proofPurpose": "assertionMethod",
            "type": "JsonWebSignature2020",
            "verificationMethod": "did:example:abcd#key"
          }
        ],
        "type": "VerifiableCredential"
      }
    }
  ],
  "attributedTo": "https://orb.domain1.com/services/orb",
  "parent": [
    "hl:uEiAsiwjaXOYDmOHxmvDl3Mx0TfJ0uCar5YXqumjFJUNIBg:uoQ-CeEdodHRwczovL2V4YW1wbGUuY29tL2Nhcy91RWlBc2l3amFYT1lEbU9IeG12RGwzTXgwVGZKMHVDYXI1WVhxdW1qRkpVTklCZ3hCaXBmczovL2JhZmtyZWlibXJtZW51eGhnYW9tb2Q0bTI2ZHM1enRkdWp4emhqb2JndnBzeWwydjJuZGNza3EyaWF5",
    "hl:uEiAn3Y7USoP_lNVX-f0EEu1ajLymnqBJItiMARhKBzAKWg:uoQ-CeEdodHRwczovL2V4YW1wbGUuY29tL2Nhcy91RWlBbjNZN1VTb1BfbE5WWC1mMEVFdTFhakx5bW5xQkpJdGlNQVJoS0J6QUtXZ3hCaXBmczovL2JhZmtyZWliaDN3aG5pc3VkNzZrbmt2N3o3dWNiZjNrMnJzNmtuaHZhamVybnJkYWJkYmZhb21ha2xp"
  ],
  "published": "2021-01-27T09:30:10Z",
  "type": "AnchorEvent",
  "url": "hl:uEiCJWrCq8ttsWob5UVueRQiQ_QUrocJY6ZA8BDgzgakuhg:uoQ-BeEJpcGZzOi8vYmFma3JlaWVqbGt5a3Y0dzNucm5pbjZrcmxvcGVrY2VxN3Vjc3hpb2NsZHV6YXBhZWhhenlka2pvcXk"
}`

//nolint:lll
const witnessProof = `{
  "@context": [
    "https://w3id.org/security/v1",
    "https://w3id.org/security/jws/v1"
  ],
  "proof": {
    "created": "2021-04-20T20:05:35.055Z",
    "domain": "http://orb.vct:8077",
    "jws": "eyJhbGciOiJFZERTQSIsImI2NCI6ZmFsc2UsImNyaXQiOlsiYjY0Il19..PahivkKT6iKdnZDpkLu6uwDWYSdP7frt4l66AXI8mTsBnjgwrf9Pr-y_BkEFqsOMEuwJ3DSFdmAp1eOdTxMfDQ",
    "proofPurpose": "assertionMethod",
    "type": "Ed25519Signature2018",
    "verificationMethod": "did:web:abc.com#2130bhDAK-2jKsOXJiEDG909Jux4rcYEpFsYzVlqdAY"
  }
}`
