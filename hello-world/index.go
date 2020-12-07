package main

import (
	"context"
	"fmt"
	prisma "hello-world/generated/prisma-client"
)

func main() {
	client := prisma.New(nil)
	ctx := context.TODO()

	email := "alice@prisma.io"
	postsByUser, err := client.User(prisma.UserWhereUniqueInput{
		Email: &email,
	}).Posts(nil).Exec(ctx)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", postsByUser)

}
