package contracts

type FinishWorkOrderUC interface {
	Execute(ID string) error
}
