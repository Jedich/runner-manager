package reconciler

import (
	"runner-controller-ecs/internal/delivery"
	"runner-controller-ecs/internal/infrastructure/logs"
	"runner-controller-ecs/internal/usecase/aws"
	"runner-controller-ecs/internal/usecase/credentials"
	gh "runner-controller-ecs/internal/usecase/github"
)

type Reconciler struct {
}

func NewReconciler() delivery.Reconciler {
	return &Reconciler{}
}

func (c *Reconciler) Init() error {
	creds := credentials.NewCredentialUC()

	awsUC := aws.NewAWSUC(creds)

	_, err := awsUC.GetTaskMetadata()
	if err != nil {
		return err
	}

	runner, err := awsUC.CreateRunner()
	if err != nil {
		return err
	}

	for _, r := range runner {
		logs.InfoF("%v", r)
	}

	githubUC := gh.NewGithubUC(creds)

	w, err := githubUC.GetWebhook(awsUC.GetPublicIP())
	if err != nil {
		return err
	}
	logs.InfoF("%v", w)

	return nil
}

func (c *Reconciler) Reconcile() error {
	logs.Info("Doing something...")
	return nil
}
