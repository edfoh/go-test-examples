package example1

type balanceGetter interface {
	GetBalance(accountID int) int
}

func Add(balanceGetter balanceGetter, accountID, amount int) int {
	bal := balanceGetter.GetBalance(accountID)

	return bal + amount
}
