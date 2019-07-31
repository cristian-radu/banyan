package domain

import (
	"fmt"
	"time"

	banyanv1alpha1 "github.com/cristian-radu/banyan/pkg/apis/banyan/v1alpha1"
	clientset "github.com/cristian-radu/banyan/pkg/generated/clientset/versioned"
	domainscheme "github.com/cristian-radu/banyan/pkg/generated/clientset/versioned/scheme"
	informers "github.com/cristian-radu/banyan/pkg/generated/informers/externalversions/banyan/v1alpha1"
	listers "github.com/cristian-radu/banyan/pkg/generated/listers/banyan/v1alpha1"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
)

const (
	controllerName = "banyan-domain-controller"
	workqueueName  = "banyan-domain"
)

// Controller is responsible for synchronizing Domain objects stored
// in the system with actual DNS Domains registered with the provider.
type Controller struct {
	kubeClientSet   kubernetes.Interface
	domainClientSet clientset.Interface
	eventRecorder   record.EventRecorder
	lister          listers.DomainLister
	workqueue       workqueue.RateLimitingInterface
}

// NewController creates a new Domain Controller.
func NewController(kubeClientSet kubernetes.Interface, domainClientSet clientset.Interface, domainInformer informers.DomainInformer) (*Controller, error) {

	utilruntime.Must(domainscheme.AddToScheme(scheme.Scheme))
	log.Info("Creating event broadcaster")

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(log.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeClientSet.CoreV1().Events("")})

	c := &Controller{
		kubeClientSet:   kubeClientSet,
		domainClientSet: domainClientSet,
		eventRecorder:   eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerName}),
		workqueue:       workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), workqueueName),
		lister:          domainInformer.Lister(),
	}

	log.Info("Setting up event handlers")
	domainInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.addDomain,
		UpdateFunc: c.updateDomain,
		DeleteFunc: c.deleteDomain,
	})

	return c, nil
}

// Run begins watching for domain changes and syncs them.
func (c *Controller) Run(workers int, stopCh <-chan struct{}) error {
	defer c.workqueue.ShutDown()

	log.Infof("Starting domain controller")
	defer log.Infof("Shutting down domain controller")

	log.Info("Starting workers")

	for i := 0; i < workers; i++ {
		go wait.Until(c.worker, time.Second, stopCh)
	}

	log.Info("Workers ready")

	<-stopCh

	log.Info("Workers shutting down")

	return nil
}

func (c *Controller) addDomain(obj interface{}) {
	d := obj.(*banyanv1alpha1.Domain)
	log.Infof("Adding domain %s", d.Spec.Name)
	c.enqueueDomain(d)
}

func (c *Controller) updateDomain(old, cur interface{}) {
	oldD := old.(*banyanv1alpha1.Domain)
	curD := cur.(*banyanv1alpha1.Domain)
	log.Infof("Updating domain %s", oldD.Spec.Name)
	c.enqueueDomain(curD)
}

func (c *Controller) deleteDomain(obj interface{}) {
	d := obj.(*banyanv1alpha1.Domain)
	log.Infof("Deleting domain %s", d.Spec.Name)
	c.enqueueDomain(d)
}

func (c *Controller) enqueueDomain(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

func (c *Controller) worker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	key, quit := c.workqueue.Get()

	if quit {
		return false
	}
	defer c.workqueue.Done(key)

	err := c.syncHandler(key.(string))
	if err == nil {
		c.workqueue.Forget(key)
		return true
	}

	utilruntime.HandleError(fmt.Errorf("Sync %q failed with %v", key, err))
	c.workqueue.AddRateLimited(key)

	return true
}

func (c *Controller) syncHandler(key string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	log.Info("syncing domain: %s in namespace: %s", name, namespace)

	// ToDo: Check if the domain exists in the cache. Sync could have been triggered by an object deletion.
	// domain, err := c.lister.Domains(namespace).Get(name)
	// if err != nil {
	// 	if errors.IsNotFound(err) {
	// Check if we own this domain and delete it from the provider if we do.
	// Check for errors during deletion. Return the error to requeue the domain for processing.
	// Record an info event.
	// 		return nil
	// 	}
	// 	return err
	// }

	// ToDo: Try to get the domain from the provider.
	// Create the domain if not found.
	// Check for errors and return them if we want to requeue the domain for processing.
	// Record an info event.
	// Set ownership because we are now managing it.

	// ToDo: If found, check if we own it before touching it.
	// Record a warning event if we don't own it. Log a warning message too.
	// c.eventRecorder.Event(domain, corev1.EventTypeWarning, ErrResourceExists, msg)
	// Update the domain if we do own it.
	// Check for errors during update and return to requeue the domain for processing.
	// Update the status of the domain object.

	return nil
}
