package k8x

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	Clientset *kubernetes.Clientset
	namespace string
}

func NewClient(kubeconfig string, namespace string) (*Client, error) {
	clientConfig, err := clientcmd.NewClientConfigFromBytes([]byte(kubeconfig))
	if err != nil {
		return nil, err
	}

	config, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		Clientset: clientset,
		namespace: namespace,
	}, nil
}

// GetDeployment 根据名称查询deployment信息
func (p *Client) GetDeployment(deploymentName string) (*appsv1.Deployment, error) {
	return p.Clientset.AppsV1().Deployments(p.namespace).Get(context.Background(), deploymentName, metav1.GetOptions{})
}

// QueryDeployment 查询deployment列表
func (p *Client) QueryDeployment(limit int64) ([]appsv1.Deployment, error) {
	if limit < 1 {
		limit = 10
	}
	back, err := p.Clientset.AppsV1().Deployments(p.namespace).List(context.Background(), metav1.ListOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	return back.Items, nil
}

// UpdateDeployment 更新deployment
func (p *Client) UpdateDeployment(deploymentName string, deployment *appsv1.Deployment) error {
	_, err := p.Clientset.AppsV1().Deployments(p.namespace).Update(context.Background(), deployment, metav1.UpdateOptions{})
	return err
}

// CreateDeployment 创建deployment
func (p *Client) CreateDeployment(deployment *appsv1.Deployment) error {
	_, err := p.Clientset.AppsV1().Deployments(p.namespace).Create(context.Background(), deployment, metav1.CreateOptions{})
	return err
}

// DeleteDeployment 删除deployment
func (p *Client) DeleteDeployment(deploymentName string) error {
	return p.Clientset.AppsV1().Deployments(p.namespace).Delete(context.Background(), deploymentName, metav1.DeleteOptions{})
}

// GetService 获取service信息
func (p *Client) GetService(serviceName string) (*corev1.Service, error) {
	return p.Clientset.CoreV1().Services(p.namespace).Get(context.Background(), serviceName, metav1.GetOptions{})
}

// QueryService 查询servic列表
func (p *Client) QueryService(limit int64) ([]corev1.Service, error) {
	back, err := p.Clientset.CoreV1().Services(p.namespace).List(context.Background(), metav1.ListOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}
	return back.Items, nil
}

// CreateService 创建service
func (p *Client) CreateService(service *corev1.Service) error {
	_, err := p.Clientset.CoreV1().Services(p.namespace).Create(context.Background(), service, metav1.CreateOptions{})
	return err
}

// UpdateService 更新service
func (p *Client) UpdateService(service *corev1.Service) error {
	_, err := p.Clientset.CoreV1().Services(p.namespace).Update(context.Background(), service, metav1.UpdateOptions{})
	return err
}

// DeleteService 删除service
func (p *Client) DeleteService(serviceName string) error {
	return p.Clientset.CoreV1().Services(p.namespace).Delete(context.Background(), serviceName, metav1.DeleteOptions{})
}

// GetIngress 查询ingress信息
func (p *Client) GetIngress(ingressName string) (*networkingv1.Ingress, error) {
	return p.Clientset.NetworkingV1().Ingresses(p.namespace).Get(context.Background(), ingressName, metav1.GetOptions{})
}

// QueryIngress 查询ingress列表
func (p *Client) QueryIngress(limit int64) ([]networkingv1.Ingress, error) {
	back, err := p.Clientset.NetworkingV1().Ingresses(p.namespace).List(context.Background(), metav1.ListOptions{
		Limit: limit,
	})
	if err != nil {
		return nil, err
	}

	return back.Items, nil
}

// CreateIngress 创建ingress
func (p *Client) CreateIngress(ingress *networkingv1.Ingress) error {
	_, err := p.Clientset.NetworkingV1().Ingresses(p.namespace).Create(context.Background(), ingress, metav1.CreateOptions{})
	return err
}

// UpdateIngress 更新ingress
func (p *Client) UpdateIngress(ingress *networkingv1.Ingress) error {
	_, err := p.Clientset.NetworkingV1().Ingresses(p.namespace).Update(context.Background(), ingress, metav1.UpdateOptions{})
	return err
}

// DeleteIngress 删除ingress
func (p *Client) DeleteIngress(ingressName string) error {
	return p.Clientset.NetworkingV1().Ingresses(p.namespace).Delete(context.Background(), ingressName, metav1.DeleteOptions{})
}
