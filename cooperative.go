package main
import (
  "fmt"
  "time"
)

type Member struct {
  ID int
  Name string
  JoinDate time.Time
  Savings float64
  Loan float64
}

type Cooperative struct {
  MemberList []Member
}

func (c *Cooperative) AddMember(name string){
  newMember := Member{
    ID: len(c.MemberList) +1,
    Name: name,
    JoinDate: time.Now(),
    Savings: 0.0,
    Loan: 0.0,
  }

  c.MemberList = append(c.MemberList, newMember)
  fmt.Printf("New member added with ID %d and name %s.\n", newMember.ID, newMember.Name)
}

func (c *Cooperative) SaveMoney(id int, amount float64){
  for i, member := range c.MemberList {
    if member.ID == id {
      c.MemberList[i].Savings += amount
      fmt.Printf("Rp%.2f saved to %s' account. Total savings:Rp%.2f\n", amount, member.Name, c.MemberList[i].Savings)
      return
    }
  }
  fmt.Println("Member Not Found.")
}

func (c *Cooperative) BorrowMoney(id int, amount float64){
  for i, member := range c.MemberList {
    if member.ID == id {
      c.MemberList[i].Loan += amount
      fmt.Printf("Rp%.2f borrowed by %s. Total loan: Rp%.2f\n", amount, member.Name, c.MemberList[i].Loan)
      return
    }
  }
  fmt.Println("Member Not Found.")
}

func (c *Cooperative) CheckSavings(id int){
  for _, member := range c.MemberList {
    if member.ID == id {
      fmt.Printf("%s's total savings are Rp%.2f\n", member.Name, member.Savings)
      return
    }
  }
  fmt.Println("Member Not Found.")
}

func (c *Cooperative) CheckLoan(id int){
  for _, member := range c.MemberList {
    if member.ID == id {
      fmt.Printf("%s's total loan is Rp%.2f\n", member.Name, member.Loan)
      return
    }
  }
  fmt.Println("Member Not Found.")
}


func main(){
  cooperative := Cooperative{}

  cooperative.AddMember("Alice")
  cooperative.AddMember("Bob")

  cooperative.SaveMoney(1, 1000000)
  cooperative.SaveMoney(2, 2500000)

  cooperative.BorrowMoney(1, 2000000)

  cooperative.CheckSavings(1)
  cooperative.CheckLoan(1)

  cooperative.CheckSavings(2)
  cooperative.CheckLoan(2)



}
