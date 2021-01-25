package informers

import (
	"log"

	broker "github.com/layer5io/meshsync/pkg/broker"
	"github.com/layer5io/meshsync/pkg/model"
	v1beta1 "istio.io/client-go/pkg/apis/security/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

func (i *Istio) AuthorizationPolicyInformer() cache.SharedIndexInformer {
	// get informer
	AuthorizationPolicyInformer := i.client.GetAuthorizationPolicyInformer().Informer()

	// register event handlers
	AuthorizationPolicyInformer.AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				AuthorizationPolicy := obj.(*v1beta1.AuthorizationPolicy)
				log.Printf("AuthorizationPolicy Named: %s - added", AuthorizationPolicy.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						metav1.TypeMeta{
							Kind:       "AuthorizationPolicy",
							APIVersion: "v1beta1",
						},
						AuthorizationPolicy.ObjectMeta,
						AuthorizationPolicy.Spec,
						AuthorizationPolicy.Status,
					)})
				if err != nil {
					log.Println("BROKER: Error publishing AuthorizationPolicy")
				}
			},
			UpdateFunc: func(new interface{}, old interface{}) {
				AuthorizationPolicy := new.(*v1beta1.AuthorizationPolicy)
				log.Printf("AuthorizationPolicy Named: %s - updated", AuthorizationPolicy.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						metav1.TypeMeta{
							Kind:       "AuthorizationPolicy",
							APIVersion: "v1beta1",
						},
						AuthorizationPolicy.ObjectMeta,
						AuthorizationPolicy.Spec,
						AuthorizationPolicy.Status,
					)})
				if err != nil {
					log.Println("BROKER: Error publishing AuthorizationPolicy")
				}
			},
			DeleteFunc: func(obj interface{}) {
				AuthorizationPolicy := obj.(*v1beta1.AuthorizationPolicy)
				log.Printf("AuthorizationPolicy Named: %s - deleted", AuthorizationPolicy.Name)
				err := i.broker.Publish(Subject, &broker.Message{
					Object: model.ConvObject(
						metav1.TypeMeta{
							Kind:       "AuthorizationPolicy",
							APIVersion: "v1beta1",
						},
						AuthorizationPolicy.ObjectMeta,
						AuthorizationPolicy.Spec,
						AuthorizationPolicy.Status,
					)})
				if err != nil {
					log.Println("BROKER: Error publishing AuthorizationPolicy")
				}
			},
		},
	)

	return AuthorizationPolicyInformer
}
