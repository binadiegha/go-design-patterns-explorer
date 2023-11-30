package main

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
	getCurrent() interface{}
}

// Now we create an aggregator interface: this describes what's combines the data and make it iterable.
type aggregator interface {
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
func (u *UserIterator) hasNext() bool {
	if u.index < len(u.users) {
		return true
	}
	return false
}

func (u *UserIterator) next() {
	if u.hasNext() {
		u.index++
	}
}

func (u *UserIterator) getCurrent() *User {
	if u.hasNext() {
		return u.users[u.index]
	}
	return nil
}



func UseIterator() {

}