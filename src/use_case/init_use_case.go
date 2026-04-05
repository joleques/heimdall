package usecase

import "context"

type InitGateway interface {
	InitTarget(ctx context.Context, request InitRequest) (InitResult, error)
}

type InitTargetUseCase struct {
	gateway InitGateway
}

func NewInitTargetUseCase(gateway InitGateway) InitTargetUseCase {
	return InitTargetUseCase{gateway: gateway}
}

func (uc InitTargetUseCase) Execute(ctx context.Context, request InitRequest) (InitResult, error) {
	return uc.gateway.InitTarget(ctx, request)
}

var _ InitTarget = InitTargetUseCase{}
