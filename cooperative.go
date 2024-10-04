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

