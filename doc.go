// Package ternary provides methods similar to other languages ternary operator
/*

Have you ever wanted those pesky ternary operators for java/c/javascript/everything else?

You know, the special sauce that allows you to do this stuff.

    var something = some_boolean ? "Oh hai! it was TRUE!!" : "Aww shucks; It was False :(";

But then you came over to the go world and realized the shortest way to do this was 4 lines of code.

    var something = "Aww shucks; It was False :("
    if some_boolean {
        something = "Oh hai! it was TRUE!!"
    }

Yeah, I had to go club some baby seals when I made that realization.

Solution

But wait!  I have crawled through the river of crappy code and propose this package to keep you clean on the other side.

Now you you can use the following to assist with your ternary operator needs!

    import "github.com/bign8/ternary"

    // ...

    var something = ternary.String(some_boolean, "Oh hai! it was TRUE!!", "Aww shucks; It was False :(")

Much better! Oh, but you say you want to use ternary operators with something other than strings?

Well go ahead!!! We support most of the core types: See index below for more!

The Catch

Unfortunantly, I can't implement EVERY type.
There are complicated ones like Array, Chan, Func, Map, Ptr, Slice and Struct that can take on any form.

Thank goodness go has typecasting!
If you run into something this packages doesn't support, just use the Interface comparison!

    type Awesome struct {
        // Many fields I don't care about
    }

    var one = Awesome{}
    var two = Awesome{}

    var answer = ternary.Interface(some_boolean, one, two).(Awesome)
*/
package ternary

/* This program is free software. It comes without any warranty, to
 * the extent permitted by applicable law. You can redistribute it
 * and/or modify it under the terms of the Do What The Fuck You Want
 * To Public License, Version 2, as published by Sam Hocevar. See
 * http://www.wtfpl.net/ for more details. */

//go:generate go run gen.go
