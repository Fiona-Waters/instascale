package e2e

import (
	"context"
	"fmt"
	"os"
	"testing"

	ocmsdk "github.com/openshift-online/ocm-sdk-go"
	cmv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	mcadv1beta1 "github.com/project-codeflare/multi-cluster-app-dispatcher/pkg/apis/controller/v1beta1"
	"google.golang.org/grpc/balancer/grpclb/state"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sclient "sigs.k8s.io/controller-runtime/pkg/client"

)

func TestScaleMachinePool(t *testing.T) {
// connect to openshift cluster
	ctx := context.Background()

	token := os.Getenv("KUBECONFIG")
	logger, err := ocmsdk.NewGoLoggerBuilder().
		Debug(false).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build logger: %v\n", err)
		os.Exit(1)
	}
	connection, err := ocmsdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		os.Exit(1)
	}
	defer connection.Close()

	// TODO set instascale to true?
	// create an appwrapper 
	aw := &mcadv1beta1.AppWrapper{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test-scale-machinepool",
			Namespace: "test",
			Labels: map[string]string{
				"orderedinstance": "m5.xlarge_g4dn.xlarge",
			},
		},
		Spec: mcadv1beta1.AppWrapperSpec{
			AggrResources: mcadv1beta1.AppWrapperResourceList{
				GenericItems: []mcadv1beta1.AppWrapperGenericResource{
					{
						CustomPodResources: []mcadv1beta1.CustomPodResourceTemplate{
							{
								Replicas: 1,
							},
							{
								Replicas: 1,
							},
						},
					},
				},
			},
		},
	}

	collection := connection.ClustersMgmt().V1().Clusters()
	createAW := 
	// log in to openshift cluster
	// create an appwrapper (use an example with instascale)
	// get appwrapper/s
	// get resource requirements
	// scale up resources
	// check that job is complete
	// check that resources are removed (are scaling down?)

// resources before :=
// resources after :=
// assert that before != after

}
