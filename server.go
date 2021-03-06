package funding
// continue: 
//Make it a service
//http://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial
import "fmt"

type FundServer struct {
	Commands chan interface{}
	fund Fund
}

func NewFundServer(initialBalance int) *FundServer {
	server := &FundServer{
		// make() creates builtins like channels, maps and slices
		Commands: make(chan interface{}),
		fund: *NewFund(initialBalance),
	}

	//Spawn off the server's main loop immediately
	go server.loop()
	return server
}

func (s *FundServer) loop() {
	
	for command := range s.Commands {
		// command is just an interface{} but we can check its real type
		switch command.(type) {
			case WithdrawCommand: 
				// And then use a "type assertion" to convert it 
				withdrawal := command.(WithdrawCommand)
				s.fund.Withdraw(withdrawal.Amount)

			case BalanceCommand: 
				getBalance := command.(BalanceCommand)
				balance := s.fund.Balance()
				getBalance.Response <- balance

			case DepositCommand:
				deposit := command.(DepositCommand)
				s.fund.Deposit(deposit.Amount)

			default:
				panic(fmt.Sprintf("Unrecognized command: %v", command))
		}
	}
}

type WithdrawCommand struct {
	Amount int
}

type BalanceCommand struct {
	Response chan int
}
type DepositCommand struct {
	Amount int
}