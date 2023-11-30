package iterator

import (
	"fmt"
)

// Author: Jones Binadiegha Gabriel

// This file explores the iterator design pattern
// What is it ?
// The iterator design pattern is a programing design pattern or concept that allows our program traverse or iterate a collection of data in a standard way.
// That data can be a collection of strings, integers, floats, or structs... basically a container (mostly structs). the data itself is inconsequencial but it's the concept that matters
// All programing languages has this pattern.const

// Why use it ?
// The iterator allows us (clients) to traverse data without knowing or exposing it's underlying data structure. ie. This allows of triverse a collection of data without prior knowledge of the type of being presented.
// allows addition of new traversal operations without modifying it's original interface

// How => see code snipet below

// First we create an Iterator interface
// This interface describes the methods that must be present to made a data iterable
type Iterator interface {
	hasNext() bool
	next()
	getCurrent() *User
}

// Now we create an aggregator interface: this describes what's combines the data and make it iterable.
type Aggregator interface {
	getIterator() Iterator
}

// the data struct for each item in the iterator
type User struct {
	name string
	age int
}

// Concrete implementation of the Iterator with the users struct
//  an implementation of an interator has to items: the index and the aggregated data. 
// the index is used to know the current level at which the aggregator is at,
type UserIterator struct {
	index int
	users []*User
}

// the has next method typically takes account for the current item which the client is at in the aggregation and indicates if there are more items in the aggregatoin or not. ie: if there is another item after the current item.
// this returns a boolean. false if the client is at the last element or true if there are more elements/items.
func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

// The next method moves the client through the elements in the aggregator. typically check if there is another item after the current item. this method greatly relies on the hasNext method. it simply increaments the index (which indicates the current item) if the hasNext method returns true.
func (u *UserIterator) next() {
	if u.hasNext() {
		u.index++
	}
}

// the get current method returns the current item being in the aggregation for the client's consumption. 
// this return nil if there isnt a current item. 
func (u *UserIterator) getCurrent() *User {
	if u.hasNext() {
		return u.users[u.index]
	}
	return nil
}

// How does it all work?
// picture this: you have 3 user structs you want to aggregate and make iterable. let's call it len_of_users = 3
// you have the index for each user struct; with 0 being the default.
// on the first index(0) where index i = 0. hasNext = i + 1 < len_of_users... this translates to hasNext = true because 1 is less than 3
// since hasNext returns true for index 0, it means the current struct is not the last struct so we can proceed to move to the next struct. 
// we can also get the current struct using getCurrent().
// this continues until hasNext returns true... then it means that [i + 1 < len_of_users is false]. ie. 3 is not less than 3.
// since hasNext returns false, then next() which is basically i + 1 will not run. simply put, you are at the last struct.
// however, you can still get the current struct using getCurrent().

// I hope this explanation makes sense. feel free to contribute and make corrections.

// let's continue.


// this produces a user iterator.
type userAggregator struct {
	users []*User
}

func (u *userAggregator) getIterator() Iterator {
		return &UserIterator{
			users: u.users,
		}
}


func App() {

	us1 := &User{
		name: "Jones",
		age: 28,
	}

	us2 := &User{
		name: "Gabriel",
		age: 30,
	}

	us3 := &User{
		name: "Binadiegha",
		age: 32,
	}

	aggr := &userAggregator{ 
		users: []*User{us1, us2, us3},
	}

	i := aggr.getIterator()



	for i.hasNext() {
		el := i.getCurrent()
		fmt.Println(el)
		i.next()
	}
}