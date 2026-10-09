package context

import (
 "fmt"
 "testing"
 "github.com/free5gc/openapi/models"
)
func TestSessionBindingStopsAtMatchingUE(t *testing.T) {
 c:=&PCFContext{}
 for i:=1;i<=3;i++ {
  ue:=&UeContext{Supi:fmt.Sprint(i),SmPolicyData:map[string]*UeSmPolicyData{}}
  ue.SmPolicyData["10"]=&UeSmPolicyData{PolicyContext:&models.SmPolicyContextData{Ipv4Address:fmt.Sprintf("10.60.0.%d",i),Dnn:"internet",SliceInfo:&models.Snssai{Sst:1,Sd:"010203"}}}
  c.UePool.Store(ue.Supi,ue)
 }
 for repeat:=0;repeat<30;repeat++ { for i:=1;i<=3;i++ {
  ip:=fmt.Sprintf("10.60.0.%d",i)
  p,err:=c.SessionBinding(&models.AppSessionContextReqData{UeIpv4:ip,Dnn:"internet",SliceInfo:&models.Snssai{Sst:1,Sd:"010203"}})
  if err!=nil || p==nil || p.PolicyContext.Ipv4Address!=ip { t.Fatalf("wrong binding for %s: %v",ip,err) }
 } }
 if c.PcfUeFindByIPv4("10.60.0.99")!=nil { t.Fatal("unknown address returned an unrelated UE") }
}
