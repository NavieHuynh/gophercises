package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "linkparser"
)

var _ = Describe("Main", func() {
	DescribeTable("Testing Examples",
		func(filename string, expected []Link) {
			actual := GetLinksFromFile(filename)
			Expect(actual).Should(Equal(expected))
		},
		Entry("Example1", "./ex1.html", []Link{
			*NewLink("/other-page", "A link to another page"),
		}),
		Entry("Example2", "./ex2.html", []Link{
			*NewLink("https://www.twitter.com/joncalhoun", "Check me out on twitter"),
			*NewLink("https://github.com/gophercises", "Gophercises is on Github!"),
		}),
		Entry("Example3", "./ex3.html", []Link{
			*NewLink("#", "Login"),
			*NewLink("/lost", "Lost? Need help?"),
			*NewLink("https://twitter.com/marcusolsson", "@marcusolsson"),
		}),
		Entry("Example4", "./ex4.html", []Link{
			*NewLink("/dog-cat", "dog cat"),
		}),
	)
})
