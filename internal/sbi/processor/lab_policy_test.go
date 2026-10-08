package processor

import (
	"github.com/free5gc/openapi/models"
	"testing"
)

func TestLabPolicyDeltaIsolatesExistingReservation(t *testing.T) {
	d := &models.SmPolicyDecision{PccRules: map[string]*models.PccRule{"active": {PccRuleId: "active", RefQosData: []string{"qa"}}, "new": {PccRuleId: "new", RefQosData: []string{"qn"}}}, QosDecs: map[string]*models.QosData{"qa": {QosId: "qa", GbrDl: "20000 Kbps"}, "qn": {QosId: "qn", GbrDl: "20000 Kbps"}}}
	delta := labPolicyDelta(d, map[string]string{"1-1": "new"})
	if len(delta.PccRules) != 1 || delta.PccRules["active"] != nil || len(delta.QosDecs) != 1 || delta.QosDecs["qa"] != nil {
		t.Fatal("existing reservation included in new admission")
	}
	delta.QosDecs["qn"].GbrDl = "changed"
	if d.QosDecs["qn"].GbrDl != "20000 Kbps" {
		t.Fatal("asynchronous notification aliases live policy")
	}
}
